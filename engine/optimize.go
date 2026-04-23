package engine

import (
	"context"
	"os"
	"strconv"

	"golang.org/x/perf/benchfmt"
	"golang.org/x/perf/benchmath"
	"golang.org/x/sys/execabs"

	"benchspotter/benchinput"
)

// ParamAssignment pins one input to a specific value.
// Value is the env-var string encoding passed to go test.
// FloatVal is the native numeric representation used by optimization algorithms.
type ParamAssignment struct {
	Input    InputInfo
	Value    string
	FloatVal float64
}

// Candidate is a complete parameter assignment for one benchmark trial.
type Candidate []ParamAssignment

// EvaluatedPoint is a candidate with its measured results.
type EvaluatedPoint struct {
	Candidate Candidate
	Metrics   map[string]float64 // unit → median across count runs
	Err       error
}

// Strategy decides which candidates to evaluate next.
// Next and Observe are always called from the same goroutine; no locking needed.
type Strategy interface {
	Name() string
	// Next returns the next candidate to evaluate. A nil return means the
	// strategy is exhausted or waiting for more Observe calls before it can
	// continue. In the sequential runner, nil always means exhausted.
	Next() Candidate
	// Observe records the result of a completed trial.
	Observe(EvaluatedPoint)
}

// BestPoint returns the index of the best result in results, and whether one was found.
func BestPoint(results []EvaluatedPoint, metric string, minimize bool) (int, bool) {
	bestIdx := -1
	var bestVal float64
	for i, p := range results {
		if p.Err != nil {
			continue
		}
		v, ok := p.Metrics[metric]
		if !ok {
			continue
		}
		if bestIdx == -1 || (minimize && v < bestVal) || (!minimize && v > bestVal) {
			bestIdx = i
			bestVal = v
		}
	}
	return bestIdx, bestIdx >= 0
}

// runPoint executes one benchmark trial with the candidate's values injected as env vars.
func runPoint(ctx context.Context, bench BenchInfo, candidate Candidate, count int) EvaluatedPoint {
	env := os.Environ()
	for _, a := range candidate {
		env = append(env, benchinput.EnvVarName(a.Input.Name())+"="+a.Value)
	}

	cmd := execabs.CommandContext(ctx, "go", "test",
		"-bench", bench.Regex(),
		"-benchmem", "-count", strconv.Itoa(count), "-run", "^$", ".")
	cmd.Dir = bench.Package
	cmd.Env = env

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return EvaluatedPoint{Candidate: candidate, Err: err}
	}
	if err := cmd.Start(); err != nil {
		return EvaluatedPoint{Candidate: candidate, Err: err}
	}

	accum := make(map[string][]float64)
	r := benchfmt.NewReader(stdout, "")
	for r.Scan() {
		res, ok := r.Result().(*benchfmt.Result)
		if !ok {
			continue
		}
		for _, v := range res.Values {
			unit, value := origUnitValue(v)
			accum[unit] = append(accum[unit], value)
		}
	}
	rerr := r.Err()
	werr := cmd.Wait()
	if rerr != nil {
		return EvaluatedPoint{Candidate: candidate, Err: rerr}
	}
	if werr != nil {
		return EvaluatedPoint{Candidate: candidate, Err: werr}
	}

	thresholds := benchmath.DefaultThresholds
	metrics := make(map[string]float64, len(accum))
	for unit, vals := range accum {
		sample := benchmath.NewSample(vals, &thresholds)
		sum := benchmath.AssumeNothing.Summary(sample, 0.95)
		metrics[unit] = sum.Center
	}
	return EvaluatedPoint{Candidate: candidate, Metrics: metrics}
}

// RunOptimize drives the optimization loop. It returns a channel that receives
// each evaluated point and is closed when the strategy is exhausted, maxTrials
// is reached (0 = unlimited), or the context is cancelled.
func RunOptimize(ctx context.Context, bench BenchInfo, strategy Strategy, count, maxTrials int) <-chan EvaluatedPoint {
	out := make(chan EvaluatedPoint)
	go func() {
		defer close(out)
		for trials := 0; maxTrials == 0 || trials < maxTrials; trials++ {
			c := strategy.Next()
			if c == nil {
				return
			}
			p := runPoint(ctx, bench, c, count)
			if ctx.Err() != nil {
				return
			}
			strategy.Observe(p)
			select {
			case out <- p:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}
