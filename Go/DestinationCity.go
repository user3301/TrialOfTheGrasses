package Go

func destCity(paths [][]string) string {
	dist := make(map[string]string)
	for i := 0; i < len(paths); i++ {
		dist[paths[i][0]] = paths[i][1]
	}

	ok := true
	start := paths[0][0]
	for ok {
		_, ok = dist[start]
		if ok {
		} else {
			break
		}
		start = dist[start]

	}
	return start
}
