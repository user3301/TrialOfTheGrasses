package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestPlusOne1(t *testing.T) {
	input := []int{1, 2, 3}
	expected := []int{1, 2, 4}
	actual := PlusOne(input)
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}

func TestPlusOne2(t *testing.T) {
	input := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}
	actual := PlusOne(input)
	for i := 0; i < len(expected); i++ {
		assert.Equal(t, expected[i], actual[i])
	}
}
