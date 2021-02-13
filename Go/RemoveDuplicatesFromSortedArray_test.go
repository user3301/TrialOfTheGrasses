package gosoln

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	testCases := []struct {
		nums          []int
		expected      int
		expectedArray []int
	}{
		{
			nums:          []int{1, 1, 2},
			expected:      2,
			expectedArray: []int{1, 2},
		},
		{
			nums:          []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected:      5,
			expectedArray: []int{0, 1, 2, 3, 4},
		},
	}

	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("TestCase%d", i), func(t *testing.T) {
			RemoveDuplicates(tc.nums)
			for i := 0; i < tc.expected; i++ {
				assert.Equal(t, tc.expectedArray[i], tc.nums[i])
			}
		})
	}
}
