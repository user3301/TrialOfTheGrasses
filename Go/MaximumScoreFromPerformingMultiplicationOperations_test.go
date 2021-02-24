package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumOperations(t *testing.T) {
	nums := []int{1, 2, 3}
	multiplers := []int{3, 2, 1}
	want := 14
	got := maximumScore(nums, multiplers)
	assert.Equal(t, want, got)
}
