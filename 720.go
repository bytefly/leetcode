package main

import "fmt"

type TrieNode struct {
	IsKey    bool
	Children [26]*TrieNode
}

func CreateTrieNode(words []string) *TrieNode {
	var root TrieNode
	for _, word := range words {
		p := &root
		for i := 0; i < len(word); i++ {
			if p.Children[word[i]-'a'] == nil {
				p.Children[word[i]-'a'] = &TrieNode{}
			}
			if i == len(word)-1 {
				p.Children[word[i]-'a'].IsKey = true
			} else {
				p = p.Children[word[i]-'a']
			}
		}
	}
	return &root
}

func (this TrieNode) Search(word string) bool {
	p := &this
	for i := 0; i < len(word); i++ {
		if p.Children[word[i]-'a'] == nil {
			return false
		}
		if p.Children[word[i]-'a'].IsKey == false {
			return false
		}
		p = p.Children[word[i]-'a']
	}
	return true
}

func longestWord(words []string) string {
	var ans string

	root := CreateTrieNode(words)
	for _, word := range words {
		if !root.Search(word) || len(word) < len(ans) {
			continue
		}
		if len(word) > len(ans) || word < ans {
			ans = word
		}
	}

	return ans
}

func main() {
	fmt.Println(longestWord([]string{"w", "wo", "wor", "worl", "world"}))
	fmt.Println(longestWord([]string{"a", "banana", "app", "appl", "ap", "apply", "apple"}))
	fmt.Println(longestWord([]string{"m", "mo", "moc", "moch", "mocha", "l", "la", "lat", "latt", "latte", "c", "ca", "cat"}))
}
