package gosoln

import (
	"gotest.tools/assert"
	"testing"
)

func TestCheckStringIsNumber(t *testing.T) {
	testCases := []struct {
		testName, input string
		want bool
	}{
		{
			testName: "test1",
			input: "0",
			want: true,
		},
		{
			testName: "test2",
			input: "+100",
			want: true,
		},
		{
			testName: "test3",
			input: "5e2",
			want: true,
		},
		{
			testName: "test4",
			input: "-123",
			want: true,
		},
		{
			testName: "test5",
			input: "3.1415926",
			want: true,
		},
		{
			testName: "test6",
			input: "-1E-16",
			want: true,
		},
		{
			testName: "test7",
			input: "0123",
			want: true,
		},
		{
			testName: "test8",
			input: "12e",
			want: false,
		},
		{
			testName: "test9",
			input: "1a3.14",
			want: false,
		},
		{
			testName: "test10",
			input: "1.2.3",
			want: false,
		},
		{
			testName: "test11",
			input: "+-5",
			want: false,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func (t *testing.T)  {
			got := IsNumberString(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
