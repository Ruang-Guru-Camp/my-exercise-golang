package main

import (
	"fmt"
	"strings"
)

func FindShortestName(names string) string {
	arrNames := strings.FieldsFunc(names, func(r rune) bool {
		return r == ',' || r == ';' || r == ' '
	})
	shortestName := arrNames[0]
	for i := 1; i < len(arrNames); i++ {
		if len(arrNames[i]) < len(shortestName) || (len(arrNames[i]) == len(shortestName) && arrNames[i] < shortestName) {
			shortestName = arrNames[i]
		}
	}

	return shortestName
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
	fmt.Println(FindShortestName("Ari,Aru,Ara,Andi,Asik"))                // "Tia"
}
