package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	pVals := GetTreeNodeVal(p)
	qVals := GetTreeNodeVal(q)
	if len(pVals) != len(qVals) {
		return false
	}
	for i := 0; i < len(qVals); i++ {
		if pVals[i] != qVals[i] {
			return false
		}
	}
	return true
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
	p := MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	q := MakeTreeNode([]int{3, 9, 20, 11, 12, 15, 7})
	fmt.Println(isSameTree(p, q))
}
