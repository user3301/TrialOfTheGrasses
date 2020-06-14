package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestRunningSum(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	expected := []int{1, 3, 6, 10}
	actual := RunningSum(nums)
	assert.DeepEqual(t, expected, actual)
}
