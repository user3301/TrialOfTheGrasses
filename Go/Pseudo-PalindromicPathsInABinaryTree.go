package gosoln

import . "github.com/leetcode/types"

type TreeItem struct {
	Vals []int
	Node *TreeNode
}

func PseudopalindromicPaths(root *TreeNode) int {
	ans := 0
	stack := make([]*TreeItem, 0)
	var vals []int
	vals = append(vals, root.Val)
	item := &TreeItem{
		Vals: vals,
		Node: root,
	}
	stack = append(stack, item)

	for len(stack) != 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			n := len(stack) - 1
			current := stack[n]
			stack = stack[:n]
			// is leaf
			if current.Node.Left == nil && current.Node.Right == nil {
				if isPalindromeList(current.Vals) {
					ans++
				}
			}
			if current.Node.Left != nil {
				vals := append(current.Vals, current.Node.Left.Val)
				item := &TreeItem{
					Vals: vals,
					Node: current.Node.Left,
				}
				stack = append(stack, item)
			}
			if current.Node.Right != nil {
				vals := append(current.Vals, current.Node.Right.Val)
				item := &TreeItem{
					Vals: vals,
					Node: current.Node.Right,
				}
				stack = append(stack, item)
			}
		}
	}
	return ans
}

func isPalindromeList(vals []int) bool {
	dist := make(map[int]int)
	seenOdd := false
	for _, v := range vals {
		dist[v]++
	}
	for _, d := range dist {
		if d%2 != 0 {
			if seenOdd {
				return false
			} else {
				seenOdd = true
			}
		}
	}
	return true
}
