package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft, maxRight := maxDepth(root.Left), maxDepth(root.Right)
	if root.Left != nil {
		maxLeft++
	}
	if root.Right != nil {
		maxRight++
	}

	if maxLeft >= maxRight {
		return maxLeft
	}
	return maxRight
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLeft, maxRight := maxDepth(root.Left), maxDepth(root.Right)
	if root.Left != nil {
		maxLeft++
	}
	if root.Right != nil {
		maxRight++
	}

	max := maxLeft + maxRight
	diameterOfLeft, diameterOfRight := diameterOfBinaryTree(root.Left), diameterOfBinaryTree(root.Right)
	if max < diameterOfLeft {
		max = diameterOfLeft
	}
	if max < diameterOfRight {
		max = diameterOfRight
	}
	return max
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

			if i >= len(vals)-1 || vals[i+1] == -1 {
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
	fmt.Println(diameterOfBinaryTree(MakeTreeNode([]int{1, 2, 3, 4, 5})))
	fmt.Println(diameterOfBinaryTree(MakeTreeNode([]int{4, -7, -3, -1, -1, -9, -3, 9, -7, -4, -1, 6, -1, -6, -6, -1, -1, 0, 6, 5, -1, 9, -1, -1, 11, -4, -1, -1, -1, -2})))
}
