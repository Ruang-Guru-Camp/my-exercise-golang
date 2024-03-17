package main

import (
	"fmt"
	"strings"
)

func CountingLetter(text string) int {
	targetCharacters := []string{"r", "s", "t", "z"}
	totCount := 0
	for i := 0; i < len(text); i++ {
		for _, char := range targetCharacters {
			if strings.ToLower(string(text[i])) == char {
				totCount++
				break
			}
		}
	}

	return totCount
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
