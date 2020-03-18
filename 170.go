package main

import (
	"fmt"
)

type TwoSum struct {
	table map[int]int
}

/** Initialize your data structure here. */
func Constructor() TwoSum {
	var obj TwoSum
	obj.table = make(map[int]int)
	return obj
}

/** Add the number to an internal data structure.. */
func (this *TwoSum) Add(number int) {
	this.table[number]++
}

/** Find if there exists any pair of numbers which sum is equal to the value. */
func (this *TwoSum) Find(value int) bool {
	for k, _ := range this.table {
		t := value - k
		if (t == k && this.table[t] > 1) ||
			(t != k && this.table[t] > 0) {
			return true
		}
	}
	return false
}

func main() {
	obj := Constructor()
	obj.Add(1)
	obj.Add(3)
	obj.Add(5)
	obj.Add(2)
	fmt.Println(obj.Find(4))
	fmt.Println(obj.Find(7))
	fmt.Println(obj.Find(2))
}
