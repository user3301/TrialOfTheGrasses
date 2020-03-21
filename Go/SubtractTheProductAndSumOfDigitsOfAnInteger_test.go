package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestSubtractProductSum(t *testing.T) {
	n := 234
	expected := 15
	actual := SubtractProductAndSum(n)
	assert.Equal(t, expected, actual)
}
