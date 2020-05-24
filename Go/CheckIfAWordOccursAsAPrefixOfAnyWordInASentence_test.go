package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestIsPrefixOfWord1(t *testing.T) {
	sentence := "i love eating burger"
	searchWord := "burg"
	expected := 4
	actual := IsPrefixOfWord(sentence, searchWord)
	assert.Equal(t, expected, actual)
}

func TestIsPrefixOfWord2(t *testing.T) {
	sentence := "this problem is an easy problem"
	searchWord := "pro"
	expected := 2
	actual := IsPrefixOfWord(sentence, searchWord)
	assert.Equal(t, expected, actual)
}

func TestIsPrefixOfWord3(t *testing.T) {
	sentence := "i am tired"
	searchWord := "you"
	expected := -1
	actual := IsPrefixOfWord(sentence, searchWord)
	assert.Equal(t, expected, actual)
}

func TestIsPrefixOfWord4(t *testing.T) {
	sentence := "hello from the other side"
	searchWord := "they"
	expected := -1
	actual := IsPrefixOfWord(sentence, searchWord)
	assert.Equal(t, expected, actual)
}
