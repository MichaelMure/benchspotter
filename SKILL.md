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

## Agent / headless use

Add `--no-prompt` (a persistent root flag) to any command to disable all
interactive prompts. If a prompt would be triggered without a required flag,
the command errors immediately naming the missing flag. Use this whenever
driving benchspotter from a script or agent.

**Working directory**: by default `benchspotter` auto-detects the project root
from the current directory. Use `-C <dir>` / `--dir <dir>` to point it at a
different directory without `cd`.

**`-c` is required for `bench` mode**: when `-p bench` is included and `-c` is
omitted, the count prompt fires. Under `--no-prompt` that becomes an error.
Always pass `-c`.

```bash
benchspotter --no-prompt bench -p bench -n "my session" --bench BenchmarkFoo -c 5
benchspotter --no-prompt session ls --format json
benchspotter --no-prompt compare bench --session <id1> --session <id2> --format json
```

## Running benchmarks

```bash
benchspotter bench -p bench -n "my session" --bench BenchmarkFoo -c 5

# regex filter — OR-ed, repeatable
benchspotter bench -p bench -n "reads" --bench Read -c 5

# all benchmarks / all profiles
benchspotter bench -p bench -n "full sweep" --all-bench -c 5
benchspotter bench --all-profile --all-bench -n "full sweep"

# multiple profile types at once
benchspotter bench -p bench,cpu,mem,escape -n "baseline" --bench BenchmarkFoo -c 5

# -p accepts: bench, cpu, mem, mutex, block, escape, inline
# --bench / -b: regexp, repeatable; --all-bench: every discovered benchmark
# --all-profile: all modes; -n: session name; -c: iteration count (bench only)
```

## Listing and managing sessions

```bash
# list
benchspotter session ls --format json
benchspotter session ls --tag my-tag --format json
benchspotter session ls --bench Foo --format json        # filter by benchmark regexp

# show full metadata for one session (including full note text)
benchspotter session show <id> --format json

# manage
benchspotter session tag    <id> <tag>
benchspotter session untag  <id> <tag>
benchspotter session rename <id> <new-name>
benchspotter session note   <id> "<text>"
benchspotter session rm  -y <id>                         # -y skips confirmation
```

`session ls` and `session show` JSON fields: `id`, `human_name`, `time`,
`commit`, `has_diff`, `profiles`, `tags`, `notes`, `benchmarks`, `machine`,
`go_version`. In `session ls` text mode, `notes` is truncated; use
`session show` when you need the full text.

When sessions originate from different machines, `session ls` flags them.
Use tags to keep machine groups filterable:

```bash
benchspotter session tag <id> ci-runner
benchspotter trend --tag ci-runner --format text
```

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

JSON output: tables → benchmarks → sessions, each with `center`, `range`,
`delta` (% change vs. baseline), and `stats` (significance).

Other useful flags:
- `--alpha 0.05` — significance threshold
- `--filter 'BenchmarkFoo*'` — scope to specific benchmarks (benchstat filter syntax)
- `--col .config` — split columns by benchmark configuration keys
- `--skip-machine-check` — suppress the cross-machine warning
- `--format raw` — raw `.bench` text, suitable for piping into `benchstat`

## Trend over time

```bash
# readable table — default for human/agent use
benchspotter trend --format text
benchspotter trend --tag my-tag --last 20 --format text
benchspotter trend --bench Foo --format text             # error if pattern matches nothing

# machine-consumable nested JSON (benchmarks[] → units[] → points[])
benchspotter trend --format json
```

JSON point fields: `session_id`, `human_name`, `time`, `center`, `lo`, `hi`, `n`.

## Inspecting sessions

Use `benchspotter show/compare` rather than raw `go tool pprof` or `gcflags`
— results land in a session and are comparable across code changes.

### Profiles

```bash
# capture
benchspotter bench -p cpu                               # or mem, mutex, block
benchspotter bench -p mem -c 10                         # mem needs ≥10 iterations;
                                                        # fewer collapses to runtime.mallocgc

# inspect — structured top-N table
benchspotter show cpu   --session <name> --format json --top 30
benchspotter show mem   --session <name> --format json --metric alloc_objects
benchspotter show block --session <name> --format json
benchspotter show mutex --session <name> --format json
# json fields: name, file, line, flat_ns, cum_ns, flat_pct, cum_pct
# mem metrics: alloc_space (default) | alloc_objects | inuse_space | inuse_objects

# compare between two sessions
benchspotter compare cpu   --base <name> --new <name> --bench BenchmarkFoo --format json
benchspotter compare mem   --base <name> --new <name> --bench BenchmarkFoo --format json
benchspotter compare block --base <name> --new <name> --bench BenchmarkFoo --format json
benchspotter compare mutex --base <name> --new <name> --bench BenchmarkFoo --format json
# json fields: name, file, line, base_flat_ns, new_flat_ns, delta_flat_ns, delta_flat_pct, ...

# raw pprof — only when you need interactive go tool pprof exploration
benchspotter show cpu    --session <name> --format raw > cpu.pprof
benchspotter compare cpu --base <name> --new <name> --format raw > diff.pprof  # positive = regression
go tool pprof cpu.pprof
```

### Compiler analysis

```bash
# capture (instead of go build -gcflags='-m=2')
benchspotter bench -p escape
benchspotter bench -p inline

# inspect
benchspotter show escape --session <name> --format json
# json fields: file, line, col, message, heap_escape, leaking_param, flow_chain

benchspotter show inline --session <name> --format json
# json fields: file, line, col, message, kind: "cannot_inline" | "can_inline" | "inlining_call"

# --all: include all compiler notes; --deps: include stdlib and dependencies

# diff between two sessions
benchspotter compare escape --base <name> --new <name> --format json
benchspotter compare inline --base <name> --new <name> --format json
# json fields: status (added|removed), file, base_line, new_line, message, ...
# --all to include unchanged entries
# --format raw: one line per change: "+ file:line:col: msg" / "- file:line:col: msg"
```

### Git diff

```bash
benchspotter show diff --session <name> --format raw    # unified diff
benchspotter show diff --session <name> --format json   # fields: session_id, session_name, diff
```

## Parameter optimization

Requires `benchinput` declarations in the codebase. Add instrumentation with a
one-line change (temporary — revert after tuning):

```diff
+import "github.com/MichaelMure/benchspotter/benchinput"

-const bufSize = 1024
+var bufSize = benchinput.Int("bufSize", 1024, 64, 65536)
```

Supported: `benchinput.Bool`, `Int`, `IntLog`, `Float`, `FloatLog`. Values are
injected via `BENCHSPOTTER_<NAME>` env vars at runtime; `optimize` sets them
automatically. `IntLog` / `FloatLog` sample on a log scale — better for
parameters spanning orders of magnitude.

```bash
benchspotter optimize \
  --bench BenchmarkFoo \   # regexp; error if zero or multiple benchmarks match
  --param bufSize \
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
trials). The correlation matrix maps each parameter to each metric — useful for
deciding which parameters actually matter.

## Scripting tips

- **Pass session names directly** — `--session`, `--base`, and `--new` all
  accept the human name (e.g. `--session "after refactor"`). Only reach for
  the UUID when a name matches multiple sessions (the command will tell you).
- `--format raw` on profile commands yields binary pprof; pipe into `go tool pprof`.
- `--format raw` on `compare bench` yields raw `.bench` text for `benchstat`.
- For `optimize`, use `--max-trials` + `--timeout` together: trials caps total
  work, timeout guards against slow benchmarks that would exceed wall-clock budget.
