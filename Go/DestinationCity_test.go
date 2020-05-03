package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestDestCity1(t *testing.T) {
	input := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	expected := "Sao Paulo"
	actual := destCity(input)
	assert.Equal(t, expected, actual)
}
