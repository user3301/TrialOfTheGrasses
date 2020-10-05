package gosoln

import (
	"testing"

	"github.com/google/go-cmp/cmp"

	"gotest.tools/assert"
)

func TestSubsets2(t *testing.T) {
	trans := cmp.Transformer("IgnoreCapacity", func(in []int) []int {
		out := append([]int(nil), in...)
		return out
	})
	input := []int{1, 2, 2}
	want := make([][]int, 6, 8)
	want[0] = []int{}
	want[1] = []int{1}
	want[2] = []int{1, 2}
	want[3] = []int{1, 2, 2}
	want[4] = []int{2}
	want[5] = []int{2, 2}
	//want := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	got := SubsetsWithDup(input)
	for i, v := range want {
		assert.DeepEqual(t, v, got[i], trans)
	}
}
