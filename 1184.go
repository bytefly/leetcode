package main

import (
	"fmt"
)

func distanceBetweenBusStops(distance []int, start int, destination int) int {
	if start > destination {
		start, destination = destination, start
	}

	total, direct := 0, 0
	for i := 0; i < len(distance); i++ {
		total += distance[i]
		if i >= start && i < destination {
			direct += distance[i]
		}
	}
	if direct < total-direct {
		return direct
	}
	return total - direct
}

func main() {
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 1))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 2))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 0, 3))
	fmt.Println(distanceBetweenBusStops([]int{1, 2, 3, 4}, 3, 0))
}
