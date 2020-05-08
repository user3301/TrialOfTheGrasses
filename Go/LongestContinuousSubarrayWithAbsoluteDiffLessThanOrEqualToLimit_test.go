package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestLongestSubarray1(t *testing.T) {
	arr := []int{8, 2, 4, 7}
	limit := 4
	expected := 2
	actual := LongestSubarray(arr, limit)
	assert.Equal(t, expected, actual)
}

func TestLongestSubarray2(t *testing.T) {
	arr := []int{10, 1, 2, 4, 7, 2}
	limit := 5
	expected := 4
	actual := LongestSubarray(arr, limit)
	assert.Equal(t, expected, actual)
}

func TestLongestSubarray3(t *testing.T) {
	arr := []int{4, 2, 2, 2, 4, 4, 2, 2}
	limit := 0
	expected := 3
	actual := LongestSubarray(arr, limit)
	assert.Equal(t, expected, actual)
}
