package engine

// DiffEscapeSite pairs base and new escape sites matched by (file, message).
// Exactly one of Base or New is nil when the status is Added or Removed.
type DiffEscapeSite struct {
	Base *EscapeSite // nil when DiffAdded
	New  *EscapeSite // nil when DiffRemoved
}

// Status returns the diff status for this site.
func (d DiffEscapeSite) Status() DiffStatus {
	if d.Base == nil {
		return DiffAdded
	}
	if d.New == nil {
		return DiffRemoved
	}
	return DiffSame
}

// Site returns the canonical site: New if available, Base otherwise.
func (d DiffEscapeSite) Site() EscapeSite {
	if d.New != nil {
		return *d.New
	}
	return *d.Base
}

// DiffEscapeAnalysis diffs base and new escape analysis sites.
// Sites are matched by (file, message) to be stable against line-number shifts.
func DiffEscapeAnalysis(base, new []EscapeSite) []DiffEscapeSite {
	return diffAnalysis(base, new,
		func(s EscapeSite) string { return s.File + "\x00" + normMsgKey(s.Message) },
		func(b, n *EscapeSite) DiffEscapeSite { return DiffEscapeSite{Base: b, New: n} },
		func(b *EscapeSite) DiffEscapeSite { return DiffEscapeSite{Base: b} },
		func(n *EscapeSite) DiffEscapeSite { return DiffEscapeSite{New: n} },
	)
}
