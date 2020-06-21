package gosoln

import (
	"fmt"
	"strconv"
)

func GetFolderNames(names []string) []string {
	dist := make(map[string]int)
	ans := make([]string, len(names))
	for i, v := range names {
		s := exists(dist, v)
		ans[i] = s
	}
	return ans
}

func exists(dist map[string]int, name string) string {
	if _, ex := dist[name]; !ex {
		dist[name]++
		return name
	} else if ex {
		newStr := fmt.Sprintf("%s(1)", name)
		return exists(dist, newStr)
	} else if name[len(name)-1:] == ")" {
		aa := name[len(name)-1:]
		fmt.Println(aa)
		i := name[len(name)-2 : len(name)-1]
		ii, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		newStr := fmt.Sprintf("%s(%d)", name[:len(name)-3], ii+1)
		return exists(dist, newStr)
	} else {
		panic("a")
	}
}
