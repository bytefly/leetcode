package main

import (
	"fmt"
	"strconv"
)

type StringIterator struct {
	S     string
	Pos   int
	Cnt   int
	Total int
}

func Constructor(compressedString string) StringIterator {
	var total, i int
	if compressedString != "" {
		for i = 1; i < len(compressedString); i++ {
			if compressedString[i] < '0' || compressedString[i] > '9' {
				break
			}
		}

		total, _ = strconv.Atoi(compressedString[1:i])
	}
	return StringIterator{compressedString, 0, 0, total}
}

func (this *StringIterator) Next() byte {
	var i int

	if this.Pos == len(this.S) {
		return ' '
	}

	ans := this.S[this.Pos]
	this.Cnt++
	if this.Cnt == this.Total {
		t := strconv.Itoa(this.Total)
		this.Pos += (len(t) + 1)
		this.Cnt = 0
		for i = this.Pos + 1; i < len(this.S); i++ {
			if this.S[i] < '0' || this.S[i] > '9' {
				break
			}
		}

		if this.Pos == len(this.S) {
			this.Total = 0
		} else {
			this.Total, _ = strconv.Atoi(this.S[this.Pos+1 : i])
		}
	}

	return ans
}

func (this *StringIterator) HasNext() bool {
	return this.Pos < len(this.S)
}

func main() {
	obj := Constructor("L1e2t1C1o1d1e1")
	for obj.HasNext() {
		fmt.Printf("%c", obj.Next())
	}
	fmt.Printf("%c", obj.Next())
	fmt.Println()
}
