package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func closestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}

		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closest
}

func GetTreeNodeVal(root *TreeNode) []int {
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	vals := make([]int, 0)
	for len(stack) > 0 {
		size := len(stack)
		nilNum := 0
		for i := 0; i < size; i++ {
			p := stack[i]
			if p != nil {
				vals = append(vals, p.Val)
				stack = append(stack, p.Left)
				stack = append(stack, p.Right)
			} else {
				vals = append(vals, -1)
				nilNum++
			}
		}
		//trim all nil leaf nodes
		if nilNum == size {
			vals = vals[:len(vals)-size]
		}
		stack = stack[size:]
	}
	return vals
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

			if i >= len(vals)-1 || vals[i+1] == -1 {
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
	fmt.Println(closestValue(MakeTreeNode([]int{4, 2, 5, 1, 3}), 3.714286))
	fmt.Println(closestValue(MakeTreeNode([]int{4, 2, 5, 1, 3}), 4.714286))
	fmt.Println(closestValue(MakeTreeNode([]int{4, 2, 5, 1, 3}), 5.714286))
	fmt.Println(closestValue(MakeTreeNode([]int{4, 2, 5, 1, 3}), 1.714286))
	fmt.Println(closestValue(MakeTreeNode([]int{41, 37, 44, 24, 39, 42, 48, 1, 35, 38, 40, -1, 43, 46, 49, 0, 2, 30, 36, -1, -1, -1, -1, -1, -1, 45, 47, -1, -1, -1, -1, -1, 4, 29, 32, -1, -1, -1, -1, -1, -1, 3, 9, 26, -1, 31, 34, -1, -1, 7, 11, 25, 27, -1, -1, 33, -1, 6, 8, 10, 16, -1, -1, -1, 28, -1, -1, 5, -1, -1, -1, -1, -1, 15, 19, -1, -1, -1, -1, 12, -1, 18, 20, -1, 13, 17, -1, -1, 22, -1, 14, -1, -1, 21, 23}), 3.285714))
}
