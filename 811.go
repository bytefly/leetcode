package main

import (
	"fmt"
	"strconv"
	"strings"
)

func subdomainVisits(cpdomains []string) []string {
	var ans []string
	m := make(map[string]int)

	for _, s := range cpdomains {
		pos := strings.IndexByte(s, ' ')
		cnt, _ := strconv.Atoi(s[:pos])
		pos2 := pos
		pos++
		for pos2 > 0 {
			m[s[pos:]] += cnt
			pos2 = strings.IndexByte(s[pos:], '.')
			pos += pos2 + 1
		}
	}
	for k, v := range m {
		ans = append(ans, fmt.Sprintf("%d %s", v, k))
	}

	return ans
}

func main() {
	fmt.Println(subdomainVisits([]string{"9001 discuss.leetcode.com"}))
	fmt.Println(subdomainVisits([]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}))
}
