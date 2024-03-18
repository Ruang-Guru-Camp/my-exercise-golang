package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var sampel = map[string]string{"hello": "world", "John": "Doe", "age": "14"}
	MapToSlice(sampel)
}

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

func DeliveryOrder(data []string, day string) map[string]float32 {
	var result = map[string]float32{}
	for _, resData := range data {
		parts := strings.Split(resData, ":")
		name := strings.Join(parts[0:2], "-")
		price, _ := strconv.Atoi(parts[2])
		destination := parts[3]

		if day == "senin" || day == "rabu" || day == "jumat" {
			price = price + (price * 10 / 100)
		} else if day == "selasa" || day == "kamis" || day == "sabtu" {
			price = price + (price * 5 / 100)
		}
		fmt.Println(name, price)

		if destination == "BDG" && day == "rabu" || day == "kamis" || day == "sabtu" {
			result[name] = float32(price)

		} else if destination == "BKS" && day == "selasa" || day == "kamis" || day == "jumat" {
			result[name] = float32(price)

		} else if destination == "DPK" && day == "senin" || day == "selasa" {
			result[name] = float32(price)

		} else if destination == "JKT" && day != "minggu" {
			result[name] = float32(price)
		}

	}
	return result // TODO: replace this
}

// func bintang(n int) {
// 	for i := n; i > 0; i-- {
// 		fmt.Print(strings.Repeat(" ", n-i))
// 		fmt.Print(strings.Repeat("*", i))
// 		fmt.Println("")
// 	}
// }
