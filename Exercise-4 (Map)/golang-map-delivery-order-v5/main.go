package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TODO: answer here
func DeliveryOrder(data []string, day string) map[string]float32 {
	var result = map[string]float32{}
	for _, resData := range data {
		parts := strings.Split(resData, ":")
		name := strings.Join(parts[0:2], "-")
		price, _ := strconv.Atoi(parts[2])
		destination := parts[3]

		var deliveryFee float32
		if day == "senin" || day == "rabu" || day == "jumat" {
			deliveryFee = float32(price) * 10 / 100
		} else if day == "selasa" || day == "kamis" || day == "sabtu" {
			deliveryFee = float32(price) * 5 / 100
		}

		if destination == "BDG" && (day == "rabu" || day == "kamis" || day == "sabtu") {
			result[name] += float32(price) + deliveryFee
		} else if destination == "BKS" && (day == "selasa" || day == "kamis" || day == "jumat") {
			result[name] += float32(price) + deliveryFee
		} else if destination == "DPK" && (day == "senin" || day == "selasa") {
			result[name] += float32(price) + deliveryFee
		} else if destination == "JKT" && day != "minggu" {
			result[name] += float32(price) + deliveryFee
		}
	}
	return result
}

func main() {
	data := []string{
		"Budi:Gunawan:10000:JKT",
		"Andi:Sukirman:20000:JKT",
		"Budi:Sukirman:30000:BDG",
		"Andi:Gunawan:40000:BKS",
		"Budi:Gunawan:50000:DPK",
	}

	day := "sabtu"

	deliveryData := DeliveryOrder(data, day)

	fmt.Println(deliveryData)
}
