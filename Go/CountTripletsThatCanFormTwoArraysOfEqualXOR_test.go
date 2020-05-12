package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestCountTriplets1(t *testing.T) {
	arr := []int{2, 3, 1, 6, 7}
	expected := 4
	actual := CountTriplets(arr)
	assert.Equal(t, expected, actual)
}

func TestCountTriplets2(t *testing.T) {
	arr := []int{1, 1, 1, 1, 1}
	expected := 10
	actual := CountTriplets(arr)
	assert.Equal(t, expected, actual)
}

func TestCountTriplets3(t *testing.T) {
	arr := []int{2, 3}
	expected := 0
	actual := CountTriplets(arr)
	assert.Equal(t, expected, actual)
}

func TestCountTriplets4(t *testing.T) {
	arr := []int{7, 11, 12, 9, 5, 2, 7, 17, 22}
	expected := 8
	actual := CountTriplets(arr)
	assert.Equal(t, expected, actual)
}
