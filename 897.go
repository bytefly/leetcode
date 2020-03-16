package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var stack []*TreeNode
	var prev, newRoot *TreeNode

	p := root
	for p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if newRoot == nil {
				newRoot = p
				prev = newRoot
			} else {
				prev.Left = nil
				prev.Right = p
				prev = p
			}

			p = p.Right
			if p != nil {
				break
			}
		}
	}

	return newRoot
}

func GetTreeNodeVal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	vals := make([]int, 0)
	for len(stack) > 0 {
		size := len(stack)
		nilNum := 0
		for i := 0; i < size; i++ {
			p := stack[i]
			if p != nil {
				vals = append(vals, p.Val)
				stack = append(stack, p.Left)
				stack = append(stack, p.Right)
			} else {
				vals = append(vals, -1)
				nilNum++
			}
		}
		//trim all nil leaf nodes
		if nilNum == size {
			vals = vals[:len(vals)-size]
		}
		stack = stack[size:]
	}
	return vals
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

		node0Num := len(queue)
		for j := 0; j < node0Num; j++ {
			node0 = queue[j]

			if i >= len(vals) {
				return root
			}
			if node0 == nil {
				//no children parse
				continue
			}

			if vals[i] == -1 {
				node1 = nil
			} else {
				node1 = &TreeNode{vals[i], nil, nil}
			}
			node0.Left = node1
			queue = append(queue, node1)

			if i == len(vals)-1 || vals[i+1] == -1 {
				node2 = nil
			} else {
				node2 = &TreeNode{vals[i+1], nil, nil}
			}
			node0.Right = node2
			queue = append(queue, node2)

			i += 2
		}
		queue = queue[node0Num:]
	}

	return root
}

func main() {
	p := increasingBST(MakeTreeNode([]int{5, 3, 6, 2, 4, -1, 8, 1, -1, -1, -1, 7, 9}))
	fmt.Println(GetTreeNodeVal(p))
	p = increasingBST(MakeTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}))
	fmt.Println(GetTreeNodeVal(p))
}
