package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumberOfMatches(t *testing.T) {
	t.Run("n=7", func(t *testing.T) {
		n := 7
		want := 6
		got := NumberOfMatches(n)
		assert.Equal(t, want, got)
	})

	t.Run("n=14", func(t *testing.T) {
		n := 14
		want := 13
		got := NumberOfMatches(n)
		assert.Equal(t, want, got)
	})
}
