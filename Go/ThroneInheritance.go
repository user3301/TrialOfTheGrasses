package gosoln

type ThroneInheritance struct {
	Name       string
	IsDead     bool
	Successors []*ThroneInheritance
}

func NewThroneInheritance(kingName string) ThroneInheritance {
	if kingName == "" {
		panic("name cannot be empty")
	}
	return ThroneInheritance{
		Name:       kingName,
		Successors: []*ThroneInheritance{},
	}
}

func (t *ThroneInheritance) Birth(parentName, childName string) {
	newThroneInheritance := NewThroneInheritance(childName)
	inheritance := locate(t, parentName)
	inheritance.Successors = append(inheritance.Successors, &newThroneInheritance)
}

func (t *ThroneInheritance) Death(name string) {
	inheritance := locate(t, name)
	inheritance.IsDead = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string

	stack := make([]*ThroneInheritance, 1)
	copy(stack, []*ThroneInheritance{t})
	for len(stack) > 0 {
		cur := stack[0]
		if !cur.IsDead {
			ans = append(ans, cur.Name)
		}
		stack = append(stack, cur.Successors...)
		stack = stack[1:]
	}
	return ans
}

func locate(root *ThroneInheritance, targetName string) *ThroneInheritance {
	if root.Name == targetName {
		return root
	}
	q := make([]*ThroneInheritance, len(root.Successors))
	copy(q, root.Successors)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			cur := q[i]
			if cur.Name == targetName {
				return q[i]
			}
			q = append(q, q[i].Successors...)
		}
		q = q[size:]
	}
	panic("cannot find")
}
