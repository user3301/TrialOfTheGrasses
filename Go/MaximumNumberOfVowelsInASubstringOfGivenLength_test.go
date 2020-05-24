package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxVowels1(t *testing.T) {
	s := "abciiidef"
	k := 3
	expected := 3
	actual := MaxVowels(s, k)
	assert.Equal(t, expected, actual)
}

func TestMaxVowels2(t *testing.T) {
	s := "aeiou"
	k := 2
	expected := 2
	actual := MaxVowels(s, k)
	assert.Equal(t, expected, actual)
}

func TestMaxVowels3(t *testing.T) {
	s := "leetcode"
	k := 3
	expected := 2
	actual := MaxVowels(s, k)
	assert.Equal(t, expected, actual)
}

func TestMaxVowels4(t *testing.T) {
	s := "rhythms"
	k := 4
	expected := 0
	actual := MaxVowels(s, k)
	assert.Equal(t, expected, actual)
}

func TestMaxVowels5(t *testing.T) {
	s := "tryhard"
	k := 4
	expected := 1
	actual := MaxVowels(s, k)
	assert.Equal(t, expected, actual)
}
