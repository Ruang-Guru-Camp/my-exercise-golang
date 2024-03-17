package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	ns := strconv.Itoa(numbers)
	bigest := 0
	pair := 0
	for i := 1; i < len(ns); i++ {
		num1, _ := strconv.Atoi(string(ns[i]))
		num2, _ := strconv.Atoi(string(ns[i-1]))

		if num1+num2 > bigest {
			bigest = num1 + num2
			pairNumb, _ := strconv.Atoi(string(ns[i-1]) + string(ns[i]))
			pair = pairNumb
		}
	}
	return pair
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
