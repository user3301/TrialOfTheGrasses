package tests

import (
	"testing"

	"github.com/leetcode/questions"
	"gotest.tools/assert"
)

func TestRemovePalindromeSub1(t *testing.T) {
	input := "ababa"
	expected := 1
	actual := questions.RemovePalindromeSub(input)
	assert.Equal(t, expected, actual)
}

func TestRemovePalindromeSub2(t *testing.T) {
	input := "abb"
	expected := 2
	actual := questions.RemovePalindromeSub(input)
	assert.Equal(t, expected, actual)
}

func TestRemovePalindromeSub3(t *testing.T) {
	input := "baabb"
	expected := 2
	actual := questions.RemovePalindromeSub(input)
	assert.Equal(t, expected, actual)
}

func TestRemovePalindromeSub4(t *testing.T) {
	input := ""
	expected := 0
	actual := questions.RemovePalindromeSub(input)
	assert.Equal(t, expected, actual)
}
