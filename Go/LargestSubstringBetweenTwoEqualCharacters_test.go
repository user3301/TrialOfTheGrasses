package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMaxLengthBetweenEqualCharacters1(t *testing.T) {
	s := "mgntdygtxrvxjnwksqhxuxtrv"
	want := 18
	got := MaxLengthBetweenEqualCharacters(s)
	assert.Equal(t, got, want)
}

func TestMaxLengthBetweenEqualCharacters2(t *testing.T) {
	s := "aa"
	want := 0
	got := MaxLengthBetweenEqualCharacters(s)
	assert.Equal(t, got, want)
}

func TestMaxLengthBetweenEqualCharacters3(t *testing.T) {
	s := "abca"
	want := 2
	got := MaxLengthBetweenEqualCharacters(s)
	assert.Equal(t, got, want)
}

func TestMaxLengthBetweenEqualCharacters4(t *testing.T) {
	s := "cbzxy"
	want := -1
	got := MaxLengthBetweenEqualCharacters(s)
	assert.Equal(t, got, want)
}

func TestMaxLengthBetweenEqualCharacters5(t *testing.T) {
	s := "cabbac"
	want := 4
	got := MaxLengthBetweenEqualCharacters(s)
	assert.Equal(t, got, want)
}
