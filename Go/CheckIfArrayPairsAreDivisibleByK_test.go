package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestCanArrange1(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 10, 6, 7, 8, 9}
	k := 5
	expected := true
	actual := CanArrange(arr, k)
	assert.Equal(t, expected, actual)
}

func TestCanArrange2(t *testing.T) {
	arr := []int{-1, 1, -2, 2, -3, 3, -4, 4}
	k := 3
	expected := true
	actual := CanArrange(arr, k)
	assert.Equal(t, expected, actual)
}

func TestCanArrange3(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6}
	k := 10
	expected := false
	actual := CanArrange(arr, k)
	assert.Equal(t, expected, actual)
}
