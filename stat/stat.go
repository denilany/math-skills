package stat

import (
	"math"
	"sort"
)

func Average(numbers []float64) float64 {
	lenArr := len(numbers)
	total := 0.0
	for _, val := range numbers {
		total += val
	}
	mean := total / float64(lenArr)
	return mean
}

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

	// fmt.Println(numbers)
	return median
}

func Variance(numbers []float64, mean float64) float64 {
	// Variance = sum{(val-mean)^2}
	lenArr := len(numbers)

	var sum float64
	for _, val := range numbers {
		sum += (val - mean) * (val - mean) // For smaller calculations multipying a number by itself is much faster than using math.Pow() function
	}
	return sum / float64(lenArr)
}

func StandDev(variance float64) float64 {
	// Find the square root of variance
	return math.Sqrt(variance)
}
