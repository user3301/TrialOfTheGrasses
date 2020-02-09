package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"testing"
)

func TestCheckIfNAndItsDoubleExist1(t *testing.T) {
	input := []int{10, 2, 5, 3}
	expected := true
	actual := questions.CheckIfExist(input)
	assert.Equal(t, expected, actual)
}

func TestCheckIfNAndItsDoubleExist2(t *testing.T) {
	input := []int{7, 1, 14, 11}
	expected := true
	actual := questions.CheckIfExist(input)
	assert.Equal(t, expected, actual)
}

func TestCheckIfNAndItsDoubleExist3(t *testing.T) {
	input := []int{3, 1, 7, 11}
	expected := false
	actual := questions.CheckIfExist(input)
	assert.Equal(t, expected, actual)
}
