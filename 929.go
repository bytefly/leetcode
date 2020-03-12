package main

import (
	"fmt"
	"strings"
)

func numUniqueEmails(emails []string) int {
	m := make(map[string]bool)

	for _, email := range emails {
		pos := strings.IndexByte(email, '@')
		local := strings.ReplaceAll(email[:pos], ".", "")
		plusPos := strings.IndexByte(local, '+')
		if plusPos >= 0 {
			m[local[:plusPos]+email[pos:]] = true
		} else {
			m[local+email[pos:]] = true
		}
	}

	return len(m)
}

func main() {
	fmt.Println(numUniqueEmails([]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}))
}
