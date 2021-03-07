package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckOnesSegment(t *testing.T) {
	t.Run("1001-false", func(t *testing.T) {
		input := "1001"
		got := checkOnesSegment(input)
		assert.False(t, got)
	})

	t.Run("110-true", func(t *testing.T) {
		input := "110"
		got := checkOnesSegment(input)
		assert.True(t, got)
	})
	t.Run("1-true", func(t *testing.T) {
		input := "1"
		got := checkOnesSegment(input)
		assert.True(t, got)
	})

	t.Run("10-true", func(t *testing.T) {
		input := "10"
		got := checkOnesSegment(input)
		assert.True(t, got)
	})
}
