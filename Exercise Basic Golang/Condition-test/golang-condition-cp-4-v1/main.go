package main

import "fmt"

func GetTicketPrice(VIP, regular, student, day int) float32 {
	priceTicketVIP := 30
	priceTicketRegular := 20
	priceTicketStudent := 10
	totTicket := VIP + regular + student
	totalPrice := (VIP * priceTicketVIP) + (regular * priceTicketRegular) + (student * priceTicketStudent)

	var discount float32
	if totalPrice >= 100 {
		if day%2 != 0 {
			if totTicket < 5 {
				discount = float32(15) / 100
			} else {
				discount = float32(25) / 100
			}
		} else {
			if totTicket < 5 {
				discount = float32(10) / 100
			} else {
				discount = float32(20) / 100
			}
		}
	} else {
		return float32(totalPrice)
	}

	return float32(totalPrice) - (float32(totalPrice) * discount)
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetTicketPrice(3, 3, 3, 20))
}
