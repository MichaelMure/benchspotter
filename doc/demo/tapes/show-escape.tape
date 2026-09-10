Output doc/demo/show-escape.gif

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
Type "benchspotter show escape"
Enter

# ── Session selection (Select) ───────────────────────────────────────────
# Only sessions carrying compiler analysis are listed:
#   0. 83b99ef format: ...        3. readers by value
#   1. a8de0a5 column/rle: ...    4. baseline   ← recorded at the current HEAD,
#   2. perf improvements                          so the annotated source matches
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

# ── Escape analysis ──────────────────────────────────────────────────────
# 1071 heap escapes across 332 functions. The sidebar lists every function with
# escapes; the pane annotates the source at each site.
Wait+Screen@60s /heap escape/
Sleep 2500ms

# [Tab] walks from one escape site to the next, sidebar tracking along.
Tab
Sleep 500ms
Tab
Sleep 500ms
Tab
Sleep 500ms
Tab
Sleep 1500ms

# [f] unfolds the dataflow chain that explains how each value reached the heap.
Type "f"
Sleep 3000ms
Tab
Sleep 800ms
Tab
Sleep 800ms
Tab
Sleep 800ms
Tab
Sleep 800ms
Tab
Sleep 2000ms

# Sweep through the rest: leaking params, &T{...} escapes to heap, one file
# after another.
PageDown
Sleep 900ms
PageDown
Sleep 900ms
PageDown
Sleep 900ms
PageDown
Sleep 1500ms
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
Sleep 1200ms
PageDown
Sleep 900ms
PageDown
Sleep 900ms
PageDown
Sleep 2000ms

Type "q"
Sleep 800ms
