package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	var min int
	for i, e := range nums {
		if i == 0 || e < min {
			min = e
		}
	}
	return min // TODO: replace this
}

func FindMax(nums ...int) int {
	var max int
	for i, e := range nums {
		if i == 0 || e > max {
			max = e
		}
	}
	return max // TODO: replace this
}

func SumMinMax(nums ...int) int {

	return FindMin(nums...) + FindMax(nums...) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
