Output doc/demo/show-escape.gif

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
Type "./benchspotter show escape"
Enter

# ── Session selection (SingleSelect) ─────────────────────────────────────
# Sessions with escape analysis:
#   0. with escape  ← recall pre-selects this
#   1. after escape
Sleep 1500ms
Down         # peek at "after escape"
Sleep 500ms
Up           # back to "with escape"
Sleep 300ms
Enter        # select "with escape"

# ── Escape analysis viewport ──────────────────────────────────────────────
# Default: heap escapes + leaking params only, project code only
Sleep 1500ms

# Scroll through escape sites
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 150ms
Down
Sleep 150ms
PageDown
Sleep 500ms

# Toggle to show all compiler notes
Type "a"
Sleep 700ms
Down
Sleep 150ms
Down
Sleep 150ms
PageDown
Sleep 400ms

# Toggle back to heap+leaking only
Type "a"
Sleep 500ms

# Include stdlib + deps
Type "p"
Sleep 700ms
PageDown
Sleep 500ms
PageDown
Sleep 400ms

# Back to project-only
Type "p"
Sleep 400ms

# Quit
Type "q"
Sleep 500ms
