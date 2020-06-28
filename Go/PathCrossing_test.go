package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestPathCrossing1(t *testing.T) {
	path := "NESWW"
	expected := true
	actual := IsPathCrossing(path)
	assert.Equal(t, expected, actual)
}

func TestPathCrossing2(t *testing.T) {
	path := "NES"
	expected := false
	actual := IsPathCrossing(path)
	assert.Equal(t, expected, actual)
}
