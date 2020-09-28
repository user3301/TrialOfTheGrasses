package gosoln

import (
	"fmt"
	"testing"
)

func TestThroneInheritance1(t *testing.T) {
	t.Parallel()
	operations := []string{"ThroneInheritance", "birth", "birth", "birth", "birth", "birth", "birth", "getInheritanceOrder", "death", "getInheritanceOrder"}
	data := [][]string{{"king"}, {"king", "andy"}, {"king", "bob"}, {"king", "catherine"}, {"andy", "matthew"}, {"bob", "alex"}, {"bob", "asha"}, {""}, {"bob"}, {"null"}}
	var root ThroneInheritance
	for i := 0; i < len(operations); i++ {
		switch operations[i] {
		case "ThroneInheritance":
			root = NewThroneInheritance(data[i][0])
		case "birth":
			root.Birth(data[i][0], data[i][1])
		case "death":
			root.Death(data[i][0])
		case "getInheritanceOrder":
			order := root.GetInheritanceOrder()
			fmt.Printf("%+v", order)
		}
	}
}
