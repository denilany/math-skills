package stat

import "math"

func StandDev(variance float64) float64 {
	// Find the square root of variance
	return math.Sqrt(variance)
}
