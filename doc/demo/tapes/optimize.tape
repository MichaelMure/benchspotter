Output doc/demo/optimize.gif

Set Theme "Catppuccin Mocha"
Set FontFamily "Fira Code"
Set FontSize 14
Set Width 800
Set Height 600
Set Framerate 12
Set TypingSpeed 40ms
Set Padding 2

Hide
Type "cd /home/michael/dev/benchspotter"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "./benchspotter optimize"
Enter

# ── Parameter selection (MultiSelect) ─────────────────────────────────────
# Recall pre-selects capacityFactor and fillFraction (BenchmarkMapOps params).
# Params in discovery order; navigate to see the full list.
Sleep 1500ms
Down         # browse
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 300ms
Up           # scroll back up
Sleep 300ms
Up
Sleep 300ms
Up
Sleep 300ms
Enter        # confirm: capacityFactor + fillFraction

# ── Benchmark selection (SingleSelect) ────────────────────────────────────
# Recall pre-selects BenchmarkMapOps (last in the list).
Sleep 1500ms
Up           # peek at previous benchmark
Sleep 400ms
Down         # back to BenchmarkMapOps
Sleep 300ms
Enter        # confirm

# ── Strategy selection (SingleSelect) ─────────────────────────────────────
# StrategyDefs order: 0=random, 1=coord (pre-selected), 2=sa
# Cursor starts on "coord". Press Up once to reach "random".
Sleep 1500ms
Up           # move to "random"
Sleep 400ms
Enter        # confirm

# ── Metric selection (SingleSelect) ──────────────────────────────────────
# Recall pre-selects ns/op.
Sleep 1500ms
Down         # peek at B/op
Sleep 300ms
Up           # back to ns/op
Sleep 300ms
Enter        # confirm

# ── Random optimization running ───────────────────────────────────────────
# Random strategy samples BenchmarkMapOps across the (capacityFactor, fillFraction)
# parameter space. Each trial runs "go test -bench BenchmarkMapOps" once (~1-2s).
Sleep 8s
# Scroll the results table while trials accumulate
Down
Sleep 200ms
Down
Sleep 200ms
Up
Sleep 6s
# More trials arriving
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Sleep 8s
# Stop and quit
Type "s"
Sleep 1s
Type "q"
Sleep 500ms
