package Go

import (
	"bytes"
	"testing"

	"gotest.tools/assert"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	ReverseString(s)
	assert.Assert(t, bytes.Equal(s, expected))
}
