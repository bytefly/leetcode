package main

import "fmt"

func intervalIntersection(A [][]int, B [][]int) [][]int {
	var ans [][]int
	var m, n int

	for m < len(A) && n < len(B) {
		//same start and end
		if A[m][0] == B[n][0] && A[m][1] == B[n][1] {
			ans = append(ans, []int{A[m][0], B[n][1]})
			m++
			n++
		} else if A[m][0] == B[n][0] { //same start
			if A[m][1] > B[n][1] {
				ans = append(ans, []int{A[m][0], B[n][1]})
				n++
			} else {
				ans = append(ans, []int{A[m][0], A[m][1]})
				m++
			}
		} else if A[m][1] == B[n][1] { //same end
			if A[m][0] < B[n][0] {
				ans = append(ans, []int{B[n][0], A[m][1]})
				m++
			} else {
				ans = append(ans, []int{A[m][0], A[m][1]})
				n++
			}
		} else {
			if A[m][0] > B[n][0] && A[m][0] <= B[n][1] && A[m][1] > B[n][1] {
				//B and A have common area
				ans = append(ans, []int{A[m][0], B[n][1]})
				n++
			} else if A[m][0] > B[n][0] && A[m][0] <= B[n][1] && A[m][1] < B[n][1] {
				//A in B
				ans = append(ans, []int{A[m][0], A[m][1]})
				m++
			} else if A[m][0] < B[n][0] && B[n][0] <= A[m][1] && B[n][1] > A[m][1] {
				//A and B have common area
				ans = append(ans, []int{B[n][0], A[m][1]})
				m++
			} else if A[m][0] < B[n][0] && B[n][0] <= A[m][1] && B[n][1] < A[m][1] {
				//B in A
				ans = append(ans, []int{B[n][0], B[n][1]})
				n++
			} else if A[m][0] > B[n][0] && A[m][0] > B[n][1] {
				//B and A have no intersection
				n++
			} else {
				//A and B have no intersection
				m++
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(intervalIntersection([][]int{
		[]int{0, 2}, []int{5, 10}, []int{13, 23}, []int{24, 25},
	},
		[][]int{
			[]int{1, 5}, []int{8, 12}, []int{15, 24}, []int{25, 26},
		}))
	fmt.Println(intervalIntersection([][]int{[]int{3, 10}}, [][]int{[]int{5, 10}}))
	fmt.Println(intervalIntersection([][]int{
		[]int{3, 5}, []int{9, 20},
	},
		[][]int{
			[]int{4, 5}, []int{7, 10}, []int{11, 12}, []int{14, 15}, []int{16, 20},
		}))
}
