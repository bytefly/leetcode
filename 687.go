package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	var ans int
	if root == nil {
		return 0
	}

	left := longestUnivaluePath(root.Left)
	right := longestUnivaluePath(root.Right)
	ans1, ans2 := dfs(root, root.Val)
	//conver node number to the edge number
	ans1--
	ans2--

	//path cross left and right nodes
	if root.Left != nil && root.Left.Val == root.Val && root.Right != nil && root.Right.Val == root.Val {
		ans = ans1 + ans2
	} else {
		//compare left and right path length with the same value
		if ans < ans1 {
			ans = ans1
		}
		if ans < ans2 {
			ans = ans2
		}
	}

	//compare left and right subtree length with differ value
	if ans < left {
		ans = left
	}
	if ans < right {
		ans = right
	}

	return ans
}

func dfs(root *TreeNode, val int) (int, int) {
	if root == nil || root.Val != val {
		return 0, 0
	}

	left, right1 := dfs(root.Left, val)
	left1, right := dfs(root.Right, val)
	if left < right1 {
		left = right1
	}
	if left1 > right {
		right = left1
	}
	return 1 + left, 1 + right
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
	fmt.Println(longestUnivaluePath(MakeTreeNode([]int{5, 4, 5, 1, 1, -1, 5})))
	fmt.Println(longestUnivaluePath(MakeTreeNode([]int{1, 4, 5, 4, 4, -1, 5})))
}
