package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {
	var result string
	var calc = (math + science + english + indonesia) / 4
	switch {
	case calc == 100:
		result = "Sempurna"
	case calc >= 90:
		result = "Sangat Baik"
	case calc >= 80:
		result = "Baik"
	case calc > 70:
		result = "Cukup"
	case calc > 60:
		result = "Kurang"
	default:
		result = "Sangat kurang"
	}
	return result
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
