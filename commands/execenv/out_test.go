package execenv

import (
	"bytes"
	"testing"

	"charm.land/lipgloss/v2"
	"github.com/charmbracelet/colorprofile"
	"github.com/stretchr/testify/require"
)

// newTestOut builds an out writing to buf. An empty environ plus a non-file
// writer makes colorprofile detect no terminal, as when output is piped.
func newTestOut(buf *bytes.Buffer) out {
	return out{out: colorprofile.NewWriter(buf, []string{})}
}

func TestOutStripsStylingWhenNotTerminal(t *testing.T) {
	styled := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("10")).Render("hello")
	require.NotEqual(t, "hello", styled, "lipgloss should have emitted escape sequences")

	t.Run("Printf", func(t *testing.T) {
		var buf bytes.Buffer
		newTestOut(&buf).Printf("%s\n", styled)
		require.Equal(t, "hello\n", buf.String())
	})

	t.Run("Println", func(t *testing.T) {
		var buf bytes.Buffer
		newTestOut(&buf).Println(styled)
		require.Equal(t, "hello\n", buf.String())
	})

	// Tables and diffs are rendered through the io.Writer interface.
	t.Run("Write", func(t *testing.T) {
		var buf bytes.Buffer
		o := newTestOut(&buf)
		n, err := o.Write([]byte(styled))
		require.NoError(t, err)
		require.Equal(t, len(styled), n, "Write must report the bytes it was given")
		require.Equal(t, "hello", buf.String())
	})
}

// Raw() is how the --format raw profile commands emit binary pprof; stripping
// ANSI sequences out of protobuf would silently corrupt it.
func TestOutRawBypassesStripping(t *testing.T) {
	var buf bytes.Buffer
	o := newTestOut(&buf)

	binary := []byte{0x1f, 0x8b, 0x08, 0x00, 0x1b, '[', '3', '1', 'm', 0x00, 0xff}
	n, err := o.Raw().Write(binary)
	require.NoError(t, err)
	require.Equal(t, len(binary), n)
	require.Equal(t, binary, buf.Bytes())
}
