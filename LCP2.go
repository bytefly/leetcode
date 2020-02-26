package main

import "fmt"

//a0 + 1/(a1/a2)
func subFraction(a0, a1, a2 int) (int, int) {
	return a2 + a0*a1, a1
}

func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	}

	//call subFraction looply
	n, m := subFraction(cont[len(cont)-2], cont[len(cont)-1], 1)
	for i := len(cont) - 3; i >= 0; i-- {
		n, m = subFraction(cont[i], n, m)
	}

	//get GCD for m and n
	a, b := n, m
	gcd := func(int, int) int {
		for {
			if 0 == b {
				break
			}
			a %= b
			if 0 == a {
				break
			}
			b %= a
		}
		return a + b
	}(a, b)

	return []int{n / gcd, m / gcd}
}

func main() {
	fmt.Println(fraction([]int{3}))
	fmt.Println(fraction([]int{0, 3}))
	fmt.Println(fraction([]int{3, 2, 0, 2}))
	fmt.Println(fraction([]int{0, 0, 3}))
}
