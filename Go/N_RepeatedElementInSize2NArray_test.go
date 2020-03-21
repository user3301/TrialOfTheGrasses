package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestN_RepeatedElementInSize2NArray(t *testing.T) {
	input := []int{1, 2, 3, 3}
	expected := 3
	actual := RepeatedNTimes(input)
	assert.Equal(t, expected, actual)
}
