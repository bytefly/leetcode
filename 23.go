package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var p, q, t *ListNode
	dummy := &ListNode{0, nil}

	if len(lists) == 0 {
		return nil
	}
	for {
		if len(lists) == 1 {
			break
		}

		for i := 0; i < len(lists); i += 2 {
			if i+1 == len(lists) {
				lists[i/2] = lists[i]
				break
			}

			p, q = lists[i], lists[i+1]
			t = dummy
			for p != nil && q != nil {
				if p.Val <= q.Val {
					t.Next = p
					p = p.Next
				} else {
					t.Next = q
					q = q.Next
				}
				t = t.Next
			}
			if p != nil {
				t.Next = p
			} else {
				t.Next = q
			}
			lists[i/2] = dummy.Next
		}

		if len(lists)%2 == 0 {
			lists = lists[:len(lists)/2]
		} else {
			lists = lists[:len(lists)/2+1]
		}
	}

	return lists[0]
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
	lists := []*ListNode{createListNode([]int{1, 4, 5}),
		createListNode([]int{1, 3, 4}),
		createListNode([]int{2, 6})}
	//lists := []*ListNode{createListNode([]int{1, 4, 5})}
	//lists := []*ListNode{createListNode([]int{})}
	//lists := []*ListNode{}*/
	fmt.Println(getValFromListNode(mergeKLists(lists)))
}
