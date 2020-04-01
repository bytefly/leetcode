package main

import "fmt"

type Node struct {
	id   int
	dist int
}

func kClosest(points [][]int, K int) [][]int {
	var ans [][]int
	var cnt int

	nodes := make([]Node, len(points))
	for i := 0; i < len(nodes); i++ {
		nodes[i] = Node{i, points[i][0]*points[i][0] + points[i][1]*points[i][1]}
	}

	n := len(nodes)
	for i := n/2 - 1; i >= 0; i-- {
		heaplize(nodes, i, n)
	}

	for i := n - 1; i >= 0; i-- {
		nodes[0], nodes[i] = nodes[i], nodes[0]
		if cnt < K {
			ans = append(ans, points[nodes[i].id])
			cnt++
		} else {
			break
		}
		heaplize(nodes, 0, i)
	}

	return ans
}

func heaplize(nodes []Node, i, end int) {
	small, left, right := i, 2*i+1, 2*i+2
	if left < end && nodes[left].dist < nodes[small].dist {
		small = left
	}
	if right < end && nodes[right].dist < nodes[small].dist {
		small = right
	}

	if small != i {
		nodes[small], nodes[i] = nodes[i], nodes[small]
		heaplize(nodes, small, end)
	}
}

func main() {
	fmt.Println(kClosest([][]int{
		[]int{1, 3}, []int{-2, 2},
	}, 1))
	fmt.Println(kClosest([][]int{
		[]int{1, 3},
	}, 1))
	fmt.Println(kClosest([][]int{
		[]int{3, 3}, []int{5, -1}, []int{-2, 4},
	}, 2))
}
