package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinOperations(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		logs []string
		want int
	}{
		{
			name: "test1",
			logs: []string{"d1/", "d2/", "../", "d21/", "./"},
			want: 2,
		},
		{
			name: "test2",
			logs: []string{"d1/", "d2/", "./", "d3/", "../", "d31/"},
			want: 3,
		},
		{
			name: "test3",
			logs: []string{"d1/", "../", "../", "../"},
			want: 0,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := MinOperations(tc.logs)
			assert.Equal(t, tc.want, got)
		})
	}
}
