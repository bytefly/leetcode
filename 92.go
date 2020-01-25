package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m > n {
		return head
	}

	var dummy, end *ListNode
	i := 1
	newHead, p := head, head.Next

	if 0 == m-1 {
		end = head
        //must set to nil else dead-loop in all reverse mode
		end.Next = nil
		dummy = &ListNode{0, head}
	} else {
		dummy = head
	}

	for p != nil {
		if i >= m-1 && i <= n-1 {
			if i == m-1 {
				end = p
                //must set to nil else dead-loop at the last node for n=5
				dummy.Next = nil
			}
			if m == 1 && i == n-1 {
				newHead = p
			}

			q := p.Next
			p.Next = dummy.Next
			dummy.Next = p

			p = q
		} else {
			if i < m-1 {
				dummy.Next = p
				dummy = p
			}
			if i > n-1 {
				end.Next = p
				end = p
			}

			p = p.Next
		}

		i++
	}

	return newHead
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
	head := createListNode([]int{1, 2, 3, 4, 5})
	//head = reverseBetween(head, 2, 4)
	//head = reverseBetween(head, 1, 4)
	//head = reverseBetween(head, 2, 5)
	head = reverseBetween(head, 1, 5)

	vals := getValFromListNode(head)
	fmt.Println(vals)
}
