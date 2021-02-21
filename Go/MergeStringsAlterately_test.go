package gosoln

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	word1 := "ab"
	word2 := "qprs"
	expected := "aqbprs"
	actual := MergeAlternately(word1, word2)
	assert.Equal(t, expected, actual)
}
