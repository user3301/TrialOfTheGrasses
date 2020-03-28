package Go

import (
	"testing"

	"gotest.tools/assert"
)

func TestBackspaceStringCompare1(t *testing.T) {
	S, T := "ab#c", "ad#c"
	actual := BackspaceCompare(S, T)
	assert.Equal(t, true, actual)
}

func TestBackspaceStringCompare2(t *testing.T) {
	S, T := "ab##", "c#d#"
	actual := BackspaceCompare(S, T)
	assert.Equal(t, true, actual)
}
