package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestCanMakeArithmeticProgression1(t *testing.T) {
	arr := []int{3, 5, 1}
	expected := true
	actual := CanMakeArithmeticProgression(arr)
	assert.Equal(t, expected, actual)
}

func TestCanMakeArithmeticProgression2(t *testing.T) {
	arr := []int{1, 2, 4}
	expected := false
	actual := CanMakeArithmeticProgression(arr)
	assert.Equal(t, expected, actual)
}
