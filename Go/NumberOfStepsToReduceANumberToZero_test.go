package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestNumberOfSteps1(t *testing.T) {
	input := 123
	expected := 12
	actual := NumberOfSteps(input)
	assert.Equal(t, expected, actual)
}

func TestNumberOfSteps2(t *testing.T) {
	input := 14
	expected := 6
	actual := NumberOfSteps(input)
	assert.Equal(t, expected, actual)
}
