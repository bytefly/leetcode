package main

import "fmt"

func addToArrayForm(A []int, K int) []int {
	var ans []int
	var t, carry int

	for i := len(A) - 1; i >= 0; i-- {
		t = carry + A[i]
		if K > 0 {
			t = K%10 + carry + A[i]
		}

		if t > 9 {
			t -= 10
			carry = 1
		} else {
			carry = 0
		}

		ans = append([]int{t}, ans...)
		if K > 0 {
			K /= 10
		}
	}

	for K > 0 {
		t = K%10 + carry
		if t > 9 {
			t -= 10
			carry = 1
		} else {
			carry = 0
		}

		ans = append([]int{t}, ans...)
		K /= 10
	}

	if carry == 1 {
		ans = append([]int{1}, ans...)
	}
	return ans
}

func main() {
	fmt.Println(addToArrayForm([]int{1, 2, 0, 0}, 34))
	fmt.Println(addToArrayForm([]int{2, 7, 4}, 181))
	fmt.Println(addToArrayForm([]int{2, 1, 5}, 806))
	fmt.Println(addToArrayForm([]int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9}, 1))
}
