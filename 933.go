package main

import "fmt"

type RecentCounter struct {
	window []int
}

func Constructor() RecentCounter {
	var obj RecentCounter
	return obj
}

func (this *RecentCounter) Ping(t int) int {
	var cnt int

	this.window = append(this.window, t)
	for i := 0; i < len(this.window); i++ {
		if t-this.window[i] <= 3000 {
			break
		}
		cnt++
	}
	this.window = this.window[cnt:]

	return len(this.window)
}

func main() {
	obj := Constructor()
	fmt.Println(obj.Ping(1))
	fmt.Println(obj.Ping(100))
	fmt.Println(obj.Ping(3001))
	fmt.Println(obj.Ping(3002))
}
