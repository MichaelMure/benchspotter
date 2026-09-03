Output demo/bench.gif

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
Type "./benchspotter bench"
Enter

# ── Profile selection (MultiSelect) ──────────────────────────────────────
# Recall pre-selects "Benchmark" (index 0). Cursor starts there.
Sleep 1500ms
Down         # peek at CPU profile option
Sleep 400ms
Down         # peek at Memory
Sleep 300ms
Up           # go back up
Sleep 300ms
Up           # back to Benchmark (pre-selected)
Sleep 300ms
Enter        # confirm: only Benchmark

# ── Benchmark selection (MultiSelect) ────────────────────────────────────
# Recall pre-selects BenchmarkFoo and BenchmarkBar.
Sleep 1500ms
Down         # browse the list
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
Enter        # confirm pre-selected BenchmarkFoo + BenchmarkBar

# ── Session name ──────────────────────────────────────────────────────────
Sleep 800ms
Type "after fix"
Enter

# ── Iteration count ───────────────────────────────────────────────────────
Sleep 500ms
Type "5"
Enter

# ── Benchmarks running ────────────────────────────────────────────────────
Sleep 15s
