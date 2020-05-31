package gosoln

import types "github.com/leetcode/types"

type TreeItem struct {
	Bitmap int
	Node   *types.TreeNode
}

func PseudopalindromicPaths(root *types.TreeNode) int {
	ans := 0
	bit := 0
	bit ^= 1 << root.Val
	rootItem := &TreeItem{
		Bitmap: bit,
		Node:   root,
	}
	stack := make([]*TreeItem, 0)
	stack = append(stack, rootItem)
	for len(stack) != 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			n := len(stack) - 1
			current := stack[n]
			stack = stack[:n]
			if current.Node.Left == nil && current.Node.Right == nil {
				if current.Bitmap == 0 || (current.Bitmap&(current.Bitmap-1) == 0) {
					ans++
				}
			}
			if current.Node.Left != nil {
				bits := current.Bitmap
				bits ^= 1 << current.Node.Left.Val
				item := &TreeItem{
					Bitmap: bits,
					Node:   current.Node.Left,
				}
				stack = append(stack, item)
			}
			if current.Node.Right != nil {
				bits := current.Bitmap
				bits ^= 1 << current.Node.Right.Val
				item := &TreeItem{
					Bitmap: bits,
					Node:   current.Node.Right,
				}
				stack = append(stack, item)
			}
		}
	}
	return ans
}
