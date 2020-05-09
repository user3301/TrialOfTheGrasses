package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestCheckIfNAndItsDoubleExist1(t *testing.T) {
	input := []int{10, 2, 5, 3}
	expected := true
	actual := CheckIfExist(input)
	assert.Equal(t, expected, actual)
}

func TestCheckIfNAndItsDoubleExist2(t *testing.T) {
	input := []int{7, 1, 14, 11}
	expected := true
	actual := CheckIfExist(input)
	assert.Equal(t, expected, actual)
}

func TestCheckIfNAndItsDoubleExist3(t *testing.T) {
	input := []int{3, 1, 7, 11}
	expected := false
	actual := CheckIfExist(input)
	assert.Equal(t, expected, actual)
}
