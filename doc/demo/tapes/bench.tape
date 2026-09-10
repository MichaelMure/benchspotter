Output doc/demo/bench.gif

Set Theme "Catppuccin Mocha"
Set FontFamily "Fira Code"
Set FontSize 14
Set Width 800
Set Height 600
Set Framerate 12
Set TypingSpeed 40ms
Set Padding 2

Hide
Type "cd /home/michael/dev/gotomerge"
Enter
# Prompts remember the last selection: start from a clean slate so the
# keystrokes below always land on the same options.
Type "rm -rf .benchspotter/recall"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter bench"
Enter

# ── Profile selection (MultiSelect) ──────────────────────────────────────
# Options: Benchmark, CPU, Memory, Mutex, Blocking, Escape analysis, Inlining.
# Nothing is pre-selected, so pick Benchmark + Memory + Escape analysis.
Sleep 1500ms
Space        # Benchmark
Sleep 500ms
Down
Sleep 250ms
Down
Sleep 250ms
Space        # Memory
Sleep 500ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Space        # Escape analysis
Sleep 800ms
# The list is taller than its viewport, so scroll back to the top: all three
# ticked options are then on screen when the form is submitted.
Up
Sleep 150ms
Up
Sleep 150ms
Up
Sleep 150ms
Up
Sleep 150ms
Up
Sleep 800ms
Enter

# ── Benchmark selection (MultiSelect) ────────────────────────────────────
# Take the first four so the demo stays short.
Sleep 1500ms
Space        # BenchmarkReadMetadata
Sleep 400ms
Down
Sleep 200ms
Space        # BenchmarkWriteMetadata
Sleep 400ms
Down
Sleep 200ms
Space        # BenchmarkUint64Reader
Sleep 400ms
Down
Sleep 200ms
Space        # BenchmarkInt64Reader
Sleep 800ms
Enter

# ── Session name ──────────────────────────────────────────────────────────
Sleep 800ms
Type "baseline"
Enter

# ── Iteration count ───────────────────────────────────────────────────────
Sleep 500ms
Type "3"
Enter

# ── Benchmarks, memory profiles, then the compiler analysis ──────────────
# The "escape" line is the very last thing printed, so wait for it rather
# than guessing a duration.
Wait+Screen@300s /escape +[0-9]/
Sleep 3s
