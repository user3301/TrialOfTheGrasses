package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"testing"
)

var input [][]int

func init() {
	input = [][]int{
		{4, 3, 2, -1},
		{3, 2, 1, -1},
		{1, 1, -1, -2},
		{-1, -1, -2, -3},
	}
}

func TestCountNegativesBF(t *testing.T) {
	expected := 8
	actual := questions.CountNegativesBF(input)
	assert.Equal(t, expected, actual)
}

func BenchmarkCountNegativesBF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		questions.CountNegativesBF(input)
	}
}

func BenchmarkCountNegativesBFOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		questions.CountNegativesBFOptimized(input)
	}
}

func BenchmarkCountNegativesLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		questions.CountNegativesLinear(input)
	}
}
