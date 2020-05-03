package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMajorityElement1(t *testing.T) {
	input := []int{3, 2, 3}
	expected := 3
	actual := MajorityElement(input)
	assert.Equal(t, expected, actual)
}

func TestMajorityElement2(t *testing.T) {
	input := []int{2, 2, 1, 1, 1, 2, 2}
	expected := 2
	actual := MajorityElement(input)
	assert.Equal(t, expected, actual)
}
