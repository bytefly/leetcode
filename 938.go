package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, L int, R int) int {
	var sum int

	if root == nil {
		return 0
	}

	if root.Val >= L && root.Val <= R {
		sum += root.Val
	}

	if root.Val >= L {
		sum += rangeSumBST(root.Left, L, R)
	}
	if root.Val <= R {
		sum += rangeSumBST(root.Right, L, R)
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
	fmt.Println(rangeSumBST(MakeTreeNode([]int{10, 5, 15, 3, 7, -1, 18}), 7, 15))
	fmt.Println(rangeSumBST(MakeTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, -1, 6}), 6, 10))
}
