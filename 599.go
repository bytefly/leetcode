package main

import (
	"fmt"
)

func findRestaurant(list1 []string, list2 []string) []string {
	min := len(list1) + len(list2)
	m := make(map[string]int)
	var ans []string

	for i, v := range list1 {
		m[v] = i
	}
	for i, v := range list2 {
		if v, ok := m[v]; ok {
			if i+v < min {
				ans = ans[:0]
				min = i + v
				ans = append(ans, list2[i])
			} else if i+v == min {//other choice with the same index sum result 
				ans = append(ans, list2[i])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Burger King", "Tapioca Express", "Shogun"}))
}
