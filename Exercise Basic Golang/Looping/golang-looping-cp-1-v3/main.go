package main

import "fmt"

func CountingNumber(n int) float64 {
	var result float64 = 0.0
	for i := 1; i <= n; i++ {
		result += float64(i)
		if i < n {
			result += float64(i) + 0.5
		}
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
