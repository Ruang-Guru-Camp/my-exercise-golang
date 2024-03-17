package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var resReverse string
	for i := 0; i < len(str); i++ {
		if i == 0 || string(str[i]) == " " {
			resReverse = string(str[i]) + resReverse
		} else {
			if i != len(str)-1 && string(str[i-1]) == " " {
				resReverse = string(str[i]) + resReverse
			} else {
				resReverse = string(str[i]) + "_" + resReverse
			}
		}
	}
	return resReverse
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(ReverseString("Hello World"))
}
