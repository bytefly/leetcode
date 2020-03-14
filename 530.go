package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) []int {
	var nums []int
	if root.Left != nil {
		nums = dfs(root.Left)
	}
	nums = append(nums, root.Val)
	if root.Right != nil {
		t := dfs(root.Right)
		nums = append(nums, t...)
	}

	return nums
}

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxUint32
	nums := dfs(root)

	for i := len(nums) - 1; i > 0; i-- {
		t := nums[i] - nums[i-1]
		if t < ans {
			ans = t
		}
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
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{2, 1, 3})))
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{4, 2, 6, 1, 3})))
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{1, 0, 48, -1, -1, 12, 49})))
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{27, -1, 34, -1, 58, 50, -1, 44})))
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{90, 69, -1, 49, 89, -1, 52})))
	fmt.Println(getMinimumDifference(MakeTreeNode([]int{956, 267, 961, 7, 296, -1, 981, -1, 138, -1, 496, -1, -1, -1, 263, 362, 638, -1, -1, 308, 409, -1, 703, -1, 322, -1, -1, -1, 769, -1, -1, -1, 883, 808, 945, -1, -1, 915})))
}
