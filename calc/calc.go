package calc

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"math-skills/stat"
)

func Calc(str string) (string, error) {
	file, err := os.Open(str)
	if err != nil {
		return "", fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	// Check if the file is empty
	fileInfo, err := file.Stat()
	if err != nil {
		return "", fmt.Errorf("error getting file info: %s", err)
	}
	if fileInfo.Size() == 0 {
		return "", fmt.Errorf("%s file is empty", str)
	}

	var arr []float64
	scanner := bufio.NewScanner(file)

	lineNum := 0 // Track line number
	for scanner.Scan() {
		scanLine := scanner.Text()
		line := strings.TrimSpace(scanLine)
		lineNum++

		if line == "" {
			continue
		}

		val, err := strconv.ParseFloat(line, 64)
		if err != nil {
			return "", fmt.Errorf("value at line %d in the %s contains an invalid character: %v", lineNum, str, err)
		}
		// Check if the value in the file is greater than 9223372036854775807
		if val >= math.MaxInt64 {
			return "", fmt.Errorf("line %d has a value of %.0f which is extremely big", lineNum, val)
		}
		if val < 0 {
			return "", fmt.Errorf("population data cannot be a negative number, %.0f", val)
		}
		arr = append(arr, val)
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("error scanning text %s", err)
	}

	// Check for file with just empty spaces
	if len(arr) == 0 {
		return "", fmt.Errorf("file data appears to be empty")
	}

	// Check for single population data
	if len(arr) == 1 {
		return "", fmt.Errorf("population data cannot be calculated from a single value")
	}

	// Calculate statistics
	mean := stat.Average(arr)
	median := stat.Median(arr)
	variance := stat.Variance(arr, mean)
	stdDev := stat.StandDev(variance)

	// Round the statistics
	roundedMean := math.Round(mean)
	roundedMedian := math.Round(median)
	roundedVariance := math.Round(variance)
	roundedStdDev := math.Round(stdDev)

	// Format the output, while checking for overflows in the results
	var expectedResult string

	if roundedMean >= math.MaxInt64 {
		expectedResult += "Value of average is extremely big\n"
	} else {
		expectedResult += fmt.Sprintf("Average: %.0f\n", roundedMean)
	}

	if roundedMedian >= math.MaxInt64 {
		expectedResult += "Value of median is extremely big\n"
	} else {
		expectedResult += fmt.Sprintf("Median: %.0f\n", roundedMedian)
	}

	if roundedVariance >= math.MaxInt64 {
		expectedResult += "Value of variance is extremely big\n"
	} else {
		expectedResult += fmt.Sprintf("Variance: %.0f\n", roundedVariance)
	}

	if roundedStdDev >= math.MaxInt64 {
		expectedResult += "Value of standard deviation is extremely big\n"
	} else {
		expectedResult += fmt.Sprintf("Standard Deviation: %.0f\n", roundedStdDev)
	}

	return expectedResult, nil
}
