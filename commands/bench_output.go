package commands

import (
	"context"
	"fmt"
	"math"
	"strings"
	"time"

	"charm.land/lipgloss/v2"
	"golang.org/x/perf/benchfmt"

	"benchspotter/commands/execenv"
)

// benchPrinter streams benchmark results with in-place updates on terminals.
// In terminal mode with count > 1, each benchmark group is shown on one line
// that updates as iterations complete, finishing with mean ± spread.
// In non-terminal mode or count == 1, it prints one line per completed group.
// Column order: [counter]  kind  name  values  [elapsed]
type benchPrinter struct {
	env        *execenv.Env
	total      int
	nameWidth  int
	kindWidth  int
	isTerminal bool

	name    string
	n       int
	sums    []float64
	mins    []float64
	maxs    []float64
	units   []string
	hasLine bool // partial line on screen; next write should start with \r\033[2K
	start   time.Time
}

func newBenchPrinter(env *execenv.Env, count, nameWidth, kindWidth int) *benchPrinter {
	return &benchPrinter{
		env:        env,
		total:      count,
		nameWidth:  nameWidth,
		kindWidth:  kindWidth,
		isTerminal: env.Err.IsTerminal(),
	}
}

// Next fetches the next benchmark result. On a terminal with count > 1, it
// calls it directly and lets Add handle in-place progress updates. Otherwise
// it wraps the call in a spinner (which is silent on non-terminals).
func (p *benchPrinter) Next(ctx context.Context, it func() (*benchfmt.Result, error)) (*benchfmt.Result, error) {
	if p.isTerminal && p.total > 1 {
		return it()
	}
	var res *benchfmt.Result
	err := p.env.Spinner().Title("Running benchmarks...").ActionWithErr(func(ctx context.Context) error {
		var e error
		res, e = it()
		return e
	}).Context(ctx).Run()
	return res, err
}

// Add records one benchmark result and updates the display.
func (p *benchPrinter) Add(res *benchfmt.Result) {
	name := string(res.Name)

	if name != p.name {
		p.finalise()
		p.name = name
		p.n = 0
		p.sums = make([]float64, len(res.Values))
		p.mins = make([]float64, len(res.Values))
		p.maxs = make([]float64, len(res.Values))
		p.units = make([]string, len(res.Values))
		for i, v := range res.Values {
			p.mins[i] = math.MaxFloat64
			p.maxs[i] = -math.MaxFloat64
			u := v.OrigUnit
			if u == "" {
				u = v.Unit
			}
			p.units[i] = u
		}
		p.hasLine = false
		p.start = time.Now()
	}

	p.n++
	for i, v := range res.Values {
		val := v.OrigValue
		if v.OrigUnit == "" {
			val = v.Value
		}
		p.sums[i] += val
		if val < p.mins[i] {
			p.mins[i] = val
		}
		if val > p.maxs[i] {
			p.maxs[i] = val
		}
	}

	if p.n == p.total {
		p.finalise()
	} else if p.isTerminal {
		prefix := ""
		if p.hasLine {
			prefix = "\r\033[2K"
		}
		p.env.Err.Printf("%s  %s  %s  %s  %s",
			prefix, p.fmtCounter(p.n), p.fmtKind("bench"), p.fmtName(), p.fmtValues(false))
		p.hasLine = true
	} else {
		p.env.Err.Printf("  %s  %s  %s  %s\n",
			p.fmtCounter(p.n), p.fmtKind("bench"), p.fmtName(), p.fmtValues(false))
	}
}

// Flush finalises any in-progress group. Call when the iterator is exhausted.
func (p *benchPrinter) Flush() {
	p.finalise()
	p.name = ""
}

func (p *benchPrinter) finalise() {
	if p.n == 0 {
		return
	}
	elapsed := time.Since(p.start).Truncate(10 * time.Millisecond)
	prefix := ""
	if p.isTerminal && p.hasLine {
		prefix = "\r\033[2K"
	}
	p.env.Err.Printf("%s  %s  %s  %s  %s    %s\n",
		prefix, p.fmtCounter(p.n), p.fmtKind("bench"), p.fmtName(),
		p.fmtValues(p.n > 1), p.env.Style.TonedDown(elapsed.String()))
	p.hasLine = false
	p.n = 0
}

// fmtCounter returns a fixed-width toned-down "[n/total]", padded so later
// columns stay at the same position regardless of n.
func (p *benchPrinter) fmtCounter(n int) string {
	maxW := len(fmt.Sprintf("[%d/%d]", p.total, p.total))
	s := fmt.Sprintf("[%d/%d]", n, p.total)
	return p.env.Style.TonedDown(s) + strings.Repeat(" ", max(0, maxW-len(s)))
}

// fmtKind returns the kind label styled and padded to kindWidth.
func (p *benchPrinter) fmtKind(kind string) string {
	styled := p.env.Style.Accent(kind)
	return styled + strings.Repeat(" ", max(0, p.kindWidth-len(kind)))
}

// fmtName returns the current benchmark name styled and padded to nameWidth.
func (p *benchPrinter) fmtName() string {
	styled := p.env.Style.Bold(displayBenchName(p.name))
	return styled + strings.Repeat(" ", max(0, p.nameWidth-lipgloss.Width(styled)))
}

func (p *benchPrinter) fmtValues(withSpread bool) string {
	var parts []string
	for i, unit := range p.units {
		avg := p.sums[i] / float64(p.n)
		if withSpread && avg != 0 && p.maxs[i] > p.mins[i] {
			pct := (p.maxs[i]-p.mins[i]) / (2 * avg) * 100
			spread := fmt.Sprintf("±%.0f%%", pct)
			if pct > 10 {
				spread = p.env.Style.Warning(spread)
			} else {
				spread = p.env.Style.TonedDown(spread)
			}
			parts = append(parts, fmt.Sprintf("%s %s %s", fmtBenchNum(avg), unit, spread))
		} else {
			parts = append(parts, fmt.Sprintf("%s %s", fmtBenchNum(avg), unit))
		}
	}
	return strings.Join(parts, "   ")
}

// profilePrinter prints progress lines for profile capture.
// Column order: [counter]  kind  name  elapsed
type profilePrinter struct {
	env       *execenv.Env
	kind      string
	total     int
	nameWidth int
	kindWidth int
	n         int
}

func newProfilePrinter(env *execenv.Env, kind string, total, nameWidth, kindWidth int) *profilePrinter {
	return &profilePrinter{
		env:       env,
		kind:      kind,
		total:     total,
		nameWidth: nameWidth,
		kindWidth: kindWidth,
	}
}

func (p *profilePrinter) Done(benchName string, elapsed time.Duration) {
	p.n++
	elapsed = elapsed.Truncate(10 * time.Millisecond)

	maxCounterW := len(fmt.Sprintf("[%d/%d]", p.total, p.total))
	rawCounter := fmt.Sprintf("[%d/%d]", p.n, p.total)
	counter := p.env.Style.TonedDown(rawCounter) + strings.Repeat(" ", max(0, maxCounterW-len(rawCounter)))

	kind := p.env.Style.Accent(p.kind) + strings.Repeat(" ", max(0, p.kindWidth-len(p.kind)))

	name := p.env.Style.Bold(displayBenchName(benchName))
	namePad := strings.Repeat(" ", max(0, p.nameWidth-lipgloss.Width(name)))

	p.env.Err.Printf("  %s  %s  %s%s  %s\n",
		counter, kind, name, namePad, p.env.Style.TonedDown(elapsed.String()))
}

// displayBenchName strips the "Benchmark" prefix and GOMAXPROCS suffix ("-N").
func displayBenchName(name string) string {
	s := strings.TrimPrefix(name, "Benchmark")
	if i := strings.LastIndex(s, "-"); i >= 0 {
		rest := s[i+1:]
		allDigits := len(rest) > 0
		for _, c := range rest {
			if c < '0' || c > '9' {
				allDigits = false
				break
			}
		}
		if allDigits {
			s = s[:i]
		}
	}
	return s
}

// fmtBenchNum formats a float with thousands separators and appropriate precision.
func fmtBenchNum(f float64) string {
	if f < 1 {
		return fmt.Sprintf("%.3f", f)
	}
	if f < 10 {
		return fmt.Sprintf("%.2f", f)
	}
	if f < 100 {
		return fmt.Sprintf("%.1f", f)
	}
	n := int64(math.Round(f))
	s := fmt.Sprintf("%d", n)
	if len(s) <= 3 {
		return s
	}
	var b strings.Builder
	offset := len(s) % 3
	for i, c := range s {
		if i > 0 && (i-offset)%3 == 0 {
			b.WriteByte(',')
		}
		b.WriteRune(c)
	}
	return b.String()
}
