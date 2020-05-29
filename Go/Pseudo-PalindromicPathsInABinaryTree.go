package gosoln

import . "github.com/leetcode/types"

func PseudopalindromicPaths(root *TreeNode) int {
	ans := 0
	paths := new([][]int)
	if root != nil {
		searchBT(root, []int{}, paths)
	}
	for _, v := range *paths {
		if isPseudoPalindrome(v) {
			ans++
		}
	}
	return ans
}

func searchBT(node *TreeNode, nums []int, paths *[][]int) {
	if node.Left == nil && node.Right == nil {
		nums = append(nums, node.Val)
		*paths = append(*paths, nums)
	}
	if node.Left != nil {
		nums = append(nums, node.Val)
		searchBT(node.Left, nums, paths)
	}
	if node.Right != nil {
		nums = append(nums, node.Val)
		searchBT(node.Right, nums, paths)
	}
}

func isPseudoPalindrome(ints []int) bool {
	dist := make(map[int]int)
	for i := 0; i < len(ints); i++ {
		dist[ints[i]]++
	}
	odd := 0
	for _, v := range dist {
		if v%2 != 0 {
			odd++
		}
	}
	if odd%2 != 0 || odd == 0 {
		return true
	} else {
		return false
	}
}
