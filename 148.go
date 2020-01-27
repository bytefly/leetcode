package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := &ListNode{0, head}
	intv, length := 1, 0
	l, r := head, head
	for r != nil {
		r = r.Next
		length++
	}

	for intv < length {
		p := dummy
		for {
			l, r = p.Next, p.Next
			for i := 0; i < intv; i++ {
				if r == nil {
					break
				}
				r = r.Next
			}

			//merge left and right ordered list together
			lnum, rnum := 0, 0
			for lnum < intv && rnum < intv {
				if l == nil || r == nil {
					break
				}
				if l.Val <= r.Val {
					p.Next = l
					p, l = p.Next, l.Next
					lnum++
				} else {
					p.Next = r
					p, r = p.Next, r.Next
					rnum++
				}
			}
            //p always point to the last ordered node
            //link the remain of r
			if rnum < intv && r != nil {
				p.Next = r
                //iterate to the end of r
				for ; rnum < intv; rnum++ {
                    //r is short than intv
                    if p.Next == nil {
                        break
                    }
					p = p.Next
				}
			}
            //link the remain of l
			if lnum < intv {
				q := p.Next
				p.Next = l
                //iterate to the end of l
				for ; lnum < intv; lnum++ {
                    //l is short than intv
                    if p.Next == nil {
                        break
                    }
					p = p.Next
				}
				if r == nil {
					p.Next = nil
				} else {
					p.Next = q
				}
			}

			if r == nil {
				break
			}
		}

		head = dummy.Next
		intv *= 2
		//fmt.Println(getValFromListNode(head))
	}

	return dummy.Next
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
	//head := createListNode([]int{3, 2, 4})
	//head := createListNode([]int{3, 4, 1})
	head := createListNode([]int{-1, 5, 2, 3, 0})
	//head := createListNode([]int{4, 2, 1, 3})
	//head := createListNode([]int{3, 1})
	//head := createListNode([]int{3})
	head = sortList(head)
	fmt.Println(getValFromListNode(head))
}
