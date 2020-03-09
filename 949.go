package main

import "fmt"

func permutation(A []int) [][]int {
	if len(A) == 1 {
		return [][]int{[]int{A[0]}}
	}

	var ans [][]int
	for i := 0; i < len(A); i++ {
		t := permutation(append(append([]int{}, A[:i]...), A[i+1:]...))
		for _, s := range t {
			s = append(s, A[i])
			ans = append(ans, s)
		}
	}

	return ans
}

func largestTimeFromDigits(A []int) string {
	ans := []byte("")
	p := permutation(A)

	for _, t := range p {
		hour, min := t[0]*10+t[1], t[2]*10+t[3]
		if hour < 24 && min < 60 {
			s := fmt.Sprintf("%02d:%02d", hour, min)
			if s > string(ans) {
				ans = []byte(s)
			}
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(largestTimeFromDigits([]int{1, 2, 3, 4}))
	fmt.Println(largestTimeFromDigits([]int{5, 5, 5, 5}))
	fmt.Println(largestTimeFromDigits([]int{9, 0, 0, 0}))
	fmt.Println(largestTimeFromDigits([]int{2, 0, 6, 6}))
	fmt.Println(largestTimeFromDigits([]int{0, 2, 7, 6}))
}
