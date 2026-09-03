<div align="center">
  <h1 align="center">Benchspotter</h1>

  <p>
    <a href="https://github.com/MichaelMure/benchspotter/tags">
        <img alt="GitHub Tag" src="https://img.shields.io/github/v/tag/MichaelMure/benchspotter">
    </a>
    <a href="https://github.com/MichaelMure/benchspotter/actions?query=">
      <img src="https://github.com/MichaelMure/benchspotter/actions/workflows/gotest.yml/badge.svg" alt="Build Status">
    </a>
    <a href="https://github.com/MichaelMure/benchspotter/blob/master/LICENSE">
        <img alt="MIT License" src="https://img.shields.io/badge/License-MIT-green">
    </a>
  </p>
</div>

Benchspotter is a companion tool for Go performance work.

Golang provides an awesome set of tools for benchmarking, but it's often a hassle to keep track of the results, analyse, extract insights, and share them. Instead, `benchspotter` runs and stores results, handles the boring parts and gives you higher-level tooling.

While every command gives interactive prompts for humans, LLMs/agents can set any parameters through flags and get JSON in return. This turns `benchspotter` into a cross-session backend for agent-assisted performance work.

## Install

```bash
go install github.com/MichaelMure/benchspotter@latest
```

Or pick a pre-compiled binary from a [release](https://github.com/MichaelMure/benchspotter/releases).

## Quickstart

```bash
benchspotter bench              # run benchmarks/profiles and save a session
benchspotter compare bench      # compare sessions with benchstat
benchspotter compare cpu        # diff CPU profiles between two sessions
benchspotter compare mem        # diff memory profiles between two sessions
benchspotter compare mutex      # diff mutex profiles between two sessions
benchspotter compare block      # diff block profiles between two sessions
benchspotter compare escape     # diff escape analysis between two sessions
benchspotter compare inline     # diff inlining decisions between two sessions
benchspotter trend              # view performance over time
benchspotter show diff          # view code changes if any
benchspotter show cpu           # view hot functions from a CPU profile
benchspotter show mem           # view allocation from a memory profile
benchspotter show mutex         # mutex contention hotspots
benchspotter show block         # view blocking contention hotspots
benchspotter show escape        # view compiler escape analysis
benchspotter show inline        # view compiler inlining decisions
benchspotter session ls         # list all sessions
benchspotter session show       # show full metadata for a session (name, commit, tags, notes, …)
benchspotter optimize           # search for optimal tuning constants
```

## Scripting and agent use

All interactive prompts can be bypassed with flags, and output can be emitted as JSON, which makes the tool easy to drive from scripts or AI agents. You can use [SKILL.md](SKILL.md) to teach your agent how to use `benchspotter`.

## Cross-machine sharing

Benchspotter writes its results to `.benchspotter`. You can safely share those with your team through your VCS or ignore it. Benchspotter is machine-aware and will let you know when results from multiple machines are used in the same context.

This can help you, for example, to figure out performance issue across architecture or deployments.

## Features

### Run benchmarks and profiles

Select benchmarks, choose which profiles to capture alongside (CPU, memory, mutex, block, ...), name the session, pick iteration count. Selections are remembered between runs. The current `git diff` is stored with the session, so results are always traceable back to the exact code that produced them.

![bench](doc/demo/bench.gif)

### Compare results

Select sessions from a list and get a `benchstat` table with confidence intervals, significance tests, and geomeans across all of them. More than two sessions works fine.

![compare](doc/demo/compare.gif)

### Observe trends

Regressions that happen one small commit at a time are easy to miss when you're only ever comparing two sessions. The trend view shows all benchmarks across all sessions at once, making drift visible. Select a benchmark to get a time-series chart and a full per-session table. Or just pat yourself on the back when looking at the results of your hard work.

![trend](doc/demo/trend.gif)

### Inspect the code changes

Benchmarks are rarely run on a clean commit, so each session remembers the `git diff` at the time of running the analysis, so you can see exactly what code has been running. You can inspect the diff with `benchspotter show diff`.

Those diffs are also used when showing [escape or inline analysis](#escape--inline-analysis).

![show diff](doc/demo/show-diff.gif)

### Profiles

Browse CPU, memory, mutex, and block profiles for any session directly in the terminal. The flat/cumulative table covers most cases quickly; toggle source annotations when you need to see exactly which lines are hot. No `.out` files to locate, no browser to open.

`compare cpu/mem/mutex/block` diffs two sessions' profiles against each other, highlighting which functions grew or shrank.

![show cpu](doc/demo/show-cpu.gif)

### Escape & inline analysis

Inspect which variables the compiler moved to the heap and which functions it refused to inline. Source context is shown inline when the session has a diff or a clean commit.

`compare escape` and `compare inline` diff two sessions side by side: entries are labelled **added**, **removed**, or **same** (with line-shift detection when the code moved but didn't change). This makes it easy to see whether a refactor accidentally introduced a new heap allocation or broke an inlining opportunity.

![show escape](doc/demo/show-escape.gif)

### Optimize magic values

There is often in programs a series of magic numbers (buffer sizes, pre-allocation, concurrency limits, batch sizes, thresholds…) whose optimal value is not obvious, and are a pain to tune by hand.

After minimal and temporary instrumentation, Benchspotter can discover and run an optimization algorithm on those values, based on their influence on benchmarks. Here is how you can do this instrumentation:

```diff
+import "benchspotter/benchinput"

-const MAGIC_NUMBER = 1000
+var MAGIC_NUMBER = benchinput.Int("magic number", 1000, 50, 10_000)
```

Supported instrumentation are `Bool`, `Int`, `IntLog`, `Float`, and `FloatLog`.

To help you understand your program's behavior, Benchspotter computes the general correlation (=influence) of those inputs on the performance and shows a graph of results. It's often possible to see hidden threshold effects, in particular due to allocations.

![optimize](doc/demo/optimize.gif)

Warning: the computed "best" values are only as good as your benchmarks. Almost always, your real workload will have different patterns. Don't treat those values as ultimate truth, they can only inform your decisions.

## License

MIT



