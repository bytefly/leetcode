package main

import (
	"fmt"
)

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	val := image[sr][sc]
	bakupColor, newColor := newColor, -1
	image[sr][sc] = newColor
	//up
	if sr > 0 && image[sr-1][sc] == val {
		floodFill(image, sr-1, sc, newColor)
	}
	//left
	if sc > 0 && image[sr][sc-1] == val {
		floodFill(image, sr, sc-1, newColor)
	}
	//right
	if sc < len(image[0])-1 && image[sr][sc+1] == val {
		floodFill(image, sr, sc+1, newColor)
	}
	//down
	if sr < len(image)-1 && image[sr+1][sc] == val {
		floodFill(image, sr+1, sc, newColor)
	}

	//don't change color in recursive running
	if bakupColor != newColor {
		for i := 0; i < len(image); i++ {
			for j := 0; j < len(image[0]); j++ {
				if image[i][j] == newColor {
					image[i][j] = bakupColor
				}
			}
		}
	}
	return image
}

func main() {
	fmt.Println(floodFill([][]int{[]int{1, 1, 1}, []int{1, 1, 0}, []int{1, 0, 1}}, 1, 1, 2))
	fmt.Println(floodFill([][]int{[]int{0, 0, 0}, []int{0, 1, 1}}, 1, 1, 1))
}
