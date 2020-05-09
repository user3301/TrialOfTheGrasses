package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 2, 2, 3}
	val := 3
	expected := 2
	actual := RemoveElement(nums, val)
	assert.Equal(t, expected, actual)

}
