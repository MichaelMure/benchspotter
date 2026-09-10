Output doc/demo/compare.gif

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
# keystrokes below always land on the same sessions.
Type "rm -rf .benchspotter/recall"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter compare bench"
Enter

# ── Session selection (MultiSelect, / to filter) ──────────────────────────
# Sessions are listed chronologically; the two to compare are 9th and 10th:
#   8. perf improvements  (8f47b64)
#   9. readers by value   (cf0af76)
Wait+Screen@60s /Select sessions/
Sleep 1200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 500ms
Space        # perf improvements — baseline
Sleep 600ms
Down
Sleep 300ms
Space        # readers by value — compared against the baseline
Sleep 1s
Enter

# ── Comparison table viewport ─────────────────────────────────────────────
# Scroll down 26 lines, which walks the column and column/rle packages: their
# rows are narrow enough to fit. Stop there — docproxy's row labels
# (Map/build/increasing_put/ops=10000-12) run past the viewport, which
# truncates rather than wraps.
Wait+Screen@60s /geomean/
# Hold on the first table: three regressions, +10.46% on WriteMetadata.
Sleep 3500ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
# column/rle sec/op: -51.80%, -52.31%, -29.07%
Sleep 3s
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
# ... paid for in B/op: +40.35%, +40.35%, +62.50%
Sleep 2500ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
# ... and won back in allocs/op: -90.00%, -90.00%, -66.67%
Sleep 3500ms
Type "q"
Sleep 800ms
