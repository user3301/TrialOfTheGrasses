package gosoln

func PeopleIndexes(favoriteCompanies [][]string) []int {
	ans := make([]int, 0)

	sets := make(map[int]map[string]bool, 0)

	for _, f := range favoriteCompanies {
		m := make(map[string]bool, 0)
		for _, v := range f {
			m[v] = true
		}
		sets[len(f)] = m
	}

	for i := 0; i < len(favoriteCompanies); i++ {
		f := 0
		for _, a := range sets {
			if len(favoriteCompanies[i]) > len(a) {
				continue
			} else {
				if subset(favoriteCompanies[i], a) {
					f++
				}
			}
		}
		if f < 2 {
			ans = append(ans, i)
		}
	}
	return ans
}

func subset(d1 []string, d2 map[string]bool) bool {
	for _, v := range d1 {
		if exist, _ := d2[v]; !exist {
			return false
		}
	}
	return true
}
