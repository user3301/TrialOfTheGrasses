package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMyAtoi1(t *testing.T) {
	input := "42"
	want := 42
	got := MyAtoi(input)
	assert.Equal(t, want, got)
}

func TestMyAtoi2(t *testing.T) {
	input := "4193 with words"
	want := 4193
	got := MyAtoi(input)
	assert.Equal(t, want, got)
}

func TestMyAtoi3(t *testing.T) {
	input := "+1"
	want := 1
	got := MyAtoi(input)
	assert.Equal(t, want, got)
}
