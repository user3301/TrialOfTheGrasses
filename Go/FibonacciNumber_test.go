package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestFib1(t *testing.T) {
	input := 2
	want := 1
	got := Fib(input)
	assert.Equal(t, want, got)
}

func TestFib2(t *testing.T) {
	input := 5
	want := 5
	got := Fib(input)
	assert.Equal(t, want, got)
}
