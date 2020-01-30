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

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{head.Val, nil, nil}
	}

	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	//fmt.Println("++", getValFromListNode(head), slow.Val)
	if prev != nil {
		prev.Next = nil
	}

	tree := &TreeNode{slow.Val, nil, nil}
	tree.Left = sortedListToBST(head)
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

func getValFromListNode(head *ListNode) []int {
	vals := make([]int, 0)
	p := head
	for p != nil {
		vals = append(vals, p.Val)
		p = p.Next
	}
	return vals
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
	p := sortedListToBST(createListNode([]int{-10, -3, 0, 5, 9}))
	fmt.Println(getValFromTreeNode(p))
	p = sortedListToBST(createListNode([]int{1, 2, 3, 4}))
	fmt.Println(getValFromTreeNode(p))
	p = sortedListToBST(createListNode([]int{0, 2, -2}))
	fmt.Println(getValFromTreeNode(p))
	p = sortedListToBST(createListNode([]int{0}))
	fmt.Println(getValFromTreeNode(p))
	p = sortedListToBST(createListNode([]int{}))
	fmt.Println(getValFromTreeNode(p))
}
