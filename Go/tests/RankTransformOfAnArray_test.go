package tests

import (
	"github.com/leetcode/questions"
	"gotest.tools/assert"
	"reflect"
	"testing"
)

func TestArrayRankTransform1(t *testing.T) {
	arr := []int{40,10,20,30}
	expected := []int{4,1,2,3}
	actual := questions.ArrayRankTransform(arr)
	assert.Equal(t, true, reflect.DeepEqual(expected,actual))
}
