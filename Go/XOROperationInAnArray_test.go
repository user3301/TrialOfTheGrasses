package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestXorOperation(t *testing.T) {
	n, start := 5, 0
	expected := 8
	actual := XorOperation(n, start)
	assert.Equal(t, expected, actual)
}
