package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestReorderSpaces1(t *testing.T) {
	t.Parallel()
	text := "  this   is  a sentence "
	want := "this   is   a   sentence"
	got := ReorderSpaces(text)
	assert.Equal(t, want, got)
}

func TestReorderSpaces2(t *testing.T) {
	t.Parallel()
	text := " practice   makes   perfect"
	want := "practice   makes   perfect "
	got := ReorderSpaces(text)
	assert.Equal(t, want, got)
}

func TestReorderSpaces3(t *testing.T) {
	t.Parallel()
	text := "a"
	want := "a"
	got := ReorderSpaces(text)
	assert.Equal(t, want, got)
}
