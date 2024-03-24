package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0)
	s = append(s, 1)
	s = append(s, 2, 3, 4)
	s = append(s, 5)
	fmt.Println(s[2:4])
	// Test cases
	// data1 := []string{"Budi;23;Jakarta;;", "Joko;30;Bandung;;true", "Susi;25;Bogor;165.42;"}
	// data2 := []string{"Jaka;25;Jakarta;false;170.1", "Anggi;24;Bandung;;"}

	// output1 := PopulationData(data1)
	// output2 := PopulationData(data2)

	// fmt.Println(output1)
	// fmt.Println(output2)

	// Print results
	// for _, p := range output1 {
	// 	println("Name:", p["name"].(string), "Age:", p["age"].(int), "Address:", p["address"].(string))
	// 	if height, ok := p["height"]; ok {
	// 		println("Height:", height.(float64))
	// 	}
	// 	if isMarried, ok := p["isMarried"]; ok {
	// 		println("Is Married:", isMarried.(bool))
	// 	}
	// }

	// for _, p := range output2 {
	// 	println("Name:", p["name"].(string), "Age:", p["age"].(int), "Address:", p["address"].(string))
	// 	if height, ok := p["height"]; ok {
	// 		println("Height:", height.(float64))
	// 	}
	// 	if isMarried, ok := p["isMarried"]; ok {
	// 		println("Is Married:", isMarried.(bool))
	// 	}
	// }

	// p1 := []Product{
	// 	{Name: "Baju", Price: 5000, Tax: 500}, {Name: "Celana", Price: 3000, Tax: 300},
	// }
	// MoneyChanges(10000, p1)

	// var numbers = [][]int{{7, 12, 19, 22}, {12, 19, 21, 23}, {7, 12, 19}, {12, 19}}
	// var numbers2 = [][]int{{1, 2, 3, 4, 5}, {2, 3, 4, 5}, {2, 3, 4, 10, 11, 12, 15}}
	// fmt.Println(SchedulableDays(numbers))
	// fmt.Println(SchedulableDays(numbers2))
}

func PopulationData(data []string) []map[string]interface{} {
	result := make([]map[string]interface{}, 0)

	for _, d := range data {
		parts := strings.Split(d, ";")

		name := parts[0]
		age, _ := strconv.Atoi(parts[1])
		address := parts[2]

		person := map[string]interface{}{
			"name":    name,
			"age":     age,
			"address": address,
		}

		if len(parts[3]) > 0 {
			height, _ := strconv.ParseFloat(parts[3], 64)
			person["height"] = height
		}

		if len(parts) > 4 && len(parts[4]) > 0 {
			isMarried, _ := strconv.ParseBool(parts[4])
			person["isMarried"] = isMarried
		}

		result = append(result, person)
	}

	return result
}

type Product struct {
	Name  string
	Price int
	Tax   int
}

func MoneyChanges(amount int, products []Product) []int {
	total := 0
	kembalian := []int{}

	for i := 0; i < len(products); i++ {
		total += products[i].Price + products[i].Tax
	}

	amount = amount - total

	pecahanan1000 := amount / 1000
	sisa1000 := amount - 1000*pecahanan1000
	for i := 0; i < pecahanan1000; i++ {
		kembalian = append(kembalian, 1000)
	}

	pecahanan500 := sisa1000 / 500
	sisa500 := sisa1000 - 500*pecahanan500
	for i := 0; i < pecahanan500; i++ {
		kembalian = append(kembalian, 500)
	}

	pecahanan200 := sisa500 / 200
	sisa200 := sisa500 - 200*pecahanan200
	for i := 0; i < pecahanan200; i++ {
		kembalian = append(kembalian, 200)
	}
	pecahanan100 := sisa200 / 100
	sisa100 := sisa200 - 100*pecahanan100
	for i := 0; i < pecahanan100; i++ {
		kembalian = append(kembalian, 100)
	}

	pecahanan50 := sisa100 / 50
	sisa50 := sisa100 - 50*pecahanan50
	for i := 0; i < pecahanan50; i++ {
		kembalian = append(kembalian, 50)
	}
	pecahanan20 := sisa50 / 20
	sisa20 := sisa50 - 20*pecahanan20
	for i := 0; i < pecahanan20; i++ {
		kembalian = append(kembalian, 20)
	}

	pecahanan10 := sisa20 / 10
	sisa10 := sisa20 - 10*pecahanan10
	for i := 0; i < pecahanan10; i++ {
		kembalian = append(kembalian, 10)
	}

	pecahanan5 := sisa10 / 5
	sisa5 := sisa10 - 5*pecahanan5
	for i := 0; i < pecahanan5; i++ {
		kembalian = append(kembalian, 5)
	}

	pecahanan1 := sisa5
	for i := 0; i < pecahanan1; i++ {
		kembalian = append(kembalian, 1)
	}

	fmt.Println(total)
	fmt.Println(kembalian)
	return nil // TODO: replace this
}

func SchedulableDays(villager [][]int) []int {
	if len(villager) == 0 {
		return []int{}
	}

	commonElements := make(map[int]bool)
	for _, num := range villager[0] {
		commonElements[num] = true
	}

	for _, sublist := range villager[1:] {
		sublistElements := make(map[int]bool)
		for _, num := range sublist {
			sublistElements[num] = true
		}

		for num := range commonElements {
			if !sublistElements[num] {
				delete(commonElements, num)
			}
		}
	}

	result := []int{}
	for num := range commonElements {
		result = append(result, num)
	}

	return result

	// for i := 1; i < len(villager); i++ {
	// 	for j := 0; j < len(villager[i]); j++ {
	// 		// fmt.Println("1. ", villager[0][j])
	// 		// fmt.Println("2. ", villager[i-1][j])
	// 		fmt.Println()
	// 		if villager[0][j] == villager[i-1][j] {
	// 			fmt.Println("masuk")
	// 		}
	// 		// fmt.Print(villager[i][j])
	// 	}
	// 	fmt.Println()
	// }
	// for index, data := range villager {
	// 	fmt.Println(villager[index])
	// 	for index2, data2 := range data {
	// 		if villager[index][index2] == villager[index][index2] {
	// 			result = append(result, data[index2])
	// 		}
	// 	}
	// }
	return result // TODO: replace this
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
