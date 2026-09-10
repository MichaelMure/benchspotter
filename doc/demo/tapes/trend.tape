Output doc/demo/trend.gif

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
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter trend"
Enter

# ── Overview: benchmarks down, sessions across ───────────────────────────
# 88 benchmarks over 12 sessions. The oldest columns are mostly empty — those
# runs predate half the benchmarks — so scroll right to where the data is.
Wait+Screen@60s /Trend/
Sleep 2500ms
Right
Sleep 350ms
Right
Sleep 350ms
Right
Sleep 350ms
Right
Sleep 350ms
Right
Sleep 350ms
Right
Sleep 2000ms

# Down into Map/build and friends: red and green in every column.
PageDown
Sleep 500ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 180ms
Down
Sleep 3000ms

# [c] switches the arrows from "vs the previous session" to "vs the first".
Type "c"
Sleep 3000ms
Type "c"
Sleep 1500ms

# ── Detail view ──────────────────────────────────────────────────────────
# Jump to Uint64Reader, which ran in every session.
PageDown
Sleep 400ms
PageDown
Sleep 400ms
Up
Sleep 250ms
Up
Sleep 250ms
Up
Sleep 250ms
Up
Sleep 250ms
Up
Sleep 1500ms
Enter

# A time series: flat through April, a spike at 83b99ef, then the drop to
# 151.6ns at "readers by value" and on down to 124.0ns.
Sleep 4000ms

# [m] redraws the chart for the other metrics.
Type "m"
Sleep 3000ms
Type "m"
# allocs/op: 10 per op, up to 16, then down to 1.
Sleep 3500ms
Type "m"
Sleep 1500ms

# [b] back to the table.
Type "b"
Sleep 2000ms

Type "q"
Sleep 800ms
