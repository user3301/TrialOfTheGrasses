package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidPalindromeII1(t *testing.T) {
	input := "abca"
	actual := ValidPalindromeII(input)
	assert.Equal(t, true, actual)
}

func TestValidPalindromeII2(t *testing.T) {
	input := "abc"
	actual := ValidPalindromeII(input)
	assert.Equal(t, false, actual)
}

func TestValidPalindromeII3(t *testing.T) {
	input := "atbbga"
	actual := ValidPalindromeII(input)
	assert.Equal(t, false, actual)
}
