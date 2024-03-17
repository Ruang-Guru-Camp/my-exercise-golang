package main

import (
	"fmt"
)

func main() {
	value1 := []int{172, 170, 150, 165, 144, 155, 159}
	fmt.Println(Sortheight(value1))
}

func Sortheight(height []int) []int {
	sortedHeight := make([]int, len(height))
	copy(sortedHeight, height)

	for i := 0; i < len(sortedHeight)-1; i++ {
		for j := 0; j < len(sortedHeight)-i-1; j++ {
			if sortedHeight[j] > sortedHeight[j+1] {
				sortedHeight[j], sortedHeight[j+1] = sortedHeight[j+1], sortedHeight[j]
			}
		}
	}

	return sortedHeight
}
