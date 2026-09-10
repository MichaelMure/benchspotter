Output doc/demo/show-diff.gif

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
# keystrokes below always land on the same session.
Type "rm -rf .benchspotter/recall"
Enter
Type "clear"
Enter
Show

Sleep 1s
Type "benchspotter show diff"
Enter

# ── Session selection (Select) ───────────────────────────────────────────
# Only sessions carrying a stored diff are listed:
#   0. baseline2   1. flate pool   2. rope-0d   3. rope-62   4. rope3
Wait+Screen@60s /Select session/
Sleep 1200ms
Down
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 300ms
Down
Sleep 900ms
Enter

# ── Diff viewport ────────────────────────────────────────────────────────
# 525 lines over 7 files: the working tree as it stood when the benchmarks ran.
Wait+Screen@60s /diff --git/
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
# docproxy/listview.go: a linear ListElements scan replaced by ListTail
Sleep 3s

PageDown
# docproxy/textview.go: ListElements → ListAt / ListLen
Sleep 3s
PageDown
Sleep 3s
PageUp
Sleep 2s

Type "q"
Sleep 800ms
