package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"testing"
)

func TestClosestDivisors1(t *testing.T) {
	input := 123
	expected := []int {5,25}
	actual := questions.ClosestDivisors(input)
	assert.DeepEqual(t, expected,actual)
}
