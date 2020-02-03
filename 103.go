package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0)
	prevStack := make([]*TreeNode, 0)
	stack := make([]*TreeNode, 0)
	p := root

    if p == nil {
        return [][]int{}
    }

	prevStack = append(prevStack, p)
	for {
		val := make([]int, 0)
		if len(prevStack) > 0 {
			for i := len(prevStack)-1; i >= 0; i-- {
				p = prevStack[i]
				val = append(val, p.Val)
				if p.Left != nil {
					stack = append(stack, p.Left)
				}
				if p.Right != nil {
					stack = append(stack, p.Right)
				}
			}
			prevStack = prevStack[:0]
			ans = append(ans, val)
			if len(stack) == 0 {
				break
			}
		} else if len(stack) > 0 {
			for i := len(stack)-1; i >= 0; i-- {
				p = stack[i]
				val = append(val, p.Val)
				if p.Right != nil {
					prevStack = append(prevStack, p.Right)
				}
				if p.Left != nil {
					prevStack = append(prevStack, p.Left)
				}
			}
			stack = stack[:0]
			ans = append(ans, val)
			if len(prevStack) == 0 {
				break
			}
		}
	}

	return ans
}

func MakeTreeNode(vals []int) *TreeNode {
	var root, node0, node1, node2 *TreeNode
	stack := make([]*TreeNode, 0)

	for i := 0; i < len(vals); {
		if len(stack) == 0 {
			node0 = &TreeNode{vals[i], nil, nil}
			if i == 0 {
				root = node0
			}
			stack = append(stack, node0)
			i++
		} else {
			node1, node2 = nil, nil
			node0 = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if vals[i] != -1 {
				node1 = &TreeNode{vals[i], nil, nil}
				node0.Left = node1
			}
			if vals[i+1] != -1 {
				node2 = &TreeNode{vals[i+1], nil, nil}
				node0.Right = node2
			}

			if node2 != nil {
				stack = append(stack, node2)
			}
			if node1 != nil {
				stack = append(stack, node1)
			}

			i += 2
		}
	}

	return root
}

func main() {
	root := MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	ans := zigzagLevelOrder(root)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
