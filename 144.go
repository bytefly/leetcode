package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	vals := make([]int, 0)
	if root != nil {
		vals = append(vals, root.Val)
		left := preorderTraversal(root.Left)
		vals = append(vals, left...)
		right := preorderTraversal(root.Right)
		vals = append(vals, right...)
	}

	return vals
}

func main() {
	node0 := &TreeNode{1, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{3, nil, nil}
	node0.Right = node1
	node1.Left = node2
	fmt.Println(preorderTraversal(node0))
}
