package execenv

import (
	"context"
	"errors"

	"charm.land/huh/v2"
)

var ErrNoPrompt = errors.New("interactive prompts are disabled (--no-prompt)")

// Form is the subset of *huh.Form used by callers.
type Form interface {
	Run() error
	RunWithContext(ctx context.Context) error
}

// ErrForm is a Form that always returns ErrNoPrompt.
type ErrForm huh.Form

func (*ErrForm) Run() error                             { return ErrNoPrompt }
func (*ErrForm) RunWithContext(_ context.Context) error { return ErrNoPrompt }
