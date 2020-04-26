package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestBullsAndCows(t *testing.T) {
	secret, guess := "1807", "7810"
	expected := "1A3B"
	actual := GetHint(secret, guess)
	assert.Equal(t, expected, actual)
}
