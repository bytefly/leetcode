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
	stack := make([]*TreeNode, 0)
	p := root
	for {
		for p != nil {
			vals = append(vals, p.Val)
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		} else {
			break
		}
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
