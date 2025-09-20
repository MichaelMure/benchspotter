package execenv

import (
	"io"
	"os"

	"github.com/mattn/go-isatty"
)

type In interface {
	io.Reader

	Raw() io.Reader

	// IsTerminal tells if the input is a user terminal (rather than a buffer,
	// a pipe ...), which tells if we can use interactive features.
	IsTerminal() bool
}

type in struct {
	io.Reader
}

func (i in) Raw() io.Reader {
	return i.Reader
}

func (i in) IsTerminal() bool {
	if f, ok := i.Reader.(*os.File); ok {
		return isTerminal(f)
	}
	return false
}

func isTerminal(file *os.File) bool {
	return isatty.IsTerminal(file.Fd()) || isatty.IsCygwinTerminal(file.Fd())
}
