package main

import (
	"fmt"
)

func licenseKeyFormatting(S string, K int) string {
	var ans []byte
	var counter int

	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == '-' {
			continue
		}

		if S[i] >= 'a' && S[i] <= 'z' {
			ans = append([]byte{S[i] - 'a' + 'A'}, ans...)
		} else {
			ans = append([]byte{S[i]}, ans...)
		}
		counter++
		if counter == K {
			ans = append([]byte{'-'}, ans...)
			counter = 0
		}
	}

	//drop '-' at the head
	if len(ans) > 1 && ans[0] == '-' {
		ans = ans[1:]
	}
	return string(ans)
}

func main() {
	fmt.Println(licenseKeyFormatting("aF3Z99aaaa", 4))     //"AF-3Z99-AAAA"
	fmt.Println(licenseKeyFormatting("aF3Z99aaaa", 1))     //"A-F-3-Z-9-9-A-A-A-A"
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9nnN-w", 2)) //"5-F3-Z2-E9-NN-NW"
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 4))
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-wii5", 4)) //"5F3-Z2E9-WII5"
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9-w", 1))    //"5-F-3-Z-2-E-9-W"
	fmt.Println(licenseKeyFormatting("5F3Z-2e-9artYUI-wN59M1mY", 4))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 2))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 3))
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 4))    //2-5G3J
	fmt.Println(licenseKeyFormatting("2-5g-3-J", 5))    //25G3J
	fmt.Println(licenseKeyFormatting("2-5g", 5))        //25G
	fmt.Println(licenseKeyFormatting("--a-a-a-a--", 2)) //AA-AA
	fmt.Println(licenseKeyFormatting("----a---", 2))    //A
	fmt.Println(licenseKeyFormatting("-------", 2))     //<EMPTY>
}
