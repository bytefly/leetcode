package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var pos int
	for pos < len(inorder) && inorder[pos] != preorder[0] {
		pos++
	}

	root := &TreeNode{preorder[0], nil, nil}
	if len(preorder) > 1 {
		root.Left = buildTree(preorder[1:pos+1], inorder[:pos])
	}
	if pos < len(inorder)-1 {
		root.Right = buildTree(preorder[pos+1:], inorder[pos+1:])
	}

	return root
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

func main() {
	fmt.Println(GetTreeNodeVal(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})))
	fmt.Println(GetTreeNodeVal(buildTree([]int{1, 2, 4, 5, 3, 6, 7}, []int{4, 2, 5, 1, 6, 3, 7})))
	fmt.Println(GetTreeNodeVal(buildTree([]int{1, 2, 4, 3, 7}, []int{4, 2, 1, 3, 7})))
	fmt.Println(GetTreeNodeVal(buildTree([]int{1}, []int{1})))
}
