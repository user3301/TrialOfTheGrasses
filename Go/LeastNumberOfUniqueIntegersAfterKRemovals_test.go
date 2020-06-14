package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestFindLeastNumOfUniqueInts1(t *testing.T) {
	arr := []int{5, 5, 4}
	k := 1
	expected := 1
	actual := FindLeastNumOfUniqueInts(arr, k)
	assert.Equal(t, expected, actual)
}

func TestFindLeastNumOfUniqueInts2(t *testing.T) {
	arr := []int{2, 1, 1, 3, 3, 3}
	k := 3
	expected := 1
	actual := FindLeastNumOfUniqueInts(arr, k)
	assert.Equal(t, expected, actual)
}
