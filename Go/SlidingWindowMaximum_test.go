package gosoln

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	nums := []int{1,3,-1,-3,5,3,6,7}
	k := 3
	want :=[]int{3,3,5,5,6,7}
	got := MaxSlidingWindow(nums,k)
	assert.ElementsMatch(t, want, got)
}