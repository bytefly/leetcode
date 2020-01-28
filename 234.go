package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	ans := true
	if head == nil {
		return false
	}

	//find the middle node
	slow, fast := head, head
	for fast != nil {
		if fast != nil {
			fast = fast.Next
		}
		if fast != nil {
			fast = fast.Next
		}
		if fast != nil {
			slow = slow.Next
		}
	}

	//reverse the next half nodes
	p, q := slow, slow.Next
	slow.Next = nil
	for q != nil {
		t := q.Next
		q.Next = p.Next
		p.Next = q
		q = t
	}

	//compare the two half nodes
	p, q = head, slow.Next
	for p != nil && q != nil {
		if p.Val != q.Val {
			ans = false
			break
		}
		p, q = p.Next, q.Next
	}

	//reverse the next half nodes
	p, q = slow, slow.Next
	slow.Next = nil
	for q != nil {
		t := q.Next
		q.Next = p.Next
		p.Next = q
		q = t
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
	//t := []int{6, 2, 6, 3, 4, 5, 6}
	t := []int{6, 2, 6, 4, 5, 6}
	//t := []int{6}
	//t := []int{6, 6}
	//t := []int{6, 2, 2, 6}
	//t := []int{6, 2, 6}
	head := createListNode(t)
	fmt.Println(isPalindrome(head))
}
