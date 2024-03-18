package main

func Sortheight(height []int) []int {
	sortedHeight := make([]int, len(height))
	copy(sortedHeight, height)

	for i := 0; i < len(sortedHeight)-1; i++ {
		for j := 0; j < len(sortedHeight)-i-1; j++ {
			if sortedHeight[j] > sortedHeight[j+1] {
				sortedHeight[j], sortedHeight[j+1] = sortedHeight[j+1], sortedHeight[j]
			}
		}
	}

	return sortedHeight // TODO: replace this
}
