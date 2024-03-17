package main

import (
	"fmt"
	"strings"
)

func SlurredTalk(words *string) {
	wordList := strings.Fields(*words)

	for i, word := range wordList {
		word = strings.Replace(word, "s", "l", -1)
		word = strings.Replace(word, "r", "l", -1)
		word = strings.Replace(word, "z", "l", -1)

		word = strings.Replace(word, "S", "L", -1)
		word = strings.Replace(word, "R", "L", -1)
		word = strings.Replace(word, "Z", "L", -1)

		wordList[i] = word
	}
	*words = strings.Join(wordList, " ")
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Saya Steven, saya suka menggoreng telur dan suka hewan zebra"
	SlurredTalk(&words)
	fmt.Println(words)
}
