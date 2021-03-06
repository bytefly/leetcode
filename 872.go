package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getLeafs(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode

	p := root
	for {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) == 0 {
			break
		}

		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if p.Left == nil && p.Right == nil {
				ans = append(ans, p.Val)
			}

			p = p.Right
			if p != nil {
				break
			}
		}
	}

	return ans
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1, leafs2 := getLeafs(root1), getLeafs(root2)
	if len(leafs1) != len(leafs2) {
		return false
	}
	for i := 0; i < len(leafs1); i++ {
		if leafs1[i] != leafs2[i] {
			return false
		}
	}
	return true
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
	fmt.Println(leafSimilar(MakeTreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4}), MakeTreeNode([]int{3, 5, 1, 6, 7, 4, 2, -1, -1, -1, -1, -1, -1, 9, 8})))
	fmt.Println(leafSimilar(MakeTreeNode([]int{3, 0, 4, -1, 2, -1, -1, 1}), MakeTreeNode([]int{1, 0, 2})))
}
