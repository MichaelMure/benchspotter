package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/klauspost/cpuid/v2"
)

// MachineInfo captures the hardware and OS context of the machine that produced
// a session. Stored in meta.json.
type MachineInfo struct {
	GOOS          string `json:"goos"`
	GOARCH        string `json:"goarch"`
	NumCPU        int    `json:"num_cpu"`
	PhysicalCores int    `json:"physical_cores,omitempty"`
	ThreadsPerCore int   `json:"threads_per_core,omitempty"`
	CPU           string `json:"cpu,omitempty"`
	CacheL1D      int    `json:"cache_l1d,omitempty"`
	CacheL2       int    `json:"cache_l2,omitempty"`
	CacheL3       int    `json:"cache_l3,omitempty"`
	VM            bool   `json:"vm,omitempty"`
}

// Hash returns a short, versioned fingerprint suitable for grouping sessions
// by machine. Input fields are null-separated to avoid boundary ambiguity.
// The v1 prefix allows the algorithm to evolve without silently invalidating
// comparisons.
func (m MachineInfo) Hash() string {
	vm := "0"
	if m.VM {
		vm = "1"
	}
	h := sha256.New()
	_, _ = h.Write([]byte(strings.Join([]string{
		m.GOOS, m.GOARCH, m.CPU,
		strconv.Itoa(m.NumCPU),
		strconv.Itoa(m.PhysicalCores),
		strconv.Itoa(m.ThreadsPerCore),
		strconv.Itoa(m.CacheL1D),
		strconv.Itoa(m.CacheL2),
		strconv.Itoa(m.CacheL3),
		vm,
	}, "\x00")))
	return "v1:" + hex.EncodeToString(h.Sum(nil)[:8])
}

// Summary returns a compact human-readable description of the machine.
func (m MachineInfo) Summary() string {
	var sb strings.Builder
	sb.WriteString(m.GOOS + "/" + m.GOARCH)
	if m.PhysicalCores > 0 && m.ThreadsPerCore > 1 {
		fmt.Fprintf(&sb, " · %dc (%dc×%dt)", m.NumCPU, m.PhysicalCores, m.ThreadsPerCore)
	} else {
		fmt.Fprintf(&sb, " · %dc", m.NumCPU)
	}
	if m.CPU != "" {
		sb.WriteString(" · " + m.CPU)
	}
	if m.VM {
		sb.WriteString(" · VM")
	}
	return sb.String()
}

// FormatMachineLine formats machine info and Go version as a single display line.
func FormatMachineLine(m *MachineInfo, goVersion string) string {
	if m == nil {
		return ""
	}
	s := m.Summary()
	if goVersion != "" {
		s += " · " + goVersion
	}
	return s
}

// MachineContext assigns numeric identifiers (1, 2, ...) to distinct machines
// present in a session collection. Nil when all sessions share one machine.
type MachineContext struct {
	Labels  map[string]int // session ID → 1-based machine number (0 = no machine info)
	Entries []MachineEntry // legend, in label order
}

// MachineEntry is one row in the machine legend.
type MachineEntry struct {
	Machine   MachineInfo
	GoVersion string
}

// NewMachineContext builds a MachineContext from sessions. Returns nil when
// all sessions with machine info share the same machine+Go version (or none
// have machine info).
func NewMachineContext(sessions []*SessionInfo) *MachineContext {
	type key struct{ hash, goVersion string }
	seen := make(map[key]int)
	var entries []MachineEntry

	for _, s := range sessions {
		if s.Machine == nil {
			continue
		}
		k := key{s.Machine.Hash(), s.GoVersion}
		if _, ok := seen[k]; !ok {
			n := len(entries) + 1
			seen[k] = n
			entries = append(entries, MachineEntry{
				Machine:   *s.Machine,
				GoVersion: s.GoVersion,
			})
		}
	}

	if len(entries) <= 1 {
		return nil
	}

	labels := make(map[string]int, len(sessions))
	for _, s := range sessions {
		if s.Machine == nil {
			continue
		}
		k := key{s.Machine.Hash(), s.GoVersion}
		labels[s.Id] = seen[k]
	}

	return &MachineContext{Labels: labels, Entries: entries}
}

func collectMachineInfo() MachineInfo {
	return MachineInfo{
		GOOS:           runtime.GOOS,
		GOARCH:         runtime.GOARCH,
		NumCPU:         runtime.NumCPU(),
		PhysicalCores:  cpuid.CPU.PhysicalCores,
		ThreadsPerCore: cpuid.CPU.ThreadsPerCore,
		CPU:            strings.Join(strings.Fields(cpuid.CPU.BrandName), " "),
		CacheL1D:       cpuid.CPU.Cache.L1D,
		CacheL2:        cpuid.CPU.Cache.L2,
		CacheL3:        cpuid.CPU.Cache.L3,
		VM:             cpuid.CPU.VM(),
	}
}
