package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, sum int) int {
	if root == nil {
		return 0
	}

	return pathSum(root.Left, sum) + pathSum(root.Right, sum) + dfs(root, sum)
}

func dfs(root *TreeNode, sum int) int {
	var ans int

	if root == nil {
		return ans
	}

	if root.Val == sum {
		ans++
	}
	return ans + dfs(root.Left, sum-root.Val) + dfs(root.Right, sum-root.Val)
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
	fmt.Println(pathSum(MakeTreeNode([]int{10, 5, -3, 3, 2, -1, 11, 3, -2, -1, 1}), 8))
	fmt.Println(pathSum(MakeTreeNode([]int{0, 1, 1}), 1))
}
