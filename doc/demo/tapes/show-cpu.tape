Output doc/demo/show-cpu.gif

Set Theme "Catppuccin Mocha"
Set FontFamily "Fira Code"
Set FontSize 21
Set Width 1140
Set Height 880
Set Framerate 12
Set TypingSpeed 40ms
Set Padding 2

Hide
Type "cd /home/michael/dev/gotomerge"
Enter
# Prompts remember the last selection: start from a clean slate so the
# keystrokes below always land on the same session and benchmark.
Type "rm -rf .benchspotter/recall"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter show cpu"
Enter

# ── Session selection (Select) ───────────────────────────────────────────
# Only sessions holding a cpu profile are listed, chronologically:
#   0. baseline2            3. a8de0a5 column/rle: optimize allocations
#   1. rope-62              4. perf improvements
#   2. 83b99ef format: ...  5. readers by value
Wait+Screen@60s /Select session/
Sleep 1200ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 800ms
Enter

# ── Benchmark selection (Select) ─────────────────────────────────────────
# Listed per package with their total cpu time:
#   0. BenchmarkInt64Reader  1. BenchmarkStringReader  2. BenchmarkUint64Reader
Wait+Screen@60s /BenchmarkUint64Reader/
Sleep 1200ms
Down
Sleep 300ms
Down
Sleep 800ms
Enter

# ── Hot functions ────────────────────────────────────────────────────────
Wait+Screen@60s /FUNCTION/
Sleep 3s

# [a] annotates each row with the source at the recorded line.
Type "a"
Sleep 3500ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 250ms
Down
Sleep 3s

# Toggle it back off to return to the compact table.
Type "a"
Sleep 2s

# [s] cycles the sort order: flat → cumulative → name.
Type "s"
Sleep 2500ms
Type "s"
Sleep 2500ms
Type "s"
Sleep 1500ms

Type "q"
Sleep 800ms
