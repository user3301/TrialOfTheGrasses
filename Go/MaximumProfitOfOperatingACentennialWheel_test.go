package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinOperationsMaxProfit1(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name                      string
		customers                 []int
		boardingCost, runningCost int
		want                      int
	}{
		//{
		//	name:         "test1",
		//	customers:    []int{10, 10, 6, 4, 7},
		//	boardingCost: 3,
		//	runningCost:  8,
		//	want:         9,
		//},
		{
			name:         "test2",
			customers:    []int{10, 10, 1, 0, 0},
			boardingCost: 4,
			runningCost:  4,
			want:         5,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := MinOperationsMaxProfit(tc.customers, tc.boardingCost, tc.runningCost)
			assert.Equal(t, tc.want, got)
		})
	}
}
