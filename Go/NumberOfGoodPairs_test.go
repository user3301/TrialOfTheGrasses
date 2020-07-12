package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumIdenticalPairs1(t *testing.T) {
	nums := []int{1, 2, 3, 1, 1, 3}
	expected := 4
	actual := NumIdenticalPairs(nums)
	assert.Equal(t, expected, actual)
}

func TestNumIdenticalPairs2(t *testing.T) {
	nums := []int{1, 1, 1, 1}
	expected := 6
	actual := NumIdenticalPairs(nums)
	assert.Equal(t, expected, actual)
}

func TestNumIdenticalPairs3(t *testing.T) {
	nums := []int{1, 2, 3}
	expected := 0
	actual := NumIdenticalPairs(nums)
	assert.Equal(t, expected, actual)
}
