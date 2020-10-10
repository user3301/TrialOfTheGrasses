package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum3(t *testing.T) {
	k, n := 3, 7
	want := [][]int{{1, 2, 4}}
	got := CombinationSum3(k, n)
	assert.ElementsMatch(t, want, got)
}
