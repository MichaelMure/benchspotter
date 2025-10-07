package engine

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/google/uuid"
	"golang.org/x/sys/execabs"

	"benchspotter/commands/execenv"
)

const sessionDir = "sessions"
const metaFilename = "meta.json"

// sessionMeta is a metadata record in the session folder
type sessionMeta struct {
	Name    string   `json:"name,omitempty"`
	Benches []string `json:"benchs,omitempty"`
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

	err = recordMeta(env, id, name, benches)
	if err != nil {
		return "", err
	}

	return id, nil
}

func recordDiffIfAvailable(ctx context.Context, env *execenv.Env, id string) error {
	filename := filepath.Join(sessionDir, id, "git.diff")
	diffFile, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create diff file: %w", err)
	}

	cmd := execabs.CommandContext(ctx, "git", "diff")
	cmd.Stdout = &CountingWriter{writer: diffFile}
	err = cmd.Run()
	_ = diffFile.Close()
	if err != nil || cmd.Stdout.(*CountingWriter).bytesWritten == 0 {
		_ = env.Repo.Storage().Remove(filename)
	}
	return nil
}

func recordMeta(env *execenv.Env, id string, name string, benches []BenchInfo) error {
	filename := filepath.Join(sessionDir, id, metaFilename)
	f, err := env.Repo.Storage().Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create meta file: %w", err)
	}
	defer f.Close()

	benchesStr := make([]string, len(benches))
	for i, bench := range benches {
		benchesStr[i] = bench.Name
	}

	err = json.NewEncoder(f).Encode(&sessionMeta{
		Name:    name,
		Benches: benchesStr,
	})
	if err != nil {
		return fmt.Errorf("failed to encode meta file: %w", err)
	}
	return nil
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
	uid       uuid.UUID
	Name      string
	HumanName string
	Root      string // root of the storage
	Path      string // path in storage
	Time      time.Time
	Benches   []string
}

func LocateSessions(fs billy.Filesystem) ([]*SessionInfo, error) {
	dirs, err := fs.ReadDir(sessionDir)
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
			Id:      dir.Name(),
			uid:     uid,
			Name:    meta.Name,
			Root:    fs.Root(),
			Path:    filepath.Join(sessionDir, dir.Name()),
			Time:    time.UnixMilli(timestampMilli),
			Benches: meta.Benches,
		}
	}

	err = generateNames(res)
	if err != nil {
		return nil, err
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

	var success bool
	for suffixLen := 0; suffixLen <= 128; suffixLen++ {
		allNames := make(map[string]struct{}, len(res))
		success = true
		for i, info := range res {
			if len(info.Name) == 0 {
				res[i].HumanName = fmt.Sprintf(`%s%s`,
					info.Time.Format("06-Jan-02"),
					uidToSuffix(info.uid, suffixLen),
				)
			} else {
				res[i].HumanName = fmt.Sprintf(`"%s"%s`,
					info.Name, uidToSuffix(info.uid, suffixLen),
				)
			}
			if _, ok := allNames[info.HumanName]; ok {
				success = false
				break
			}
			allNames[info.HumanName] = struct{}{}
		}
		if success {
			break
		}
	}
	if !success {
		return fmt.Errorf("failed to generate unique names")
	}

	return nil
}

func (s SessionInfo) BenchPath() string {
	return filepath.Join(s.Path, benchFilename)
}

func (s SessionInfo) BenchFullPath() string {
	return filepath.Join(s.Root, s.Path, benchFilename)
}
