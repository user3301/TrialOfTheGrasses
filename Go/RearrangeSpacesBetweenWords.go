package gosoln

import "strings"

func ReorderSpaces(text string) string {
	spaceCnt := strings.Count(text, " ")
	words := strings.Fields(text)
	spaceAfter, spaceBetween := 0, 0
	if len(words) <= 1 {
		spaceAfter = spaceCnt
		spaceBetween = 0
	} else {
		spaceBetween = spaceCnt / (len(words) - 1)
		spaceAfter = spaceCnt % (len(words) - 1)
	}
	var sb strings.Builder
	for _, w := range words {
		sb.WriteString(w)
		sb.WriteString(strings.Repeat(" ", spaceBetween))
	}
	sb.WriteString(strings.Repeat(" ", spaceAfter))
	return strings.TrimSuffix(sb.String(), strings.Repeat(" ", spaceBetween))
}
