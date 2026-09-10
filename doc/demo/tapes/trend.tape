Output doc/demo/trend.gif

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
Type "./benchspotter trend"
Enter

# ── Overview: benchmark list with session columns ─────────────────────────
Sleep 1500ms

# Navigate down the benchmark list
Down
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 300ms

# Cycle unit
Type "m"
Sleep 400ms
Type "m"
Sleep 400ms

# Open detail view for selected benchmark
Enter
Sleep 1s

# ── Detail view: time-series chart + table ────────────────────────────────
# Cycle unit in detail view
Type "m"
Sleep 400ms

# Step through benchmarks
Right
Sleep 600ms
Right
Sleep 600ms

# Go back to overview
Type "b"
Sleep 500ms

# Quit
Type "q"
Sleep 500ms
