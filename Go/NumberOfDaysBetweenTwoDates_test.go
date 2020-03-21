package Go

import (
	"gotest.tools/assert"
	"testing"
)

func TestNumberOfDates1(t *testing.T) {
	from := "1971-06-29"
	to := "2010-09-23"
	expected := 14331
	actual := DaysBetweenDates(from, to)
	assert.Equal(t, expected, actual)
}
