package gosoln

import "strings"

func ArrangeWords(text string) string {
	var ans strings.Builder
	words := strings.Fields(text)
	list := make([][]string, 100)
	for i := 0; i < len(list); i++ {
		list[i] = make([]string, 0)
	}

	for _, word := range words {
		list[len(word)] = append(list[len(word)], word)
	}
	firstFound := false
	for i := 0; i < len(list); i++ {
		if len(list[i]) != 0 {
			for _, w := range list[i] {
				if !firstFound {
					ans.WriteString(strings.Title(w))
					ans.WriteString(" ")
					firstFound = true
				} else {
					ans.WriteString(strings.ToLower(w))
					ans.WriteString(" ")
				}
			}
		}
	}
	return strings.TrimRight(ans.String(), "\t \n")
}
