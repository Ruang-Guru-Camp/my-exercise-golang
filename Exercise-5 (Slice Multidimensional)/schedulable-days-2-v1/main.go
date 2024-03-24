package main

import "fmt"

func main() {
	var numbers1 = [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	fmt.Println(SchedulableDays(numbers1))
}

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	commonElements := make(map[int]bool)
	for _, num := range villager[0] {
		commonElements[num] = true
	}

	for _, sublist := range villager[1:] {
		sublistElements := make(map[int]bool)
		for _, num := range sublist {
			sublistElements[num] = true
		}

		for num := range commonElements {
			if !sublistElements[num] {
				delete(commonElements, num)
			}
		}
	}

	result := []int{}
	for num := range commonElements {
		result = append(result, num)
	}

	return result
}
