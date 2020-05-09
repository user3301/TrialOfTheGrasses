package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestReverseVowelsOfAString(t *testing.T) {
	input := "hello"
	expected := "holle"
	actual := ReverseVowels(input)
	assert.Equal(t, expected, actual)
}
