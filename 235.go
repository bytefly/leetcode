package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if p.Val > q.Val {
		p, q = q, p
	}

	if root == nil || root.Val == p.Val || root.Val == q.Val {
		return root
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	if left != nil {
		return left
	}
	right := lowestCommonAncestor(root.Right, p, q)
	return right
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

			if vals[i+1] == -1 {
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
	root := MakeTreeNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 3, 5})
	fmt.Println(lowestCommonAncestor(root, &TreeNode{2, nil, nil}, &TreeNode{8, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{0, nil, nil}, &TreeNode{9, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{0, nil, nil}, &TreeNode{3, nil, nil}).Val)
	fmt.Println(lowestCommonAncestor(root, &TreeNode{3, nil, nil}, &TreeNode{9, nil, nil}).Val)
}
