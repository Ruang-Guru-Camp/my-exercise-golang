package main

import "fmt"

func main() {
	var test = [][][2]int{
		{{1000, 500}, {500, 150}, {600, 100}, {800, 750}},
	}
	var test2 = [][][2]int{
		{{1000, 200}}, {{500, 100}}, {{600, 100}}, {{450, 150}}, {{100, 50}},
	}
	var test3 = [][][2]int{}
	fmt.Println(CountProfit(test))
	fmt.Println(CountProfit(test2))
	fmt.Println(CountProfit(test3))
}

func CountProfit(data [][][2]int) []int {
	if len(data) == 0 {
		return []int{}
	}

	if len(data[0]) == 0 {
		return []int{}
	}

	profits := make([]int, len(data[0]))

	for _, branch := range data {
		for monthIndex, month := range branch {
			profit := month[0] - month[1]
			profits[monthIndex] += profit
		}
	}

	return profits
}
