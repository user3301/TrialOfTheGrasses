package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestRemoveDuplicatesFromSortedArray1(t *testing.T) {
	nums := []int{1, 1, 2}
	expected := 2
	expectedArr := []int{1, 2}
	RemoveDuplicates(nums)

	for i := 0; i < expected; i++ {
		assert.Equal(t, expectedArr[i], nums[i])
	}
}
