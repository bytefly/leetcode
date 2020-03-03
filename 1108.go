package main

import "fmt"

func defangIPaddr(address string) string {
	var ans []byte
	for _, b := range address {
		if b == '.' {
			ans = append(ans, []byte("[.]")...)
		} else {
			ans = append(ans, byte(b))
		}
	}

	return string(ans)
}

func main() {
	fmt.Println(defangIPaddr("1.1.1.1"))
	fmt.Println(defangIPaddr("255.100.50.0"))
}
