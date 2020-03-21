package Go

import (
	"bytes"
	"gotest.tools/assert"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := []byte{'h', 'e', 'l', 'l', 'o'}
	expected := []byte{'o', 'l', 'l', 'e', 'h'}
	ReverseString(s)
	assert.Assert(t, bytes.Compare(s, expected) == 0)
}
