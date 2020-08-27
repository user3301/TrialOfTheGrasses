package gosoln

import "fmt"

func IsPathCrossing(path string) bool {
	x, y := 0, 0
	set := make(map[string]struct{})
	set[fmt.Sprintf("%d&%d", x, y)] = struct{}{}
	for _, c := range path {
		switch c {
		case 'N':
			y++
			if _, ex := set[fmt.Sprintf("%d&%d", x, y)]; ex {
				return true
			}
			set[fmt.Sprintf("%d&%d", x, y)] = struct{}{}
		case 'S':
			y--
			if _, ex := set[fmt.Sprintf("%d&%d", x, y)]; ex {
				return true
			}
			set[fmt.Sprintf("%d&%d", x, y)] = struct{}{}
		case 'E':
			x++
			if _, ex := set[fmt.Sprintf("%d&%d", x, y)]; ex {
				return true
			}
			set[fmt.Sprintf("%d&%d", x, y)] = struct{}{}
		case 'W':
			x--
			if _, ex := set[fmt.Sprintf("%d&%d", x, y)]; ex {
				return true
			}
			set[fmt.Sprintf("%d&%d", x, y)] = struct{}{}
		}
	}
	return false
}
