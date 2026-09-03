Output demo/show-cpu.gif

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
Type "./benchspotter show cpu"
Enter

# ── Session selection (SingleSelect) ─────────────────────────────────────
# Sessions with CPU profiles (chronological):
#   0. based       ← recall pre-selects this
#   1. with escape
#   2. after escape
#   3. (unnamed)
Sleep 1500ms
Down         # peek at "with escape"
Sleep 400ms
Up           # back to "based"
Sleep 300ms
Enter        # select "based"

# ── Benchmark selection (SingleSelect) ────────────────────────────────────
# Benchmarks in session order from meta.json:
#   0. BenchmarkBaz   (anotherpackage)
#   1. BenchmarkBoz   (anotherpackage)
#   2. BenchmarkFoo   ← recall pre-selects this
#   3. BenchmarkBar
#   ...
Sleep 1500ms
Down         # peek at BenchmarkBar
Sleep 400ms
Up           # back to BenchmarkFoo
Sleep 300ms
Enter        # select BenchmarkFoo

# ── CPU profile table ─────────────────────────────────────────────────────
Sleep 1500ms

# Scroll down through top functions
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 300ms

# Toggle source annotations
Type "a"
Sleep 800ms

# Scroll through annotated view
Down
Sleep 200ms
Down
Sleep 200ms
PageDown
Sleep 600ms

# Cycle sort: flat → cumulative → name → flat
Type "s"
Sleep 500ms
Type "s"
Sleep 500ms
Type "s"
Sleep 400ms

# Quit
Type "q"
Sleep 500ms
