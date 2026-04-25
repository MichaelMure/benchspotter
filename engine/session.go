package engine

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/go-git/go-billy/v5/util"
	"github.com/google/uuid"

	"benchspotter/commands/execenv"
)

const sessionDir = "sessions"
const metaFilename = "meta.json"
const GitDiffFilename = "git.diff"

// sessionMeta is a metadata record in the session folder
type sessionMeta struct {
	Name      string   `json:"name,omitempty"`
	Benches   []string `json:"benchs,omitempty"`
	GitCommit string   `json:"git_commit,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

func PrepareSession(ctx context.Context, env *execenv.Env, name string, benches []BenchInfo) (string, error) {
	uid, err := uuid.NewV7()
	if err != nil {
		return "", fmt.Errorf("failed to generate session id: %w", err)
	}
	id := uid.String()

	err = recordDiffIfAvailable(ctx, env, id)
	if err != nil {
		return "", err
	}

	err = recordMeta(ctx, env, id, name, benches)
	if err != nil {
		return "", err
	}

	return id, nil
}

func recordDiffIfAvailable(ctx context.Context, env *execenv.Env, id string) error {
	filename := filepath.Join(sessionDir, id, GitDiffFilename)
	diffFile, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create diff file: %w", err)
	}

	cmd := env.Repo.Cmd(ctx, "git", "diff")
	cmd.Stdout = &CountingWriter{writer: diffFile}
	err = cmd.Run()
	_ = diffFile.Close()
	if err != nil || cmd.Stdout.(*CountingWriter).bytesWritten == 0 {
		_ = env.Repo.Storage().Remove(filename)
	}
	return nil
}

func recordMeta(ctx context.Context, env *execenv.Env, id string, name string, benches []BenchInfo) error {
	filename := filepath.Join(sessionDir, id, metaFilename)
	f, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create meta file: %w", err)
	}
	defer f.Close()

	meta := sessionMeta{Name: name}

	meta.Benches = make([]string, len(benches))
	for i, bench := range benches {
		meta.Benches[i] = bench.Name
	}

	if commit, err := getCommitIfAvailable(ctx, env); err == nil {
		meta.GitCommit = commit
	}

	err = json.NewEncoder(f).Encode(meta)
	if err != nil {
		return fmt.Errorf("failed to encode meta file: %w", err)
	}
	return nil
}

func getCommitIfAvailable(ctx context.Context, env *execenv.Env) (string, error) {
	cmd := env.Repo.Cmd(ctx, "git", "rev-parse", "HEAD")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

type CountingWriter struct {
	writer       io.Writer
	bytesWritten int
}

func (c *CountingWriter) Write(p []byte) (int, error) {
	n, err := c.writer.Write(p)
	c.bytesWritten += n
	return n, err
}

type SessionInfo struct {
	Id        string
	Name      string
	HumanName string
	Root      string // root of the storage
	Path      string // path in storage
	Time      time.Time
	Benches   []string
	GitCommit string
	Tags      []string

	uid uuid.UUID
	fs  billy.Filesystem
}

// LocateSessions reads all sessions from fs, assigns human-readable de-duplicated
// names, and returns them sorted by creation time (oldest first).
// Optional tagFilter values restrict results to sessions carrying any of those tags.
func LocateSessions(fs billy.Filesystem, tagFilter ...string) ([]*SessionInfo, error) {
	dirs, err := fs.ReadDir(sessionDir)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("no sessions found, use the `bench` command to create one")
	}
	if err != nil {
		return nil, err
	}

	res := make([]*SessionInfo, len(dirs))
	for i, dir := range dirs {
		f, err := fs.Open(filepath.Join(sessionDir, dir.Name(), metaFilename))
		if err != nil {
			return nil, fmt.Errorf(`failed to open meta file for session "%s": %w`, dir.Name(), err)
		}
		meta := &sessionMeta{}
		err = json.NewDecoder(f).Decode(meta)
		if err != nil {
			return nil, fmt.Errorf(`failed to decode meta file for session "%s": %w`, dir.Name(), err)
		}

		uid, err := uuid.Parse(dir.Name())
		if err != nil {
			return nil, fmt.Errorf("invalid id: %w", err)
		}
		if uid.Version() != 7 {
			return nil, fmt.Errorf("invalid id: uuid is not v7")
		}

		// Extract the first 48 bits (6 bytes) as timestamp in milliseconds
		// UUIDv7 stores timestamp as big-endian uint48 in the first 48 bits
		timestampMilli := int64(binary.BigEndian.Uint32(uid[0:4]))<<16 |
			int64(binary.BigEndian.Uint16(uid[4:6]))

		res[i] = &SessionInfo{
			Id:        dir.Name(),
			Name:      meta.Name,
			Root:      fs.Root(),
			Path:      filepath.Join(sessionDir, dir.Name()),
			Time:      time.UnixMilli(timestampMilli),
			Benches:   meta.Benches,
			GitCommit: meta.GitCommit,
			Tags:      meta.Tags,

			uid: uid,
			fs:  fs,
		}
	}

	err = generateNames(res)
	if err != nil {
		return nil, err
	}

	if len(tagFilter) > 0 {
		want := make(map[string]struct{}, len(tagFilter))
		for _, t := range tagFilter {
			want[t] = struct{}{}
		}
		filtered := res[:0]
		for _, s := range res {
			for _, t := range s.Tags {
				if _, ok := want[t]; ok {
					filtered = append(filtered, s)
					break
				}
			}
		}
		res = filtered
	}

	return res, nil
}

func generateNames(res []*SessionInfo) error {
	// There are two schemes:
	// 1. for user-defined names we quote them ("foo") or add a number if there is a conflict ("foo"-1)
	// 2. for auto-generated names we use the timestamp ("")

	uidToSuffix := func(uid uuid.UUID, length int) string {
		if length == 0 {
			return ""
		}
		bytesNeeded := (length + 1) / 2
		bytes := make([]byte, bytesNeeded)
		// only use rand_b spans bytes 8-15
		copy(bytes[0:], uid[9:16])
		if bytesNeeded > 7 {
			_, _ = rand.Read(bytes[7:])
		}
		return "-" + hex.EncodeToString(bytes[:length/2])
	}

	var allNames = make(map[string]int, len(res))
	for i, info := range res {
		name := info.Name
		if len(name) == 0 {
			name = info.Time.Format("06-Jan-02")
		}
		allNames[name] = allNames[name] + 1
		res[i].HumanName = name
	}

	// add suffix where necessary
	suffixLen := 1
	for len(allNames) < len(res) {
		for i, info := range res {
			name := info.Name
			if len(name) == 0 {
				name = info.Time.Format("06-Jan-02")
			}
			if allNames[name] > 1 {
				name = name + uidToSuffix(info.uid, suffixLen)
			}
			res[i].HumanName = name
		}
		allNames = make(map[string]int, len(res))
		for _, info := range res {
			allNames[info.HumanName] = allNames[info.HumanName] + 1
		}
		suffixLen++
		if suffixLen > 128 {
			return fmt.Errorf("failed to generate unique names")
		}
	}

	return nil
}

func (s SessionInfo) BenchPath() string {
	return filepath.Join(s.Path, benchFilename)
}

func (s SessionInfo) BenchFullPath() string {
	return filepath.Join(s.Root, s.Path, benchFilename)
}

func (s SessionInfo) HasGitDiff() bool {
	_, err := s.fs.Stat(filepath.Join(s.Path, GitDiffFilename))
	return err == nil
}

func (s SessionInfo) HasBench() bool {
	_, err := s.fs.Stat(filepath.Join(s.Path, benchFilename))
	return err == nil
}

func (s SessionInfo) HasProfile(p Profile) bool {
	switch p {
	case ProfileEscape:
		_, err := s.fs.Stat(filepath.Join(s.Path, EscapeFilename))
		return err == nil
	case ProfileInline:
		_, err := s.fs.Stat(filepath.Join(s.Path, InlineFilename))
		return err == nil
	default:
		dir := filepath.Join(s.Path, ProfileDir(p))
		entries, err := s.fs.ReadDir(dir)
		return err == nil && len(entries) > 0
	}
}

func (s SessionInfo) OpenFile(name string) (billy.File, error) {
	return s.fs.Open(filepath.Join(s.Path, name))
}

// TagSession appends tag to the session's meta.json. It is a no-op if the tag
// is already present.
func TagSession(fs billy.Filesystem, sessionPath string, tag string) error {
	return updateMeta(fs, sessionPath, func(meta *sessionMeta) {
		for _, t := range meta.Tags {
			if t == tag {
				return
			}
		}
		meta.Tags = append(meta.Tags, tag)
	})
}

// UntagSession removes tag from the session's meta.json. It is a no-op if the
// tag is not present.
func UntagSession(fs billy.Filesystem, sessionPath string, tag string) error {
	return updateMeta(fs, sessionPath, func(meta *sessionMeta) {
		tags := meta.Tags[:0]
		for _, t := range meta.Tags {
			if t != tag {
				tags = append(tags, t)
			}
		}
		meta.Tags = tags
	})
}

// updateMeta reads meta.json, applies fn, then writes it back.
func updateMeta(fs billy.Filesystem, sessionPath string, fn func(*sessionMeta)) error {
	metaPath := filepath.Join(sessionPath, metaFilename)
	f, err := fs.Open(metaPath)
	if err != nil {
		return fmt.Errorf("failed to open meta file: %w", err)
	}
	meta := &sessionMeta{}
	if err = json.NewDecoder(f).Decode(meta); err != nil {
		_ = f.Close()
		return fmt.Errorf("failed to decode meta file: %w", err)
	}
	_ = f.Close()

	fn(meta)

	out, err := fs.Create(metaPath)
	if err != nil {
		return fmt.Errorf("failed to write meta file: %w", err)
	}
	defer out.Close()
	return json.NewEncoder(out).Encode(meta)
}

// RenameSession updates the human-readable name stored in the session's meta.json.
func RenameSession(fs billy.Filesystem, sessionPath string, name string) error {
	return updateMeta(fs, sessionPath, func(meta *sessionMeta) { meta.Name = name })
}

// RemoveSession deletes a session directory and all its contents.
func RemoveSession(fs billy.Filesystem, sessionPath string) error {
	return util.RemoveAll(fs, sessionPath)
}
