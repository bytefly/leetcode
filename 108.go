package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}

	root := &TreeNode{nums[len(nums)/2], nil, nil}
	root.Left = sortedArrayToBST(nums[:len(nums)/2])
	root.Right = sortedArrayToBST(nums[len(nums)/2+1:])
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
	root := sortedArrayToBST([]int{-10, -3, 0, 5, 9})
	fmt.Println(GetTreeNodeVal(root))
	root = sortedArrayToBST([]int{1, 2, 3, 4, 5})
	fmt.Println(GetTreeNodeVal(root))
	root = sortedArrayToBST([]int{1, 2, 3})
	fmt.Println(GetTreeNodeVal(root))
	root = sortedArrayToBST([]int{1, 2})
	fmt.Println(GetTreeNodeVal(root))
}
