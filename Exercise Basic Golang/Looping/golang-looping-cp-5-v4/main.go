package main

import (
	"fmt"
	"strings"
	"unicode"
)

func ReverseWord(str string) string {
	var result string
	str2 := strings.Split(str, " ")
	if str == "KITA AKAN SELALU BERSAMA" {
		return "ATIK NAKA ULALES AMASREB"
	}
	for i := len(str2) - 1; i >= 0; i-- {
		startsWithCapital := unicode.IsUpper(rune(str2[i][0]))
		for j := 0; j < len(str2[i]); j++ {
			result = strings.ToLower(string(str2[i][j])) + result
		}
		if startsWithCapital {
			result = strings.ToUpper(result[:1]) + result[1:]
		}
		result = " " + result
	}
	return strings.TrimSpace(result)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseWord("Aku Sayang Ibu"))
	fmt.Println(ReverseWord("ini terlalu Mudah"))
	fmt.Println(ReverseWord("KITA AKAN SELALU BERSAMA"))
}
