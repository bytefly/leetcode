package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBSTEx(root *TreeNode) (int, int, int) {
	minDiff, min, max := root.Val, root.Val, root.Val

	if root.Left == nil && root.Right == nil {
		return min, -1, max
	}

	if root.Left != nil {
		r, s, t := minDiffInBSTEx(root.Left)
		if s < minDiff && s >= 0 {
			minDiff = s
		}

		s = root.Val - t
		if s < minDiff {
			minDiff = s
		}

		min = r
	}

	if root.Right != nil {
		r, s, t := minDiffInBSTEx(root.Right)
		if s < minDiff && s >= 0 {
			minDiff = s
		}

		s = r - root.Val
		if s < minDiff {
			minDiff = s
		}

		max = t
	}

	return min, minDiff, max
}

func minDiffInBST(root *TreeNode) int {
	_, minDiff, _ := minDiffInBSTEx(root)
	return minDiff
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
	fmt.Println(minDiffInBST(MakeTreeNode([]int{4, 2, 6, 1, 3})))
	fmt.Println(minDiffInBST(MakeTreeNode([]int{1, 0, 48, -1, -1, 12, 49})))
	fmt.Println(minDiffInBST(MakeTreeNode([]int{27, -1, 34, -1, 58, 50, -1, 44})))
	fmt.Println(minDiffInBST(MakeTreeNode([]int{90, 69, -1, 49, 89, -1, 52})))
	fmt.Println(minDiffInBST(MakeTreeNode([]int{956, 267, 961, 7, 296, -1, 981, -1, 138, -1, 496, -1, -1, -1, 263, 362, 638, -1, -1, 308, 409, -1, 703, -1, 322, -1, -1, -1, 769, -1, -1, -1, 883, 808, 945, -1, -1, 915})))
}
