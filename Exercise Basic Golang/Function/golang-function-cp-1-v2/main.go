package main

import (
	"fmt"
	"strconv"
)

func DateFormat(day, month, year int) string {
	resDay := fmt.Sprintf("%02d", day)
	var resMonth string
	switch month {
	case 1:
		resMonth = "January"
	case 2:
		resMonth = "February"
	case 3:
		resMonth = "March"
	case 4:
		resMonth = "April"
	case 5:
		resMonth = "May"
	case 6:
		resMonth = "June"
	case 7:
		resMonth = "July"
	case 8:
		resMonth = "August"
	case 9:
		resMonth = "September"
	case 10:
		resMonth = "October"
	case 11:
		resMonth = "November"
	case 12:
		resMonth = "December"
	default:
		resMonth = "Unknown"
	}
	result := resDay + "-" + resMonth + "-" + strconv.Itoa(year)
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))
}
