package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	var result string
	for _, res := range data {
		if strings.Contains(res, input) {
			if len(result) > 0 {
				result += ","
			}
			result += res
		}
	}
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("motor", "mobil APV", "mobil Avanza", "motor matic", "motor gede", "iphone 14", "iphone 13", "iphone 12", "pengering baju", "Kemeja flannel"))
}
