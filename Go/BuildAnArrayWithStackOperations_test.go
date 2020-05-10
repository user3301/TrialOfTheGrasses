package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestBuildArray1(t *testing.T) {
	target := []int{1, 3}
	n := 3
	expected := []string{"Push", "Push", "Pop", "Push"}
	actual := BuildArray(target, n)
	assert.DeepEqual(t, expected, actual)
}
