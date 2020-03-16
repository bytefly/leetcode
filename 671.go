package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode) (int, int) {
	one, two := root.Val, root.Val
	if root.Left != nil && root.Left.Val != two {
		two = root.Left.Val
	}
	if root.Right != nil && root.Right.Val != two && root.Right.Val < two {
		two = root.Right.Val
	}

	if two == root.Val {
		m := two
		if root.Left != nil {
			s, t := dfs(root.Left)
			if s != m {
				two = s
			} else if t != m {
				two = t
			}
		}
		if root.Right != nil {
			s, t := dfs(root.Right)
			if (s < two && s != m) || (two == m && s != m) {
				two = s
			} else if (t < two && t != m) || (two == m && t != m) {
				two = t
			}
		}
	}

	return one, two
}

func findSecondMinimumValue(root *TreeNode) int {
	one, two := dfs(root)
	if one == two {
		return -1
	}
	return two
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
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 2, 5, -1, -1, 5, 7})))
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 2, 2, 3, 4, 5, 6})))
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 2, 2, 3, 4})))
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 2, 2})))
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 4, 3})))
	fmt.Println(findSecondMinimumValue(MakeTreeNode([]int{2, 3, 4})))
}
