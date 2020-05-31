package gosoln

import (
	"testing"

	"gotest.tools/assert"
)

func TestPseudopalindromicPaths1(t *testing.T) {
	root := TreeNode{Left: &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: nil,
		},
	},
		Right: &TreeNode{
			Val:   1,
			Left:  nil,
			Right: &TreeNode{Left: nil, Right: nil, Val: 1},
		},
		Val: 2}
	expected := 2
	actual := PseudopalindromicPaths(&root)
	assert.Equal(t, expected, actual)
}

//func TestIsPalindrome1(t *testing.T) {
//	s := []int{2, 3, 3}
//	expected := true
//	actual := isPalindromeList(s)
//	assert.Equal(t, expected, actual)
//}
//
//func TestIsPalindrome2(t *testing.T) {
//	s := []int{2, 1, 1}
//	expected := true
//	actual := isPalindromeList(s)
//	assert.Equal(t, expected, actual)
//}
//
//func TestIsPalindrome3(t *testing.T) {
//	s := []int{2, 3, 1}
//	expected := false
//	actual := isPalindromeList(s)
//	assert.Equal(t, expected, actual)
//}
//
//func TestIsPalindrome4(t *testing.T) {
//	s := []int{9, 5, 5, 5, 1, 1, 9, 5}
//	expected := true
//	actual := isPalindromeList(s)
//	assert.Equal(t, expected, actual)
//}
