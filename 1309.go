package main

import (
	"fmt"
	"strconv"
)

func freqAlphabets(s string) string {
	var ans []byte

	for i := 0; i < len(s); {
		if i+2 < len(s) && s[i+2] == '#' {
			n, _ := strconv.Atoi(s[i : i+2])
			ans = append(ans, 'a'+byte(n-1))
			i += 3
			continue
		}
		ans = append(ans, 'a'+s[i]-'1')
		i++
	}
	return string(ans)
}

func main() {
	fmt.Println(freqAlphabets("10#11#12"))
	fmt.Println(freqAlphabets("1326#"))
	fmt.Println(freqAlphabets("25#"))
	fmt.Println(freqAlphabets("12345678910#11#12#13#14#15#16#17#18#19#20#21#22#23#24#25#26#"))
}
