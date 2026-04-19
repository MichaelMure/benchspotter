package engine

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"iter"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/go-git/go-billy/v5"
	"github.com/google/pprof/profile"
	"golang.org/x/sys/execabs"
)

type Profile int

const (
	ProfileBench Profile = iota
	ProfileCPU
	ProfileMem
	ProfileMutex
	ProfileBlock
)

var replacer = strings.NewReplacer(
	"/", "ᚋ",
	".", "ᚗ",
	"-", "ᚑ",
	"~", "א",
)

var reverseReplacer = strings.NewReplacer(
	"ᚋ", "/",
	"ᚗ", ".",
	"ᚑ", "-",
	"א", "~",
)

func RunProfile(ctx context.Context, storage billy.Filesystem, id string, benches []BenchInfo, profile Profile) func() error {
	next, _ := iter.Pull(func(yield func(error) bool) {
		for _, infos := range benches {
			var dir, filename, option string
			switch profile {
			case ProfileCPU:
				dir, filename, option = "cpu", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-cpuprofile"
			case ProfileMem:
				dir, filename, option = "mem", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-memprofile"
			case ProfileMutex:
				dir, filename, option = "mutex", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-mutexprofile"
			case ProfileBlock:
				dir, filename, option = "block", replacer.Replace(infos.Package+"."+infos.Name)+".profile", "-blockprofile"
			default:
				panic("invalid profile")
			}

			subDir := storage.Join(sessionDir, id, dir)
			if err := storage.MkdirAll(subDir, 0755); err != nil {
				yield(err)
				return
			}
			outFilename := storage.Join(storage.Root(), subDir, filename)

			cmd := execabs.CommandContext(ctx, "go", "test",
				"-bench", infos.Regex(),
				option+"="+outFilename,
				"-run", "^$", ".")
			cmd.Dir = infos.Package

			if out, err := cmd.CombinedOutput(); err != nil {
				yield(fmt.Errorf("failed to profile: %w, %s", err, out))
				return
			}
			if !yield(nil) {
				return
			}
		}
	})
	return func() error {
		err, _ := next()
		return err
	}
}

type ProfileFunc struct {
	Name       string
	Flat       time.Duration
	Cumulative time.Duration
	FlatPct    float64
	CumPct     float64
}

// ProfileDir returns the storage subdirectory for a given profile type.
func ProfileDir(p Profile) string {
	switch p {
	case ProfileCPU:
		return "cpu"
	case ProfileMem:
		return "mem"
	case ProfileMutex:
		return "mutex"
	case ProfileBlock:
		return "block"
	default:
		panic("invalid profile")
	}
}

// ListProfileBenchmarks returns the benchmark names that have profiles stored
// for the given session and profile type.
func ListProfileBenchmarks(fs billy.Filesystem, sessionPath string, p Profile) ([]string, error) {
	dir := filepath.Join(sessionPath, ProfileDir(p))
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("no %s profile found for this session", ProfileDir(p))
	}
	var names []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".profile") {
			continue
		}
		decoded := reverseReplacer.Replace(strings.TrimSuffix(e.Name(), ".profile"))
		// decoded is "package.BenchmarkName" — benchmark name is after the last dot
		if idx := strings.LastIndex(decoded, "."); idx >= 0 {
			names = append(names, decoded[idx+1:])
		}
	}
	return names, nil
}

// ReadProfileFunctions parses the pprof file for the given session, profile
// type, and benchmark name, and returns the functions sorted by flat time
// descending. Use ListProfileBenchmarks to enumerate valid bench values.
func ReadProfileFunctions(fs billy.Filesystem, sessionPath string, p Profile, bench string) ([]ProfileFunc, error) {
	prof, err := readProfile(fs, sessionPath, p, bench)
	if err != nil {
		return nil, err
	}
	return profileToFuncs(prof, p), nil
}

// ReadProfileRaw parses the pprof file for the given session, profile type, and
// benchmark name, and returns the re-serialized binary suitable for piping to
// go tool pprof. Use ListProfileBenchmarks to enumerate valid bench values.
func ReadProfileRaw(fs billy.Filesystem, sessionPath string, p Profile, bench string) ([]byte, error) {
	prof, err := readProfile(fs, sessionPath, p, bench)
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	if err := prof.Write(&buf); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// readProfile opens the single pprof file for bench and parses it.
// bench must be a non-empty name returned by ListProfileBenchmarks.
func readProfile(fs billy.Filesystem, sessionPath string, p Profile, bench string) (*profile.Profile, error) {
	dir := filepath.Join(sessionPath, ProfileDir(p))
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("no %s profile found for this session", ProfileDir(p))
	}

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".profile") {
			continue
		}
		decoded := reverseReplacer.Replace(strings.TrimSuffix(e.Name(), ".profile"))
		name := decoded
		if idx := strings.LastIndex(decoded, "."); idx >= 0 {
			name = decoded[idx+1:]
		}
		if name != bench {
			continue
		}
		f, err := fs.Open(filepath.Join(dir, e.Name()))
		if err != nil {
			return nil, err
		}
		defer f.Close()
		data, err := io.ReadAll(f)
		if err != nil {
			return nil, err
		}
		return profile.ParseData(data)
	}

	return nil, fmt.Errorf("no %s profile found for benchmark %q in this session", ProfileDir(p), bench)
}

// profileToFuncs aggregates pprof samples into per-function flat/cumulative
// totals. Flat counts a function only when it appears at the top of the call
// stack (i.e. it was on-CPU or holding the lock at sample time). Cumulative
// counts it whenever it appears anywhere in the stack, deduplicated per sample.
// Percentages are relative to the total value across all samples.
// Results are sorted by flat descending, then cumulative descending.
//
// This reimplements the core of github.com/google/pprof/internal/report, which
// is not importable outside the pprof module.
func profileToFuncs(prof *profile.Profile, p Profile) []ProfileFunc {
	valueIdx := pickValueIndex(prof, p)

	type entry struct {
		name string
		flat int64
		cum  int64
	}
	byFunc := make(map[string]*entry)

	var total int64
	for _, s := range prof.Sample {
		if len(s.Value) <= valueIdx {
			continue
		}
		v := s.Value[valueIdx]
		total += v

		if len(s.Location) > 0 {
			for _, line := range s.Location[0].Line {
				name := line.Function.Name
				e := byFunc[name]
				if e == nil {
					e = &entry{name: name}
					byFunc[name] = e
				}
				e.flat += v
			}
		}
		seen := make(map[string]bool)
		for _, loc := range s.Location {
			for _, line := range loc.Line {
				name := line.Function.Name
				if seen[name] {
					continue
				}
				seen[name] = true
				e := byFunc[name]
				if e == nil {
					e = &entry{name: name}
					byFunc[name] = e
				}
				e.cum += v
			}
		}
	}

	result := make([]ProfileFunc, 0, len(byFunc))
	for _, e := range byFunc {
		var flatPct, cumPct float64
		if total > 0 {
			flatPct = float64(e.flat) / float64(total) * 100
			cumPct = float64(e.cum) / float64(total) * 100
		}
		result = append(result, ProfileFunc{
			Name:       e.name,
			Flat:       time.Duration(e.flat),
			Cumulative: time.Duration(e.cum),
			FlatPct:    flatPct,
			CumPct:     cumPct,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Flat != result[j].Flat {
			return result[i].Flat > result[j].Flat
		}
		return result[i].Cumulative > result[j].Cumulative
	})

	return result
}

// pickValueIndex returns the index into prof.SampleType whose Type we want to
// aggregate. A pprof file stores multiple parallel value columns per sample;
// which column is meaningful depends on the profile type:
//
//	CPU:         [samples/count, cpu/nanoseconds]                                       → cpu
//	Mem:         [alloc_objects/count, alloc_space/bytes, inuse_objects/count, inuse_space/bytes] → inuse_space
//	Block/Mutex: [contentions/count, delay/nanoseconds]                                → delay
//
// Falls back to index 0 if the preferred type is not found.
func pickValueIndex(prof *profile.Profile, p Profile) int {
	preferred := map[Profile]string{
		ProfileCPU:   "cpu",
		ProfileMem:   "inuse_space",
		ProfileBlock: "delay",
		ProfileMutex: "delay",
	}
	want, ok := preferred[p]
	if ok {
		for i, st := range prof.SampleType {
			if st.Type == want {
				return i
			}
		}
	}
	return 0
}
