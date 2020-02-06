package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	sum := 0

	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			p := stack[i]
			if p.Left != nil {
				if p.Left.Left == nil && p.Left.Right == nil {
					sum += p.Left.Val
				} else {
					stack = append(stack, p.Left)
				}
			}
			if p.Right != nil {
				if p.Right.Left != nil || p.Right.Right != nil {
					stack = append(stack, p.Right)
				}
			}
		}
		stack = stack[size:]
	}

	return sum
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

			if vals[i+1] == -1 {
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
	root := MakeTreeNode([]int{1, 2, 3, -1, 4, -1, 5})
	fmt.Println(sumOfLeftLeaves(root))
	root = MakeTreeNode([]int{1, 2, 3, -1, -1, -1, 4})
	fmt.Println(sumOfLeftLeaves(root))
	root = MakeTreeNode([]int{1})
	fmt.Println(sumOfLeftLeaves(root))
	root = MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})
	fmt.Println(sumOfLeftLeaves(root))
	root = MakeTreeNode([]int{1, -1, 2, -1, 3})
	fmt.Println(sumOfLeftLeaves(root))
}
