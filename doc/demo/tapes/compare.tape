Output demo/compare.gif

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
Type "./benchspotter compare stat"
Enter

# ── Session selection (MultiSelect, / to filter) ──────────────────────────
# Recall pre-selects "based", "with escape", "after escape".
# Sessions in chronological order:
#   0. based        (pre-selected)
#   1. with escape  (pre-selected)
#   2. after escape (pre-selected)
#   3. (unnamed)
#   4. after fix    (newest, not pre-selected)
Sleep 1500ms
Down         # browse to "with escape"
Sleep 300ms
Down         # browse to "after escape"
Sleep 300ms
Down         # browse to unnamed
Sleep 300ms
Down         # browse to "after fix" — toggle it in too
Space
Sleep 300ms
Up           # scroll back up
Sleep 300ms
Up
Sleep 300ms
Up
Sleep 300ms
Up
Sleep 300ms
Enter        # confirm selection (4 sessions)

# ── Comparison table viewport ─────────────────────────────────────────────
Sleep 2s
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 300ms
Down
Sleep 300ms
Up
Sleep 200ms
Up
Sleep 300ms
# Quit
Type "q"
Sleep 500ms
