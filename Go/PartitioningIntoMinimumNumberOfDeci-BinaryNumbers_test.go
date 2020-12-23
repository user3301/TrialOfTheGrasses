package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinPartitions(t *testing.T) {
	expected := 3
	got := MinPartitions("32")
	assert.Equal(t, expected, got)
}
