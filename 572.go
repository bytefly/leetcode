package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtreeEx(s *TreeNode, t *TreeNode) bool {
	//both are nil
	if s == t {
		return true
	} else if s == nil || t == nil || s.Val != t.Val {
		return false
	}

	return isSubtreeEx(s.Left, t.Left) && isSubtreeEx(s.Right, t.Right)
}

func isSubtree(s *TreeNode, t *TreeNode) bool {
	//both are nil
	if s == t {
		return true
	} else if s == nil || t == nil {
		return false
	}

	//have same root/left/right structure or value
	if isSubtreeEx(s, t) {
		return true
	}

	if isSubtree(s.Left, t) {
		return true
	}
	return isSubtree(s.Right, t)
}

func GetTreeNodeVal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	vals := make([]int, 0)
	for len(stack) > 0 {
		size := len(stack)
		nilNum := 0
		for i := 0; i < size; i++ {
			p := stack[i]
			if p != nil {
				vals = append(vals, p.Val)
				stack = append(stack, p.Left)
				stack = append(stack, p.Right)
			} else {
				vals = append(vals, -1)
				nilNum++
			}
		}
		//trim all nil leaf nodes
		if nilNum == size {
			vals = vals[:len(vals)-size]
		}
		stack = stack[size:]
	}
	return vals
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
	fmt.Println(isSubtree(MakeTreeNode([]int{3, 4, 5, 1, 2}), MakeTreeNode([]int{4, 1, 2})))
	fmt.Println(isSubtree(MakeTreeNode([]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0}), MakeTreeNode([]int{4, 1, 2})))
	fmt.Println(isSubtree(MakeTreeNode([]int{3, 4, 5, 1, 2, -1, -1, -1, -1, 0}), MakeTreeNode([]int{4, 1, 2, -1, -1, 0})))
	fmt.Println(isSubtree(MakeTreeNode([]int{3, 4, 5, 1, -1, 2}), MakeTreeNode([]int{3, 1, 2})))
	fmt.Println(isSubtree(MakeTreeNode([]int{4, 1}), MakeTreeNode([]int{1})))
}
