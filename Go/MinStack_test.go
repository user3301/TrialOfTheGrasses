package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinStack1(t *testing.T) {
	minStack := NewMinStack()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	assert.Equal(t, -3, minStack.Min())
	minStack.Pop()
	assert.Equal(t, 0, minStack.Top())
	assert.Equal(t, -2, minStack.Min())
}

//["MinStack","push","push","min","min","push","min","min","top","min","pop","push","push","min","push","pop","top","min","pop"]
//[[],[-10],[14],[],[],[-20],[],[],[],[],[],[10],[-7],[],[-7],[],[],[],[]]
//[null,null,null,-10,-10,null,-20,-20,-20,-20,null,null,null,-10,null,null,-7,-10,null]
func TestMinStack2(t *testing.T) {
	ms := NewMinStack()
	ms.Push(-10)
	ms.Push(14)
	assert.Equal(t, -10, ms.Min())
	assert.Equal(t, -10, ms.Min())
	ms.Push(-20)
	assert.Equal(t, -20, ms.Min())
	assert.Equal(t, -20, ms.Min())
	assert.Equal(t, -20, ms.Top())
	assert.Equal(t, -20, ms.Min())
	ms.Pop()
	ms.Push(10)
	ms.Push(-7)
	assert.Equal(t, -10, ms.Min())
	ms.Push(-7)
	ms.Pop()
	assert.Equal(t, -7, ms.Top())
	assert.Equal(t, -10, ms.Min())
	ms.Pop()
}
