package execenv

import (
	"context"
	"errors"
	"fmt"
)

var ErrNoPrompt = errors.New("interactive prompts are disabled (--no-prompt)")

// Form is the subset of *huh.Form used by callers.
type Form interface {
	Run() error
	RunWithContext(ctx context.Context) error
}

// errForm is a Form that always returns ErrNoPrompt, naming the missing flag.
type errForm struct{ flag string }

func (f *errForm) Run() error { return fmt.Errorf("%w: use %s", ErrNoPrompt, f.flag) }
func (f *errForm) RunWithContext(_ context.Context) error {
	return fmt.Errorf("%w: use %s", ErrNoPrompt, f.flag)
}
