package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestIntersectionOfTwoArraysII(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2, 2}
	actual := Intersect(nums1, nums2)
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}
