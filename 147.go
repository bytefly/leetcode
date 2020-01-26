package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	p := head
	for p != nil {
        var prev, q *ListNode
		t := p.Next
		//find the first node larger than p
		for q = head; q != p; {
			if q.Val > p.Val {
				break
			}
			prev = q
			q = q.Next
		}
		//swap p and q
		if q != p {
			if prev != nil {
				prev.Next = p
			} else { //head replaced
				//reset head when the head changed
				head = p
			}

			//find the last node before p, then link to the p.Next
			for last := q; last != p; {
				if last.Next == p {
					last.Next = p.Next
					break
				} else {
					last = last.Next
				}
			}
			p.Next = q
			//fmt.Println(getValFromListNode(head))
		}

		p = t
	}

	return head
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
	head := createListNode([]int{3, 4, 1})
	//head := createListNode([]int{-1, 5, 3, 4, 0})
	//head := createListNode([]int{4, 2, 1, 3})
	//head := createListNode([]int{3, 1})
	//head := createListNode([]int{3})
	head = insertionSortList(head)
	fmt.Println(getValFromListNode(head))
}
