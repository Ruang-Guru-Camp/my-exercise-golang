package main

func ReverseData(arr [5]int) [5]int {
	var result [5]int
	for i := 0; i < len(arr); i++ {
		reversed := 0
		num := arr[i]
		for num > 0 {
			digit := num % 10
			reversed = reversed*10 + digit
			num /= 10
		}
		result[i] = reversed
	}
	resRev := [5]int{}
	for i := len(result) - 1; i >= 0; i-- {
		resRev[len(result)-1-i] = result[i]
	}
	return [5]int(resRev) // TODO: replace this
}
