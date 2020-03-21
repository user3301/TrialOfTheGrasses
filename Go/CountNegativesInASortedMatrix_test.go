package Go

import (
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
	actual := CountNegativesBF(input)
	assert.Equal(t, expected, actual)
}

func BenchmarkCountNegativesBF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountNegativesBF(input)
	}
}

func BenchmarkCountNegativesBFOptimized(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountNegativesBFOptimized(input)
	}
}

func BenchmarkCountNegativesLinear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CountNegativesLinear(input)
	}
}
