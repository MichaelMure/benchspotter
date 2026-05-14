package engine

import "regexp"

// numberRE matches decimal digit sequences used to strip variable cost/count
// tokens from compiler messages before key comparison.
var numberRE = regexp.MustCompile(`[0-9]+`)

// normMsgKey returns a normalized version of msg with all digit sequences
// replaced by "#", so messages that differ only in cost or count values
// (e.g. "cost 1088" vs "cost 1253") compare as equal.
func normMsgKey(msg string) string {
	return numberRE.ReplaceAllString(msg, "#")
}

// DiffStatus classifies a diagnostic site relative to the base session.
type DiffStatus int

const (
	DiffAdded   DiffStatus = iota // in new session only
	DiffRemoved                   // in base session only
	DiffSame                      // in both sessions
)

// DiffInlineSite pairs base and new inline sites matched by (file, message).
// Exactly one of Base or New is nil when the status is Added or Removed.
type DiffInlineSite struct {
	Base *InlineSite // nil when DiffAdded
	New  *InlineSite // nil when DiffRemoved
}

// Status returns the diff status for this site.
func (d DiffInlineSite) Status() DiffStatus {
	if d.Base == nil {
		return DiffAdded
	}
	if d.New == nil {
		return DiffRemoved
	}
	return DiffSame
}

// Site returns the canonical site: New if available, Base otherwise.
func (d DiffInlineSite) Site() InlineSite {
	if d.New != nil {
		return *d.New
	}
	return *d.Base
}

// DiffInlineAnalysis diffs base and new inline analysis sites.
// Sites are matched by (file, message) to be stable against line-number shifts
// caused by unrelated code edits between sessions.
func DiffInlineAnalysis(base, new []InlineSite) []DiffInlineSite {
	return diffAnalysis(base, new,
		func(s InlineSite) string { return s.File + "\x00" + normMsgKey(s.Message) },
		func(b, n *InlineSite) DiffInlineSite { return DiffInlineSite{Base: b, New: n} },
		func(b *InlineSite) DiffInlineSite { return DiffInlineSite{Base: b} },
		func(n *InlineSite) DiffInlineSite { return DiffInlineSite{New: n} },
	)
}

// diffAnalysis matches sites from base and new by key, returning diff pairs.
// Multiple sites with the same key are matched in order (greedy).
func diffAnalysis[S, D any](
	base, new []S,
	keyOf func(S) string,
	same func(*S, *S) D,
	removed func(*S) D,
	added func(*S) D,
) []D {
	basePool := make(map[string][]int)
	for i, s := range base {
		k := keyOf(s)
		basePool[k] = append(basePool[k], i)
	}

	consumed := make(map[string]int)
	var out []D

	for i := range new {
		k := keyOf(new[i])
		pool := basePool[k]
		if idx := consumed[k]; idx < len(pool) {
			consumed[k]++
			out = append(out, same(&base[pool[idx]], &new[i]))
		} else {
			out = append(out, added(&new[i]))
		}
	}

	for k, pool := range basePool {
		for _, bi := range pool[consumed[k]:] {
			out = append(out, removed(&base[bi]))
		}
	}

	return out
}
