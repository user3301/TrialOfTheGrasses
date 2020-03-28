package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestIntersection1(t *testing.T) {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	expected := []int{2}
	actual := Intersection(nums1, nums2)
	for idx, val := range expected {
		assert.Equal(t, val, actual[idx])
	}
}

func TestIntersection2(t *testing.T) {
	nums1 := []int{4, 9, 5}
	nums2 := []int{9, 4, 9, 8, 4}
	expected := []int{9, 4}
	actual := Intersection(nums1, nums2)
	for idx, val := range expected {
		assert.Equal(t, val, actual[idx])
	}
}
