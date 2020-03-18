package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(t *TreeNode) string {
	if t == nil {
		return ""
	}
	ans := fmt.Sprintf("%d", t.Val)

	left := tree2str(t.Left)
	right := tree2str(t.Right)

	switch {
	case t.Left != nil && t.Right == nil:
		ans += "(" + left + ")"
	case t.Left == nil && t.Right != nil:
		ans += "()" + "(" + right + ")"
	case t.Left != nil && t.Right != nil:
		ans += "(" + left + ")" + "(" + right + ")"
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
	fmt.Println(tree2str(MakeTreeNode([]int{1, 2, 3, 4})))
	fmt.Println(tree2str(MakeTreeNode([]int{1, 2, 3, -1, 4})))
	fmt.Println(tree2str(MakeTreeNode([]int{1, 2, 3, 4, 5})))
	fmt.Println(tree2str(MakeTreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4})))
	fmt.Println(tree2str(MakeTreeNode([]int{3, 0, 4, -1, 2, -1, -1, 1})))
}
