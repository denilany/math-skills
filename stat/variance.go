package stat

func Variance(numbers []float64, mean float64) float64 {
	// Variance = sum{(val-mean)^2}
	lenArr := len(numbers)

	var sum float64
	for _, val := range numbers {
		sum += (val - mean) * (val - mean) // For smaller calculations multipying a number by itself is much faster than using math.Pow() function
	}
	return sum / float64(lenArr)
}
