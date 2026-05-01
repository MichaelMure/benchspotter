---
name: benchspotter
description: benchspotter is a performance measurement backend CLI tool. It runs Go benchmarks, persists the results as named **sessions**, and provides tools for comparison, trend analysis, profiling, compiler analysis, and parameter optimization. Think of it as a long-lived performance logbook: sessions accumulate across weeks or months, across machines, and across code changes.
---

## Core concept: sessions

Every `benchspotter bench` run creates a session. A session can contain any
combination of:

- **bench results** — raw `go test -bench` output (ns/op, B/op, allocs/op, …)
- **CPU profile** — where wall-clock time is spent
- **memory profile** — heap allocation sites
- **mutex / block profiles** — lock contention and goroutine blocking
- **escape analysis** — which variables the compiler moved to the heap
- **inlining decisions** — which functions the compiler refused to inline
- **git diff snapshot** — the uncommitted diff at the time of the run

Sessions have a human name, a git commit hash, machine metadata, and
optionally tags.

## Running benchmarks

All interactive prompts can be bypassed with flags:

```bash
# bench results only
benchspotter bench -p bench -n "my session" -b BenchmarkFoo -c 5

# multiple profile types at once
benchspotter bench -p bench,cpu,mem,escape -n "baseline" -b BenchmarkFoo

# -p accepts: bench, cpu, mem, mutex, block, escape, inline
# -b accepts a benchmark name (repeatable)
# -n sets the session name
# -c sets the iteration count (bench mode only)
```

## Listing and managing sessions

```bash
# list all sessions
benchspotter session ls --format json

# filter by tag
benchspotter session ls --tag my-tag --format json

# filter by benchmark name
benchspotter session ls --benchmark BenchmarkFoo --format json

# tag / untag / rename / delete (all accept id + value as args)
benchspotter session tag    <id> <tag>
benchspotter session untag  <id> <tag>
benchspotter session rename <id> <new-name>
benchspotter session rm  -y <id>   # -y skips confirmation
```

The JSON output of `session ls` includes: `id`, `human_name`, `time`,
`commit`, `has_diff`, `profiles`, `tags`, `benchmarks`, `machine`, `go_version`.

`--session` flags throughout benchspotter accept either a UUID session ID or a
human name. Names that match multiple sessions return an error — use the ID in
that case.

## Comparing benchmark results

```bash
benchspotter compare bench \
  --session "before refactor" \
  --session "after refactor" \
  --format json
```

The first `--session` is the baseline unless `--baseline` overrides it:

```bash
benchspotter compare bench \
  --session <id1> --session <id2> --session <id3> \
  --baseline <id2> \
  --format json
```

The JSON output contains tables → benchmarks → sessions, each with `center`,
`range`, `delta` (% change vs. baseline), and `stats` (significance).

Other useful flags:
- `--alpha 0.05` — significance threshold
- `--filter 'BenchmarkFoo*'` — scope to specific benchmarks (benchstat filter syntax)
- `--col .config` — split columns by benchmark configuration keys
- `--skip-machine-check` — suppress the cross-machine warning
- `--format raw` — raw `.bench` text, suitable for piping into `benchstat`

## Trend over time

```bash
# all sessions
benchspotter trend --format json

# scoped to specific sessions
benchspotter trend --session <id1> --session <id2> --format json

# scoped by tag and count
benchspotter trend --tag my-tag --last 20 --format json
```

JSON output: `benchmarks[]` → `units[]` → `points[]`, each point with
`session_id`, `human_name`, `time`, `center`, `lo`, `hi`, `n`.

## Inspecting profiles

```bash
# structured top-N table
benchspotter show cpu   --session <id> --format json --top 30
benchspotter show mem   --session <id> --format json --metric alloc_objects
benchspotter show block --session <id> --format json
benchspotter show mutex --session <id> --format json
# json fields: name, file, line, flat_ns, cum_ns, flat_pct, cum_pct

# raw pprof binary — pipe into go tool pprof
benchspotter show cpu --session <id> --bench BenchmarkFoo --format raw > cpu.pprof
go tool pprof cpu.pprof

# mem metrics: alloc_space (default) | alloc_objects | inuse_space | inuse_objects
```

## Compiler analysis

```bash
# escape analysis
benchspotter show escape --session <id> --format json
# json fields: file, line, col, message, heap_escape, leaking_param, flow_chain

# inlining decisions
benchspotter show inline --session <id> --format json
# json fields: file, line, col, message, kind
# kind: "cannot_inline" | "can_inline" | "inlining_call"

# --all includes all compiler notes (not just actionable ones)
# --deps includes stdlib and dependencies
```

## Git diff snapshot

```bash
# raw unified diff
benchspotter show diff --session <id> --format raw

# structured
benchspotter show diff --session <id> --format json
# json fields: session_id, session_name, diff
```

## Parameter optimization

Requires `benchinput` declarations in the codebase. Each parameter has a name,
a default value, and a valid range — these represent tuning knobs (buffer
sizes, concurrency limits, thresholds, etc.) whose optimal value is not obvious.

For headless use, supply all flags and use `--format json`. Stop the run with
`--max-trials` (count-based) or `--timeout` (time-based), or both:

```bash
benchspotter optimize \
  --bench BenchmarkFoo \
  --param size \
  --metric ns/op \
  --strategy random \
  --max-trials 50 \
  --timeout 5m \
  --format json
```

Three strategies:
- `random` — uniform sampling; unbiased, reliable correlation matrix.
- `coord` — coordinate descent; fast for smooth surfaces, stops on convergence.
- `sa` — simulated annealing; better than `coord` at escaping local optima.

Use `--maximize` to flip the objective. `--count` sets benchmark iterations
per trial (increase for noisy benchmarks).

JSON output: `benchmark`, `strategy`, `metric`, `minimize`, `trials[]` (each
with `n`, `params`, `metrics`), `best`, and a `correlation` matrix (once ≥3
trials have run). The correlation matrix maps each parameter to each metric —
useful for deciding which parameters actually matter.

## Cross-machine work

Sessions record the host's CPU model, OS, and Go version. `session ls --format json`
exposes this in the `machine` field. `compare bench` warns when mixing sessions
from different machines; use `--skip-machine-check` if intentional.

Tag sessions by machine to keep groups distinct and filterable:

```bash
benchspotter session tag <id> ci-runner
benchspotter trend --tag ci-runner --format json
```

## Scripting tips

- Use `session ls --format json` to discover session IDs and filter by tag,
  benchmark, or machine before passing IDs to other commands.
- `--format raw` on profile commands yields binary pprof; pipe into `go tool pprof`.
- `--format raw` on `compare bench` yields raw `.bench` text for `benchstat`.
- `--session` flags accept human names; use the UUID if the name is ambiguous.
- For `optimize`, use `--max-trials` + `--timeout` together: trials caps total
  work, timeout guards against slow benchmarks that would exceed wall-clock budget.
