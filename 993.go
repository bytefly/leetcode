package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil || x == y {
		return false
	}

	var queue []*TreeNode
	queue = append(queue, root)

	for {
		size := len(queue)
		cnt, empty, last := 0, 0, -1

		for i := 0; i < size; i++ {
			if queue[i] == nil {
				empty++
				continue
			}

			if queue[i].Val == x || queue[i].Val == y {
				cnt++
				if last < 0 {
					last = i
				} else if i-last > 1 || i%2 == 0 { //find it(in neightbour or live so far)
					return true
				} else { //have same parent
					return false
				}
			}
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}

		//search over
		if empty == size {
			break
		}
		//at least in different depth
		if cnt == 1 {
			return false
		}

		queue = queue[size:]
	}

	return false
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
	/*
		fmt.Println(isCousins(MakeTreeNode([]int{1, 2, 3, 4}), 4, 3))
		fmt.Println(isCousins(MakeTreeNode([]int{1, 2, 3, -1, 4, -1, 5}), 5, 4))
		fmt.Println(isCousins(MakeTreeNode([]int{1, 2, 3, -1, 4}), 2, 3))
	*/
	fmt.Println(isCousins(MakeTreeNode([]int{1, 2, 10, 3, 4, -1, 20, 6, 7, 5, -1, 54, 25, 65, 11, 38, 8, 61, 16, 55, -1, 34, 95, -1, -1, 13, 14, -1, -1, 22, 9, 71, -1, 47, 31, -1, 98, 50, -1, -1, -1, -1, 29, 21, 15, 32, -1, 12, 19, 76, -1, 62, 52, 37, 49, -1, -1, 83, -1, -1, -1, 23, -1, 27, 44, 36, 42, 17, 26, 39, 68, -1, -1, 70, -1, -1, 56, -1, 46, 53, 66, -1, -1, 75, 24, 59, 48, -1, -1, 89, 84, -1, 80, 18, 28, 33, -1, 40, -1, -1, 97, -1, 90, -1, -1, -1, 73, 69, -1, -1, 78, -1, -1, 77, 35, -1, -1, -1, 79, -1, -1, -1, -1, -1, 85, -1, 30, 60, -1, -1, 45, 41, 64, -1, -1, -1, -1, -1, -1, -1, 88, 82, -1, 99, -1, -1, -1, -1, 86, 91, -1, 93, 63, 81, 74, 57, -1, 43, 51, 96, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, 100, 67, -1, -1, -1, -1, -1, -1, -1, 72, -1, 58, -1, -1, -1, -1, -1, -1, -1, -1, -1, 87, 92, 94}), 18, 80))
}
