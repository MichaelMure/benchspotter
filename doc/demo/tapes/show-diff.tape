Output doc/demo/show-diff.gif

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
Type "./benchspotter show diff"
Enter

# ── Session selection (SingleSelect) ─────────────────────────────────────
# Only sessions with a stored diff are shown:
#   0. based       ← recall pre-selects this
#   1. with escape
Sleep 1500ms
Down         # peek at "with escape"
Sleep 500ms
Up           # back to "based"
Sleep 300ms
Enter        # select "based"

# ── Diff viewport ─────────────────────────────────────────────────────────
Sleep 1500ms
Down
Sleep 100ms
Down
Sleep 100ms
Down
Sleep 100ms
Down
Sleep 100ms
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 200ms
Down
Sleep 200ms
Down
Sleep 200ms
PageDown
Sleep 500ms
PageDown
Sleep 500ms
PageUp
Sleep 400ms
# Quit
Type "q"
Sleep 500ms
