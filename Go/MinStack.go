package gosoln

type MinStack struct {
	stack      []int
	currentMin int
}

func NewMinStack() MinStack {
	return MinStack{
		stack: make([]int, 0),
	}
}

func (m *MinStack) Push(x int) {
	if len(m.stack) == 0 {
		m.stack = append(m.stack, x)
		m.currentMin = x
		m.stack = append(m.stack, x)
	} else {
		m.stack = append(m.stack, x)
		if x < m.currentMin {
			m.currentMin = x
		}
		m.stack = append(m.stack, m.currentMin)
	}
}

func (m *MinStack) Pop() {
	m.stack = m.stack[:len(m.stack)-2]
	if len(m.stack) > 0 {
		m.currentMin = m.stack[len(m.stack)-1]
	}
}

func (m *MinStack) Top() int {
	if len(m.stack) < 2 {
		panic("there are no elements in stack")
	}
	return m.stack[len(m.stack)-2]
}

func (m *MinStack) Min() int {
	if len(m.stack) < 2 {
		panic("no elements")
	}
	return m.stack[len(m.stack)-1]
}
