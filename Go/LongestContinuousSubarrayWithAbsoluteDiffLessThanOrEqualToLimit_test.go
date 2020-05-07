package gosoln

import(
	"testing"

	"gotest.tools/assert"
)
func TestLongestSubarray(t *testing.T) {
	arr := []int {8,2,4,7}
	limit := 4
	expected := 2
	actual := LongestSubarray(arr, limit)
	assert.Equal(t, expected, actual)
}
