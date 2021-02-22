package gosoln

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestMinOperationsBoxes(t *testing.T) {
	testCases := []struct {
		boxes string
		want  []int
	}{
		{
			boxes: "110",
			want:  []int{1, 1, 3},
		},
		{
			boxes: "001011",
			want:  []int{11, 8, 5, 4, 3, 4},
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%dth test case", i+1), func(t *testing.T) {
			got := minOperations(tc.boxes)
			assert.DeepEqual(t, tc.want, got)
		})
	}
}
