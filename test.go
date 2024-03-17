package main

import "fmt"

func main() {
	fmt.Println(ExchangeCoin(1234))
}

func ExchangeCoin(amount int) []int {
	hasil := []int{}
	pecahanan1000 := amount / 1000
	sisa1000 := amount - 1000*pecahanan1000
	for i := 0; i < pecahanan1000; i++ {
		hasil = append(hasil, 1000)
	}

	pecahanan500 := sisa1000 / 500
	sisa500 := sisa1000 - 500*pecahanan500
	for i := 0; i < pecahanan500; i++ {
		hasil = append(hasil, 500)
	}

	pecahanan200 := sisa500 / 200
	sisa200 := sisa500 - 200*pecahanan200
	for i := 0; i < pecahanan200; i++ {
		hasil = append(hasil, 200)
	}

	pecahanan50 := sisa200 / 50
	sisa50 := sisa200 - 50*pecahanan50
	for i := 0; i < pecahanan50; i++ {
		hasil = append(hasil, 50)
	}
	pecahanan20 := sisa50 / 20
	sisa20 := sisa50 - 50*pecahanan20
	for i := 0; i < pecahanan20; i++ {
		hasil = append(hasil, 20)
	}

	pecahanan10 := sisa20 / 10
	sisa10 := sisa20 - 10*pecahanan10
	for i := 0; i < pecahanan10; i++ {
		hasil = append(hasil, 10)
	}

	pecahanan5 := sisa10 / 5
	sisa5 := sisa10 - 10*pecahanan5
	for i := 0; i < pecahanan5; i++ {
		hasil = append(hasil, 5)
	}

	pecahanan1 := sisa5
	for i := 0; i < pecahanan1; i++ {
		hasil = append(hasil, 1)
	}

	// fmt.Println(hasil)
	return hasil
}
