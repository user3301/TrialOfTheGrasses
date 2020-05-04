package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestCheckIfAll1sAreAtLeastLengthKPlaceAway1(t *testing.T) {
	nums := []int{1, 0, 0, 0, 1, 0, 0, 1}
	k := 2
	actual := kLengthApart(nums, k)
	assert.Equal(t, true, actual)
}

func TestCheckIfAll1sAreAtLeastLengthKPlaceAway2(t *testing.T) {
	nums := []int{1, 0, 0, 1, 0, 1}
	k := 2
	actual := kLengthApart(nums, k)
	assert.Equal(t, false, actual)
}

func TestCheckIfAll1sAreAtLeastLengthKPlaceAway3(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	k := 0
	actual := kLengthApart(nums, k)
	assert.Equal(t, true, actual)
}

func TestCheckIfAll1sAreAtLeastLengthKPlaceAway4(t *testing.T) {
	nums := []int{0, 1, 0, 1}
	k := 1
	actual := kLengthApart(nums, k)
	assert.Equal(t, true, actual)
}
