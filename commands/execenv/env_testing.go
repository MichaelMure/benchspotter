package execenv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"

	"charm.land/huh/v2"

	"github.com/MichaelMure/benchspotter/repository"
)

var _ In = &TestIn{}

type TestIn struct {
	*bytes.Buffer
	forceIsTerminal bool
}

func (t *TestIn) Raw() io.Reader {
	return t.Buffer
}

func (t *TestIn) IsTerminal() bool {
	return t.forceIsTerminal
}

var _ Out = &TestOut{}

type TestOut struct {
	*bytes.Buffer
	forceIsTerminal bool
}

func (te *TestOut) Printf(format string, a ...interface{}) {
	_, _ = fmt.Fprintf(te.Buffer, format, a...)
}

func (te *TestOut) Print(a ...interface{}) {
	_, _ = fmt.Fprint(te.Buffer, a...)
}

func (te *TestOut) Println(a ...interface{}) {
	_, _ = fmt.Fprintln(te.Buffer, a...)
}

func (te *TestOut) PrintJSON(v interface{}) error {
	raw, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		return err
	}
	te.Println(string(raw))
	return nil
}

func (te *TestOut) IsTerminal() bool {
	return te.forceIsTerminal
}

func (te *TestOut) Width() int {
	return 80
}

func (te *TestOut) Raw() io.Writer {
	return te.Buffer
}

func TestStyle() Style {
	return NewStyle(huh.ThemeCharm, false)
}

func NewTestEnv(ctx context.Context, repo *repository.Repository) *Env {
	return &Env{
		Ctx:    ctx,
		Repo:   repo,
		In:     &TestIn{Buffer: &bytes.Buffer{}},
		Out:    &TestOut{Buffer: &bytes.Buffer{}},
		Err:    &TestOut{Buffer: &bytes.Buffer{}},
		Style:  TestStyle(),
		Format: FormatText,
	}
}
