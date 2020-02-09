package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"testing"
)

func TestNumberOfSteps1(t *testing.T) {
	input := 123
	expected := 12
	actual := questions.NumberOfSteps(input)
	assert.Equal(t, expected, actual)
}

func TestNumberOfSteps2(t *testing.T) {
	input := 14
	expected := 6
	actual := questions.NumberOfSteps(input)
	assert.Equal(t, expected, actual)
}