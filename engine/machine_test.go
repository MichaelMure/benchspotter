package engine

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	machineA = MachineInfo{
		GOOS: "linux", GOARCH: "amd64", NumCPU: 12,
		PhysicalCores: 6, ThreadsPerCore: 2,
		CPU: "Intel Core i7-8700K", CacheL1D: 32768, CacheL2: 262144, CacheL3: 12582912,
	}
	machineB = MachineInfo{
		GOOS: "darwin", GOARCH: "arm64", NumCPU: 10,
		CPU: "Apple M1 Pro",
	}
)

func TestMachineInfoHash(t *testing.T) {
	t.Run("same machine same hash", func(t *testing.T) {
		require.Equal(t, machineA.Hash(), machineA.Hash())
	})

	t.Run("different machines different hash", func(t *testing.T) {
		require.NotEqual(t, machineA.Hash(), machineB.Hash())
	})

	t.Run("hash prefix", func(t *testing.T) {
		require.Contains(t, machineA.Hash(), "v1:")
	})

	t.Run("field sensitivity", func(t *testing.T) {
		m := machineA
		m.NumCPU = 8
		require.NotEqual(t, machineA.Hash(), m.Hash())

		m = machineA
		m.PhysicalCores = 4
		require.NotEqual(t, machineA.Hash(), m.Hash())

		m = machineA
		m.VM = true
		require.NotEqual(t, machineA.Hash(), m.Hash())

		m = machineA
		m.CacheL3 = 0
		require.NotEqual(t, machineA.Hash(), m.Hash())
	})
}

func TestMachineInfoSummary(t *testing.T) {
	t.Run("with HT topology", func(t *testing.T) {
		s := machineA.Summary()
		require.Contains(t, s, "linux/amd64")
		require.Contains(t, s, "12c (6c×2t)")
		require.Contains(t, s, "Intel Core i7-8700K")
	})

	t.Run("without HT", func(t *testing.T) {
		m := MachineInfo{GOOS: "linux", GOARCH: "amd64", NumCPU: 4, PhysicalCores: 4, ThreadsPerCore: 1}
		s := m.Summary()
		require.Contains(t, s, "4c")
		require.NotContains(t, s, "×")
	})

	t.Run("VM flag", func(t *testing.T) {
		m := machineA
		m.VM = true
		require.Contains(t, m.Summary(), "VM")
		require.NotContains(t, machineA.Summary(), "VM")
	})

	t.Run("no CPU name", func(t *testing.T) {
		m := MachineInfo{GOOS: "linux", GOARCH: "amd64", NumCPU: 2}
		s := m.Summary()
		require.Contains(t, s, "linux/amd64")
		require.Contains(t, s, "2c")
	})
}

func TestFormatMachineLine(t *testing.T) {
	t.Run("nil machine", func(t *testing.T) {
		require.Equal(t, "", FormatMachineLine(nil, "go1.22"))
	})

	t.Run("with go version", func(t *testing.T) {
		s := FormatMachineLine(&machineA, "go1.22.3")
		require.Contains(t, s, "go1.22.3")
		require.Contains(t, s, "linux/amd64")
	})

	t.Run("without go version", func(t *testing.T) {
		s := FormatMachineLine(&machineA, "")
		require.NotContains(t, s, "go")
	})
}

func TestNewMachineContext(t *testing.T) {
	mkSession := func(id string, m *MachineInfo, goVer string) *SessionInfo {
		return &SessionInfo{Id: id, Machine: m, GoVersion: goVer}
	}
	mA := machineA
	mB := machineB

	t.Run("nil when no sessions", func(t *testing.T) {
		require.Nil(t, NewMachineContext(nil))
	})

	t.Run("nil when all same machine", func(t *testing.T) {
		sessions := []*SessionInfo{
			mkSession("s1", &mA, "go1.22"),
			mkSession("s2", &mA, "go1.22"),
		}
		require.Nil(t, NewMachineContext(sessions))
	})

	t.Run("nil when no machine info", func(t *testing.T) {
		sessions := []*SessionInfo{
			mkSession("s1", nil, ""),
			mkSession("s2", nil, ""),
		}
		require.Nil(t, NewMachineContext(sessions))
	})

	t.Run("two machines", func(t *testing.T) {
		sessions := []*SessionInfo{
			mkSession("s1", &mA, "go1.22"),
			mkSession("s2", &mB, "go1.22"),
		}
		mc := NewMachineContext(sessions)
		require.NotNil(t, mc)
		require.Equal(t, 2, len(mc.Entries))
		require.Equal(t, 1, mc.Labels["s1"])
		require.Equal(t, 2, mc.Labels["s2"])
	})

	t.Run("same hardware different go version", func(t *testing.T) {
		sessions := []*SessionInfo{
			mkSession("s1", &mA, "go1.21"),
			mkSession("s2", &mA, "go1.22"),
		}
		mc := NewMachineContext(sessions)
		require.NotNil(t, mc)
		require.Equal(t, 2, len(mc.Entries))
		require.NotEqual(t, mc.Labels["s1"], mc.Labels["s2"])
	})

	t.Run("session without machine info not labeled", func(t *testing.T) {
		sessions := []*SessionInfo{
			mkSession("s1", &mA, "go1.22"),
			mkSession("s2", &mB, "go1.22"),
			mkSession("s3", nil, ""),
		}
		mc := NewMachineContext(sessions)
		require.NotNil(t, mc)
		require.Equal(t, 0, mc.Labels["s3"])
	})
}
