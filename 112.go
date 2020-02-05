package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeSum struct {
	Node *TreeNode
	Sum  int
}

func hasPathSum(root *TreeNode, sum int) bool {
	var t *TreeNodeSum
	if root == nil {
		return false
	}

	rootSum := &TreeNodeSum{root, root.Val}
	stack := make([]*TreeNodeSum, 0)
	stack = append(stack, rootSum)

	for len(stack) > 0 {
		size := len(stack)
		for i := 0; i < size; i++ {
			p := stack[i]
			if p.Node.Left != nil {
				t = &TreeNodeSum{p.Node.Left, p.Sum + p.Node.Left.Val}
				stack = append(stack, t)
			}
			if p.Node.Right != nil {
				t = &TreeNodeSum{p.Node.Right, p.Sum + p.Node.Right.Val}
				stack = append(stack, t)
			}
			//find the leaf and check the sum
			if p.Node.Left == nil && p.Node.Right == nil {
				if p.Sum == sum {
					return true
				}
			}
		}
		stack = stack[size:]
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
	root := MakeTreeNode([]int{1, 2, 2, 3, 4, 4, 3})
	fmt.Println(hasPathSum(root, 7))
	root = MakeTreeNode([]int{1, 2, 2, -1, 3, -1, 3})
	fmt.Println(hasPathSum(root, 6))
	root = MakeTreeNode([]int{1, 2, 2, -1, -1, -1, 3})
	fmt.Println(hasPathSum(root, 3))
	root = MakeTreeNode([]int{1})
	fmt.Println(hasPathSum(root, 1))
	root = MakeTreeNode([]int{})
	fmt.Println(hasPathSum(root, 1))
	root = MakeTreeNode([]int{9, -42, -42, -1, 76, 76, -1, -1, 13, -1, 13})
	fmt.Println(hasPathSum(root, 3))
}
