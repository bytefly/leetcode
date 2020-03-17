package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inOrder(root *TreeNode) (int, int) {
	var sum, tilt int
	if root == nil {
		return 0, 0
	}

	sum += root.Val
	sumLeft, tiltLeft := inOrder(root.Left)
	sumRight, tiltRight := inOrder(root.Right)

	sum += sumLeft + sumRight
	tilt += tiltLeft + tiltRight

	t := sumLeft - sumRight
	if t < 0 {
		t = -t
	}
	tilt += t
	return sum, tilt
}

func findTilt(root *TreeNode) int {
	_, tilt := inOrder(root)
	return tilt
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
	fmt.Println(findTilt(MakeTreeNode([]int{3, 5, 1, 6, 2, 9, 8, -1, -1, 7, 4})))
	fmt.Println(findTilt(MakeTreeNode([]int{3, 0, 4, -1, 2, -1, -1, 1})))
	fmt.Println(findTilt(MakeTreeNode([]int{1, 2, 3})))
	fmt.Println(findTilt(MakeTreeNode([]int{1, 2, 3, 4, -1, 5})))
}
