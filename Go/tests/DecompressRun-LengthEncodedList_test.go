package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestDecompressRLEList(t *testing.T) {
	input := []int{1, 2, 3, 4}
	expected := []int{2, 4, 4, 4}
	actual := questions.DecompressRLEList(input)
	assert.Assert(t, true, reflect.DeepEqual(expected, actual))
}
