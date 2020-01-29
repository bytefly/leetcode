package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	ans := make([]int, 0)
	p := head
	for p != nil {
		find := false
		for q := p.Next; q != nil; q = q.Next {
			if q.Val > p.Val {
				ans = append(ans, q.Val)
				find = true
				break
			}
		}
		if !find {
			ans = append(ans, 0)
		}
		p = p.Next
	}
	return ans
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

func main() {
	p := nextLargerNodes(createListNode([]int{2, 1, 5}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{2, 7, 4, 3, 5}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{1, 7, 5, 1, 9, 2, 5, 1}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{2}))
	fmt.Println(p)
	p = nextLargerNodes(createListNode([]int{}))
	fmt.Println(p)
}
