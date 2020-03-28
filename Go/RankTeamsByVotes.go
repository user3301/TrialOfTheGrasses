package Go

import (
	"bytes"
	"math"
)

func RankTeams(votes []string) string {
	if len(votes) == 1 {
		return votes[0]
	}
	stats := map[byte]int{}
	teams := votes[0]
	ranks := len(votes[0])

	for i := 0; i < ranks; i++ {
		stats[teams[i]] = 0
	}

	for i := 0; i < len(votes); i++ {
		for j := 0; j < len(votes[i]); j++ {
			aa := ranks - j
			stats[votes[i][j]] += int(math.Pow(10, float64(float32(aa))))
		}
	}

	var buf bytes.Buffer
	cur := byte('A')
	max := 0
	for len(stats) != 0 {
		for k, v := range stats {
			if v > max {
				cur = k
				max = v
			} else if v == max {
				if cur < k {
					cur = k
					max = v
				}
			}
		}
		_ = buf.WriteByte(cur)
		delete(stats, cur)
		cur = byte('A')
		max = 0
	}
	return buf.String()
}
