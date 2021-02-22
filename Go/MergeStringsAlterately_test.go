package gosoln

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	testCases := []struct {
		word1, word2 string
		want         string
	}{
		{
			word1: "abc",
			word2: "pqr",
			want:  "apbqcr",
		},
		{
			word1: "ab",
			word2: "pqrs",
			want:  "apbqrs",
		},
		{
			word1: "abcd",
			word2: "pq",
			want:  "apbqcd",
		},
	}
	for i, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%dth test case", i), func(t *testing.T) {
			got := mergeAlternately(tc.word1, tc.word2)
			assert.Equal(t, tc.want, got)
		})
	}
}
