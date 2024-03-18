package main

func SchedulableDays(date1 []int, date2 []int) []int {
	result := []int{}
	for _, imam := range date1 {
		for _, permana := range date2 {
			if imam == permana {
				// result[i] = imam
				result = append(result, imam)
			}
		}
	}
	return result // TODO: replace this
}
