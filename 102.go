package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	var q []*TreeNode

	q = append(q, root)
	for {
		size := len(q)
		var item []int
		for i := 0; i < size; i++ {
			p := q[i]
			if p != nil {
				item = append(item, p.Val)
				q = append(q, p.Left)
				q = append(q, p.Right)
			}
		}

		q = q[size:]
		if len(q) == 0 {
			break
		}
		ans = append(ans, item)
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
	fmt.Println(levelOrder(MakeTreeNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1})))
	fmt.Println(levelOrder(MakeTreeNode([]int{0, 1, 1})))
	fmt.Println(levelOrder(MakeTreeNode([]int{3, 9, 20, -1, -1, 15, 7})))
	fmt.Println(levelOrder(MakeTreeNode([]int{})))
}
