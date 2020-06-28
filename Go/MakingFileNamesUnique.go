package gosoln

import (
	"fmt"
)

func GetFolderNames(names []string) []string {
	dist := make(map[string]int)
	ans := make([]string, 0)
	for _, s := range names {
		if _, exist := dist[s]; exist {
			for i := dist[s] + 1; ; i++ {
				newStr := fmt.Sprintf("%s(%d)", s, i)
				if _, ex := dist[newStr]; !ex {
					dist[newStr] = 0
					dist[s] = i
					ans = append(ans, newStr)
					break
				}
			}
		} else {
			dist[s] = 0
			ans = append(ans, s)
		}
	}
	return ans
}
