package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
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
	root := MakeTreeNode([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})
	fmt.Println(lowestCommonAncestor(root, &TreeNode{5, nil, nil}, &TreeNode{1, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{5, nil, nil}, &TreeNode{4, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{2, nil, nil}, &TreeNode{0, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{4, nil, nil}, &TreeNode{1, nil, nil}).Val)
}
