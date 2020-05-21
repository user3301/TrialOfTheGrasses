package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestPeopleIndexes1(t *testing.T) {
	favoriteCompanies := [][]string{{"leetcode", "google", "facebook"},
		{"google", "microsoft"},
		{"google", "facebook"},
		{"google"},
		{"amazon"},
	}
	expected := []int{0, 1, 4}
	actual := PeopleIndexes(favoriteCompanies)
	assert.DeepEqual(t, expected, actual)
}

func TestPeopleIndexes2(t *testing.T) {
	favoriteCompanies := [][]string{
		{"leetcode", "google", "facebook"},
		{"leetcode", "amazon"},
		{"facebook", "google"},
	}
	expected := []int{0, 1}
	actual := PeopleIndexes(favoriteCompanies)
	assert.DeepEqual(t, expected, actual)
}
