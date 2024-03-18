package main

import "fmt"

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string
	for key, value := range mapData {
		var temp []string
		temp = append(temp, key)
		temp = append(temp, value)
		result = append(result, temp)
	}
	fmt.Println(result)
	return result // TODO: replace this
}
