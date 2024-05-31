package stat

import "sort"

func Median(numbers []float64) float64 {
	sort.Float64s(numbers)
	lenArr := len(numbers)
	var sum, median float64

	// Check for even length of slice
	if lenArr%2 == 0 {
		index := (lenArr / 2) - 1
		sum = numbers[index] + numbers[index+1]
		median = sum / 2

	} else if lenArr%2 == 1 { // Check for an odd length of the slice
		index := lenArr / 2
		median = numbers[index]
	}

	return median
}
