package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplyUnifiedDiff_basic(t *testing.T) {
	lines := []string{"a", "b", "c", "d", "e"}
	diff := []byte(`diff --git a/f.go b/f.go
--- a/f.go
+++ b/f.go
@@ -2,3 +2,3 @@
 b
-c
+C
 d
`)
	got := ApplyUnifiedDiff(lines, diff, "f.go")
	assert.Equal(t, []string{"a", "b", "C", "d", "e"}, got)
}

func TestApplyUnifiedDiff_addition(t *testing.T) {
	lines := []string{"a", "b", "c"}
	diff := []byte(`diff --git a/f.go b/f.go
--- a/f.go
+++ b/f.go
@@ -2,1 +2,2 @@
 b
+b2
`)
	got := ApplyUnifiedDiff(lines, diff, "f.go")
	assert.Equal(t, []string{"a", "b", "b2", "c"}, got)
}

func TestApplyUnifiedDiff_deletion(t *testing.T) {
	lines := []string{"a", "b", "c", "d"}
	diff := []byte(`diff --git a/f.go b/f.go
--- a/f.go
+++ b/f.go
@@ -2,2 +2,1 @@
-b
 c
`)
	got := ApplyUnifiedDiff(lines, diff, "f.go")
	assert.Equal(t, []string{"a", "c", "d"}, got)
}

func TestApplyUnifiedDiff_noMatch(t *testing.T) {
	lines := []string{"a", "b", "c"}
	diff := []byte(`diff --git a/other.go b/other.go
--- a/other.go
+++ b/other.go
@@ -1,1 +1,1 @@
-a
+A
`)
	got := ApplyUnifiedDiff(lines, diff, "f.go")
	assert.Equal(t, lines, got)
}

func TestApplyUnifiedDiff_multipleHunks(t *testing.T) {
	lines := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	diff := []byte(`diff --git a/f.go b/f.go
--- a/f.go
+++ b/f.go
@@ -2,1 +2,1 @@
-2
+TWO
@@ -8,1 +8,1 @@
-8
+EIGHT
`)
	got := ApplyUnifiedDiff(lines, diff, "f.go")
	assert.Equal(t, []string{"1", "TWO", "3", "4", "5", "6", "7", "EIGHT", "9", "10"}, got)
}

func TestApplyUnifiedDiff_emptyDiff(t *testing.T) {
	lines := []string{"a", "b"}
	got := ApplyUnifiedDiff(lines, nil, "f.go")
	assert.Equal(t, lines, got)
}
