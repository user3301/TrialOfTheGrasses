package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestKDiffPairsInAnArray1(t *testing.T) {
	input, k := []int{3, 1, 4, 1, 5}, 2
	expected := 2
	actual := FindPairs(input, k)
	assert.Equal(t, expected, actual)
}

func TestKDiffPairsInAnArray2(t *testing.T) {
	input, k := []int{1, 2, 3, 4, 5}, 1
	expected := 4
	actual := FindPairs(input, k)
	assert.Equal(t, expected, actual)
}

func TestKDiffPairsInAnArray3(t *testing.T) {
	input, k := []int{1, 3, 1, 5, 4}, 0
	expected := 1
	actual := FindPairs(input, k)
	assert.Equal(t, expected, actual)
}

func TestKDiffPairsInAnArray4(t *testing.T) {
	input, k := []int{1, 2, 3, 4, 5}, 2
	expected := 3
	actual := FindPairs(input, k)
	assert.Equal(t, expected, actual)
}
