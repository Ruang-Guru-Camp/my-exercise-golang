package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CountVowelConsonant(str string) (int, int, bool) {
	targetCharacters := "aiueo"
	countVocal := 0
	countConsonant := 0

	for _, char := range str {
		if !unicode.IsLetter(char) {
			continue
		}
		if strings.ContainsRune(targetCharacters, unicode.ToLower(char)) {
			countVocal++
		} else {
			countConsonant++
		}
	}
	return countVocal, countConsonant, countVocal < 1
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountVowelConsonant("Hidup Itu Indah"))
	fmt.Println(CountVowelConsonant("kopi"))
	fmt.Println(CountVowelConsonant("bbbbbbbb vvvvvvvv  ddddd sssss  wwww"))
}
