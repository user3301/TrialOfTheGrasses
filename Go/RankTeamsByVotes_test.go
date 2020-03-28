package Go

import (
	"testing"
)

func TestRankTeams1(t *testing.T) {
	votes := []string{
		"ABC",
		"ACB",
		"ABC",
		"ACB",
		"ACB",
	}
	_ = RankTeams(votes)
}
