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
			for i := len(prevStack) - 1; i >= 0; i-- {
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
			for i := len(stack) - 1; i >= 0; i-- {
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
	queue := make([]*TreeNode, 0)
	i := 0

	for {
		if i == 0 {
			if i >= len(vals) {
				return root
			}
			node0 = &TreeNode{vals[i], nil, nil}
			root = node0
			queue = append(queue, node0)
			i++
			continue
		}

		node0Num, childNum := len(queue), 0
		for j := 0; j < node0Num; j++ {
			node0 = queue[j]

			if i >= len(vals) {
				return root
			}
			if vals[i] != -1 {
				node1 = &TreeNode{vals[i], nil, nil}
				node0.Left = node1
				queue = append(queue, node1)
				childNum++
			}
			i++

			if i >= len(vals) {
				return root
			}
			if vals[i] != -1 {
				node2 = &TreeNode{vals[i], nil, nil}
				node0.Right = node2
				queue = append(queue, node2)
				childNum++
			}
			i++
		}
		queue = queue[node0Num:]
	}

	return root
}

func main() {
	root := MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	//root := MakeTreeNode([]int{3, 9, 20, 11, 12, 15, 7})
	ans := zigzagLevelOrder(root)
	for i := 0; i < len(ans); i++ {
		fmt.Println(ans[i])
	}
}
