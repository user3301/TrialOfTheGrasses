package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestContainsPattern1(t *testing.T) {
	arr, m, k := []int{1, 2, 4, 4, 4, 4}, 1, 3
	expected := true
	actual := ContainasPattern(arr, m, k)
	assert.Equal(t, expected, actual)
}

func TestContainsPattern2(t *testing.T) {
	arr, m, k := []int{1, 2, 1, 2, 1, 1, 1, 3}, 2, 2
	expected := true
	actual := ContainasPattern(arr, m, k)
	assert.Equal(t, expected, actual)
}

func TestContainsPattern3(t *testing.T) {
	arr, m, k := []int{1, 2, 1, 2, 1, 3}, 2, 3
	expected := false
	actual := ContainasPattern(arr, m, k)
	assert.Equal(t, expected, actual)
}

func TestContainsPattern4(t *testing.T) {
	arr, m, k := []int{1, 2, 3, 1, 2}, 2, 2
	expected := false
	actual := ContainasPattern(arr, m, k)
	assert.Equal(t, expected, actual)
}

func TestContainsPattern5(t *testing.T) {
	arr, m, k := []int{2, 2, 2, 2}, 2, 3
	expected := false
	actual := ContainasPattern(arr, m, k)
	assert.Equal(t, expected, actual)
}
