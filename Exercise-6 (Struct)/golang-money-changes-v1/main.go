package main

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

	return kembalian // TODO: replace this
}

func main() {
	p1 := []Product{
		{Name: "Baju",
			Price: 5000,
			Tax:   500},
	}
	MoneyChanges(10000, p1)
}
