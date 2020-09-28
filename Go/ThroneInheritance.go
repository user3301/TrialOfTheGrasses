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
	if t.Name == parentName {
		t.Successors = append(t.Successors, &newThroneInheritance)
	}
	inheritance := locate(t.Successors, parentName, 0)
	inheritance.Successors = append(inheritance.Successors, &newThroneInheritance)
}

func (t *ThroneInheritance) Death(name string) {
	if t.Name == name {
		t.IsDead = true
	}
	inheritance := locate(t.Successors, name, 0)
	inheritance.IsDead = true
}

func (t *ThroneInheritance) GetInheritanceOrder() []string {
	var ans []string
	if !t.IsDead {
		ans = append(ans, t.Name)
	}
	var q []*ThroneInheritance
	q = append(q, t.Successors...)
	for len(q) > 0 {
		size := len(q)
		for i := 0; i < size; i++ {
			if !q[i].IsDead {
				ans = append(ans, q[i].Name)
				q = append(q, q[i].Successors...)
			}
		}
	}
	return ans
}

func locate(tarr []*ThroneInheritance, parentName string, i int) *ThroneInheritance {
	if len(tarr) <= 0 {
		panic("not successors")
	}
	if tarr[i].Name == parentName {
		return tarr[i]
	}
	i++
	return locate(tarr, parentName, i)
}
