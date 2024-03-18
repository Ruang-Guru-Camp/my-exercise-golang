package main

import (
	"fmt"
	"strconv"
)

func PhoneNumberChecker(number string, result *string) {
	// TODO: answer here
	if string(number[0:3]) != "628" && string(number[0:2]) != "08" {
		*result = "invalid"
	} else {

		var code int
		if number[0:3] == "628" {
			if len(number) < 11 {
				*result = "invalid"
				return
			}
			code, _ = strconv.Atoi(number[3:5])
		} else {
			if len(number) < 10 {
				*result = "invalid"
				return
			}
			code, _ = strconv.Atoi(number[2:4])
		}
		// fmt.Println(code)
		for i := 11; i < 88; i++ {
			switch {
			case code >= 11 && code <= 15:
				*result = "Telkomsel"
			case code >= 16 && code <= 19:
				*result = "Indosat"
			case code >= 21 && code <= 23:
				*result = "XL"
			case code >= 27 && code <= 29:
				*result = "Tri"
			case code >= 52 && code <= 53:
				*result = "AS"
			case code >= 81 && code <= 88:
				*result = "Smartfren"
			default:
				*result = "invalid"
			}
		}
	}
}

func main() {
	// bisa digunakan untuk pengujian test case
	var number = "081234567"
	var result string

	PhoneNumberChecker(number, &result)
	fmt.Println(result)
}
