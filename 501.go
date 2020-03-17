package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var ans []int
	var stack []*TreeNode
	var prev, cur, max, cnt int

	p := root
	for {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}

		if len(stack) == 0 {
			if cnt >= max && cnt > 0 {
				if cnt > max {
					ans = ans[:0]
				}
				ans = append(ans, prev)
			}
			break
		}

		for len(stack) > 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			cur = p.Val
			if cnt == 0 {
				prev = cur
				cnt++
			} else if cur == prev {
				cnt++
			} else if cur != prev {
				if cnt >= max {
					if cnt > max {
						ans = ans[:0]
						max = cnt
					}
					ans = append(ans, prev)
				}
				prev = cur
				cnt = 1
			}

			p = p.Right
			if p != nil {
				break
			}
		}
	}

	return ans
}

func MakeTreeNode(vals []int) *TreeNode {
	var root, node0, node1, node2 *TreeNode
	queue := make([]*TreeNode, 0)
	i := 0

	for {
		if i == 0 {
			if i >= len(vals) {
				return root
			}
			node0 = &TreeNode{vals[i], nil, nil}
			root = node0
			queue = append(queue, node0)
			i++
			continue
		}

		node0Num := len(queue)
		for j := 0; j < node0Num; j++ {
			node0 = queue[j]

			if i >= len(vals) {
				return root
			}
			if node0 == nil {
				//no children parse
				continue
			}

			if vals[i] == -1 {
				node1 = nil
			} else {
				node1 = &TreeNode{vals[i], nil, nil}
			}
			node0.Left = node1
			queue = append(queue, node1)

			if i == len(vals)-1 || vals[i+1] == -1 {
				node2 = nil
			} else {
				node2 = &TreeNode{vals[i+1], nil, nil}
			}
			node0.Right = node2
			queue = append(queue, node2)

			i += 2
		}
		queue = queue[node0Num:]
	}

	return root
}

func main() {
	fmt.Println(findMode(MakeTreeNode([]int{1, -1, 2, 2})))
	fmt.Println(findMode(MakeTreeNode([]int{4, 2, 6, 1, 3, 5, 7})))
	fmt.Println(findMode(MakeTreeNode([]int{})))
	fmt.Println(findMode(MakeTreeNode([]int{1})))
	fmt.Println(findMode(MakeTreeNode([]int{6, 2, 8, 0, 4, 7, 9, -1, -1, 2, 6})))
}
