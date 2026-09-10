Output doc/demo/optimize.gif

Set Theme "Catppuccin Mocha"
Set FontFamily "Fira Code"
Set FontSize 21
Set Width 1140
Set Height 880
Set Framerate 12
Set TypingSpeed 40ms
Set Padding 2

Hide
Type "cd /home/michael/dev/benchspotter"
Enter
# Prompts remember the last selection: start from a clean slate so the
# keystrokes below always land on the same options.
Type "rm -rf .benchspotter/recall"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter optimize"
Enter

# ── Parameter selection (MultiSelect) ─────────────────────────────────────
# benchinput declarations found by static analysis, in discovery order:
#   0. name1 [bool]              3. capacityFactor [float 0.25..6]
#   1. name2 [int 0..24]         4. fillFraction [float 0.1..1]
#   2. threshold [float 1..128]  5. name3   6. name4
# Sweeping a single parameter is what unlocks the line chart.
Wait+Screen@60s /Select parameters/
Sleep 1500ms
Down
Sleep 350ms
Down
Sleep 500ms
Space        # threshold — the insertion-sort cutoff of a hybrid quicksort
Sleep 1000ms
Enter

# ── Benchmark selection (Select) ─────────────────────────────────────────
# BenchmarkHybridSort is 13th of 14.
Wait+Screen@60s /Select benchmark/
Sleep 1200ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 130ms
Down
Sleep 1000ms
Enter

# ── Strategy (Select): random, coord, sa ─────────────────────────────────
Wait+Screen@60s /Optimization strategy/
Sleep 1200ms
Down
Sleep 350ms
Down
Sleep 900ms
Enter        # sa — random exploration, then simulated annealing

# ── Metric (Select): ns/op, B/op, allocs/op ──────────────────────────────
Wait+Screen@60s /Metric to optimize/
Sleep 1500ms
Enter        # ns/op

# ── Search running ───────────────────────────────────────────────────────
# Each trial re-runs "go test -bench BenchmarkHybridSort" with
# BENCHSPOTTER_THRESHOLD set, and plots ns/op against the parameter value.
Wait+Screen@120s /trials/
Sleep 8s

# The trials table is scrollable while the search keeps running.
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 2000ms
Up
Sleep 250ms
Up
Sleep 250ms
Up
Sleep 250ms
Up
Sleep 8s

# The curve fills in: expensive at threshold=1, a clear interior minimum,
# rising again as insertion sort takes over larger partitions.
Sleep 6s

# [s] stops the search and leaves the final chart up.
Type "s"
Sleep 4000ms

Type "q"
Sleep 800ms
