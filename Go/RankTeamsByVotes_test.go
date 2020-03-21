package Go

import (
	"testing"
)

func TestRankTeams1(T *testing.T) {
	votes := []string{
		"ABC",
		"ACB",
		"ABC",
		"ACB",
		"ACB",
	}
	_ = RankTeams(votes)
}