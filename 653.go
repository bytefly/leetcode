package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var stack []*TreeNode
	var vals []int

	p := root
	for p != nil {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			vals = append(vals, p.Val)

			p = p.Right
			if p != nil {
				break
			}
		}
	}

	l, r := 0, len(vals)-1
	for l < r {
		t := vals[l] + vals[r]
		switch {
		case t == k:
			return true
		case t > k:
			r--
		case t < k:
			l++
		}
	}
	return false
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
	fmt.Println(findTarget(MakeTreeNode([]int{5, 3, 6, 2, 4, -1, 7}), 9))
	fmt.Println(findTarget(MakeTreeNode([]int{5, 3, 6, 2, 4, -1, 7}), 28))
	fmt.Println(findTarget(MakeTreeNode([]int{1}), 1))
	fmt.Println(findTarget(MakeTreeNode([]int{1}), 2))
	fmt.Println(findTarget(MakeTreeNode([]int{1, 0, 4, -2, -1, 3}), 7))
}
