package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestMinDays1(t *testing.T) {
	bloomDay := []int{1, 10, 3, 10, 2}
	m, k := 3, 1
	expected := 3
	actual := MinDays(bloomDay, m, k)
	assert.Equal(t, expected, actual)
}

func TestMinDays2(t *testing.T) {
	bloomDay := []int{1, 10, 3, 10, 2}
	m, k := 3, 2
	expected := -1
	actual := MinDays(bloomDay, m, k)
	assert.Equal(t, expected, actual)
}

func TestMinDays3(t *testing.T) {
	bloomDay := []int{7, 7, 7, 7, 12, 7, 7}
	m, k := 2, 3
	expected := 12
	actual := MinDays(bloomDay, m, k)
	assert.Equal(t, expected, actual)
}
