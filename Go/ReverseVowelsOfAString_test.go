package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestReverseVowelsOfAString(t *testing.T) {
	input := "hello"
	expected := "holle"
	actual := ReverseVowels(input)
	assert.Equal(t, expected, actual)
}
