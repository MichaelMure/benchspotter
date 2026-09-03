---
name: benchspotter
description: Run, store and analyse Go benchmarks — compare runs, track trends over time, read cpu/mem/mutex/block profiles, diff compiler escape and inlining decisions, and search for optimal tuning constants. Use whenever working on Go performance: measuring a change, chasing a regression, or hunting allocations.
---

## Sessions

Every `benchspotter bench` run records a **session** under `.benchspotter/`: the
benchmark results, whichever profiles were captured, the git commit, the uncommitted
diff, and machine metadata. Sessions accumulate over weeks and across machines, so any
two runs stay comparable. Every analysis command reads sessions; none of them re-run
anything.

Prefer `benchspotter` over raw `go test -bench`, `go tool pprof` or `-gcflags=-m`: the
output lands in a session and becomes comparable against every other run.

## Start here

Run `benchspotter --no-prompt session ls` before anything else. It shows what has
already been recorded — names, commits, which profiles each session carries, which
machine produced it — so you find an existing baseline instead of spending minutes
re-measuring one. If it reports no sessions, nothing has been recorded yet.

**Record the baseline before touching the code.** A session captures the git commit and
the uncommitted diff at the moment it runs. Once you have edited, the "before" state is
gone and no later command can reconstruct it. This is the only mistake here that cannot
be undone.

## Driving it headlessly

Pass `--no-prompt` (a root flag) on every invocation. Any prompt that would have fired
becomes an immediate error naming the flag to supply instead, so nothing ever hangs.
`-C <dir>` runs against another directory without `cd`.

Under `--no-prompt`, `bench` requires:

| flag | |
|---|---|
| `-p/--profile` (or `--all-profile`) | what to capture: `bench,cpu,mem,mutex,block,escape,inline` |
| `-b/--bench` (or `--all-bench`) | which benchmarks; regexp, repeatable, OR-ed |
| `-n/--name` | session name; `""` is valid and means unnamed |
| `-c/--count` | **only** when `-p` includes `bench` |

## Naming a session

`bench -n` names the session, and that name is what every later command takes:

```bash
benchspotter --no-prompt bench -p bench -b BenchmarkFoo -n "baseline" -c 10
benchspotter --no-prompt compare bench --session "baseline" --session "after"
```

`--session`, `--base` and `--new` all accept a UUID, the name shown by `session ls`, or
the name passed to `-n`. Reusing a name is safe: `session ls` disambiguates it
(`baseline-3c`), and an ambiguous lookup fails with the alternatives spelled out.

`--print-id` makes `bench` print only the UUID on stdout — progress goes to stderr, so
`ID=$(benchspotter ... --print-id)` is clean. Text format only.

## Choosing an output format

The rule: **text, unless the command renders source code.** Text output is a table, and
JSON costs 3–6× the tokens to say the same thing — `session ls`, `trend`,
`compare bench` and every `show`/`compare` of a cpu, mem, mutex or block profile are all
smaller and just as precise as text.

The exception is `escape` and `inline`. In text mode those annotate the source line by
line, which is enormous: on a small project `show escape` is 864 KB as text against
459 KB as JSON, and `compare escape` is 24 KB against 1.5 KB. Use `-f json` for all four
of `show escape|inline` and `compare escape|inline`.

Narrow `show escape` and `show inline` with `--func`, using function names taken from
`show cpu` — that is what turns compiler output into something actionable.

Use `-f raw` only to pipe into another tool: it gives binary pprof from the profile
commands, `.bench` text from `compare bench` (feed it to `benchstat`), a unified diff
from `show diff`, and one line per change from `compare escape|inline`.

## Avoid

- Running `go test -bench`, `go tool pprof` or `go build -gcflags=-m` directly. Nothing
  is persisted and nothing becomes comparable.
- Re-running a benchmark to look at it again. Sessions are permanent; read the stored
  one with `show`.
- `--web` on any profile command. It opens a browser and blocks.
- `show escape` or `show inline` without `--func`. Unfiltered they dump the whole
  project, ~175 KB of JSON.

## Recipes

### Did my change help?

```bash
# 1. before editing anything
benchspotter --no-prompt bench -p bench,cpu,mem --all-bench -n "before" -c 10

# 2. make the change

# 3. after
benchspotter --no-prompt bench -p bench,cpu,mem --all-bench -n "after" -c 10
benchspotter --no-prompt compare bench --session "before" --session "after"
```

Capture `cpu` and `mem` on the baseline even when you only want timings. Each profile
type costs one extra benchmark run, but `compare cpu` refuses to run when either side
lacks the profile — and by the time you want it, the old code is gone. Add
`escape,inline` too if compiler behaviour is in question; those only compile, so they
are nearly free.

Use `-c 10` or more. A result that stays `~` at a high count is an answer — the change
is below the benchmark's noise floor — not a reason to keep re-running.

### Where did the time go?

`compare bench` says a benchmark regressed; `compare cpu` says which function did.

```bash
benchspotter --no-prompt compare cpu --base "before" --new "after" --bench BenchmarkFoo
```

Positive deltas are the regression. Sort with `--sort` and widen with `--top` if the
culprit is not in the default 20.

### Where are the allocations?

```bash
benchspotter --no-prompt compare mem --base "before" --new "after" --bench BenchmarkFoo
benchspotter --no-prompt show mem --session "after" --bench BenchmarkFoo --metric alloc_objects
```

A memory profile needs `-c 10` or more at capture time; with fewer iterations everything
collapses into `runtime.mallocgc`.

Then ask why the compiler put it on the heap, using the function names the profile gave
you:

```bash
benchspotter --no-prompt show escape --session "after" --func myHotFunc -f json
```

### Is a hot function being inlined?

```bash
benchspotter --no-prompt show inline --session "after" --func myHotFunc -f json
```

`cannot_inline` sites carry the compiler's reason, e.g. `function too complex: cost 122
exceeds budget 80`. To see what a change did to inlining across the board:

```bash
benchspotter --no-prompt compare inline --base "before" --new "after" -f json
```

### Tune a constant

Requires a `benchinput` declaration in the code — a one-line temporary change, reverted
after tuning:

```diff
+import "github.com/MichaelMure/benchspotter/benchinput"

-const bufSize = 1024
+var bufSize = benchinput.Int("bufSize", 1024, 64, 65536)
```

```bash
benchspotter --no-prompt optimize -b BenchmarkFoo --param bufSize \
  -m ns/op -s random --max-trials 50 --timeout 5m -f json
```

`Bool`, `Int`, `IntLog`, `Float`, `FloatLog` are supported; the `Log` variants sample on
a log scale. `--param` is repeatable. `-s` picks `random`, `coord` (coordinate descent)
or `sa` (simulated annealing) — prefer `random` when you want the correlation matrix,
since the others sample too narrowly for it to mean anything. `--maximize` flips the
objective, `--count` sets iterations per trial. Always pair `--max-trials` with
`--timeout`.

JSON: `benchmark`, `strategy`, `metric`, `minimize`, `trials[]` (`n`, `params`,
`metrics`), `best`, and a `correlation` matrix once ≥3 trials have run — that matrix
tells you which parameters actually matter.

## Command reference

### bench

```bash
benchspotter bench -p bench,cpu,mem -b BenchmarkFoo -b BenchmarkBar -n "baseline" -c 10
benchspotter bench --all-profile --all-bench -n "full sweep" -c 10
```

### session

```bash
benchspotter session ls                       # add --tag <t> or --bench <regexp> to filter
benchspotter session show <id-or-name>        # full metadata, including the whole note
benchspotter session tag    <id-or-name> <tag>
benchspotter session untag  <id-or-name> <tag>
benchspotter session rename <id-or-name> <new-name>
benchspotter session note   <id-or-name> "<text>"
benchspotter session rm -y  <id-or-name>
```

JSON fields: `id`, `name`, `human_name`, `time`, `commit`, `has_diff`, `profiles`,
`tags`, `notes`, `benchmarks`, `machine`, `go_version`. `name` is what was passed to
`-n` (absent if unnamed); `human_name` is the de-duplicated display name. In `ls` text
mode `notes` is truncated — use `session show` for the full text.

Two markers are benchspotter's own, and appear in `trend` and `compare bench` tables
too: `±` after a commit hash means the session ran with uncommitted changes (`show diff`
retrieves them), and `⚙1`/`⚙2` mark different machines, with a legend under the table.
Timings across machines are not comparable — say so rather than reporting the delta. Tag
machine groups to keep them filterable: `session tag <id> ci-runner`, then
`trend --tag ci-runner`.

### compare bench

```bash
benchspotter compare bench --session "before" --session "after"
benchspotter compare bench --session <a> --session <b> --session <c> --baseline <b>
```

First `--session` is the baseline unless `--baseline` says otherwise. Also:
`--alpha 0.05` (significance), `--table/--row/--col` (projections; defaults
`.config`/`.fullname`/`.file`), `--skip-machine-check`.

`--filter` takes benchfilter syntax — `key:value`, never a bare pattern. The name
carries no `Benchmark` prefix, and `/…/` is a regexp: `--filter '.name:Foo'` for one
benchmark, `--filter '.name:/Foo/'` for every name containing `Foo`.

JSON is an array of `{unit, benchmarks[]}` — one entry per unit (`sec/op`, `B/op`,
`allocs/op`). Each benchmark holds `sessions[]` of `{name, center, range}`, and every
non-baseline session adds `delta` and `stats`. `delta` and `range` are display
**strings** (`"~"`, `"+2.31%"`, `"∞"`), not numbers.

### trend

```bash
benchspotter trend                            # every session, one column each
benchspotter trend --tag ci-runner --last 20
benchspotter trend --bench Foo                # errors if the regexp matches nothing
benchspotter trend --comparison baseline      # arrows vs first session, not vs previous
```

Also `--session` (repeatable) and `--confidence` (default 0.95).

JSON is `{sessions: [{i, name, time}], benchmarks: [{name, metrics: {unit: [...]}}]}`.
The metric arrays are **formatted display strings** (`"5.92µs"`) positionally aligned
with `sessions[]` — no easier to compute with than the text table, and 5× larger.

### show / compare profiles

```bash
benchspotter show cpu   --session "after" --bench BenchmarkFoo --top 30
benchspotter show mem   --session "after" --bench BenchmarkFoo --metric alloc_objects
benchspotter compare cpu --base "before" --new "after" --bench BenchmarkFoo
```

`--bench` is required whenever the session holds more than one profile of that kind;
omitting it fails with the list of available benchmarks. `show` and `compare` both take
`--top` and `--sort`. `--metric` exists only on the `mem` commands: `alloc_space`
(default), `alloc_objects`, `inuse_space`, `inuse_objects`.

JSON: `show` gives `name`, `file`, `line`, `flat_ns`, `cum_ns`, `flat_pct`, `cum_pct`;
`compare` gives `base_flat_ns`, `new_flat_ns`, `delta_flat_ns`, `delta_flat_pct`, ….

```bash
benchspotter show cpu    --session "after" --bench BenchmarkFoo -f raw > cpu.pprof
benchspotter compare cpu --base "before" --new "after" --bench BenchmarkFoo \
  -f raw > diff.pprof                                          # positive = regression
go tool pprof cpu.pprof
```

### show / compare escape and inline

```bash
benchspotter show escape    --session "after" --func myFunc -f json
benchspotter compare inline --base "before" --new "after" -f json
```

`--func` is repeatable, OR-ed, case-insensitive; it matches the message for `escape` and
the function name for `inline`.

`show escape` JSON: `{summary: {heap_escape, leaking_param, other}, files: [{file,
sites: [{line, kind, message}]}]}`. Only `heap_escape` and `leaking_param` are shown by
default; `--all` adds the rest.

`show inline` JSON: `{summary: {cannot_inline, inlining_call, can_inline}, files:
[{file, sites: [{line, kind, function, reason}]}]}`. Only `cannot_inline` by default;
`--all` adds the rest.

`compare` replaces the summary with `{added, removed, same}` and adds `status`
(`"added"`/`"removed"`/`"same"`) plus `base_line` to each site; `--all` includes
unchanged sites.

In both, project code only unless `--deps`, except that `heap_escape`, `leaking_param`
and `cannot_inline` are always reported from dependencies too.

### show diff

```bash
benchspotter show diff --session "after" -f raw     # unified diff
benchspotter show diff --session "after" -f json    # session_id, session_name, diff
```
