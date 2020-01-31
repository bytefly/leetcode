package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGstWithSum(root *TreeNode, b int) int {
	if root == nil {
		return b
	}
	root.Val += bstToGstWithSum(root.Right, b)
	return bstToGstWithSum(root.Left, root.Val)
}

func bstToGst(root *TreeNode) *TreeNode {
	bstToGstWithSum(root, 0)
	return root
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}

	var prev *ListNode
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	if prev != nil {
		prev.Next = nil
	}

	tree := &TreeNode{slow.Val, nil, nil}
	if slow != head {
		tree.Left = sortedListToBST(head)
	}
	tree.Right = sortedListToBST(slow.Next)
	return tree
}

func createListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	var head, p *ListNode
	for i := 0; i < len(vals); i++ {
		node := &ListNode{vals[i], nil}
		if i == 0 {
			head = node
		} else {
			p.Next = node
		}
		p = node
	}
	return head
}

func getValFromTreeNode(head *TreeNode) []int {
	vals := make([]int, 0)
	if head != nil {
		vals = append(vals, head.Val)
		left := getValFromTreeNode(head.Left)
		right := getValFromTreeNode(head.Right)
		vals = append(vals, left...)
		vals = append(vals, right...)
	}

	return vals
}

func main() {
	p := sortedListToBST(createListNode([]int{0, 1, 2, 3, 4, 5, 6, 7, 8}))
	p = bstToGst(p)
	fmt.Println(getValFromTreeNode(p))
}
