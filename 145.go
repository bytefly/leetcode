package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	vals := make([]int, 0)
	stack := make([]*TreeNode, 0)
	p := root

	stack = append(stack, p)
	for {
		if len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, p.Val)
		} else {
			break
		}

		if p.Right != nil {
			stack = append(stack, p.Right)
		}
		if p.Left != nil {
			stack = append(stack, p.Left)
		}
	}
	//reverse vals slice
	ans := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		ans[i] = vals[len(vals)-1-i]
	}
	return ans
}

func main() {
	node0 := &TreeNode{3, nil, nil}
	node1 := &TreeNode{2, nil, nil}
	node2 := &TreeNode{1, nil, nil}
	node0.Right = node1
	node1.Left = node2
	fmt.Println(postorderTraversal(node0))
}
