package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	ans, _, _ := dfs(root)
	return ans
}

func dfs(root *TreeNode) (bool, int, int) {
	if root == nil {
		return true, -1, -1
	}

	min, max := root.Val, root.Val
	if root.Left != nil {
		leftValid, leftMin, leftMax := dfs(root.Left)
		if !leftValid || leftMin >= root.Val || leftMax >= root.Val {
			return false, -1, -1
		}
		min = leftMin
	}
	if root.Right != nil {
		rightValid, rightMin, rightMax := dfs(root.Right)
		if !rightValid || rightMin <= root.Val || rightMax <= root.Val {
			return false, -1, -1
		}
		max = rightMax
	}

	return true, min, max
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
	fmt.Println(isValidBST(MakeTreeNode([]int{5, 1, 4, -1, -1, 3, 6})))
	fmt.Println(isValidBST(MakeTreeNode([]int{2, 1, 3})))
}
