package chapter_02

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

func Stats() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Generating arguments...")
		for i := 0; i < 10; i++ {
			arguments = append(arguments, strconv.FormatFloat(generateFloat(0, 100), 'f', -1, 64))
		}
		fmt.Println("Generated:", arguments[1:])
	}
	var min, max float64
	var initialized = 0

	nValues := 0
	var values []float64
	var sum float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

		values = append(values, n)
		nValues += 1
		sum += n

		if initialized == 0 {
			min = n
			max = n
			initialized = 1
			continue
		}
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}

	fmt.Println("Number of values:", nValues)
	fmt.Println("Min", min)
	fmt.Println("Min", max)

	if nValues == 0 {
		return
	}
	// Mean
	meanValue := sum / float64(nValues)
	fmt.Printf("Mean value: %5f\n", meanValue)

	// Standard Deviation
	var squared float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		squared += math.Pow(n-meanValue, 2)
	}
	standardDeviation := math.Sqrt(squared / float64(nValues))
	fmt.Printf("Standard deviation: %.5f\n", standardDeviation)

	normalized := normalize(values, meanValue, standardDeviation)
	fmt.Printf("Normalized data: %f\n", normalized)
}

func normalize(data []float64, mean float64, stdDev float64) []float64 {
	if stdDev == 0 {
		return data
	}

	normalized := make([]float64, len(data))
	for i, val := range data {
		// floor(x*100000)/100000 -> 5 digits precision
		normalized[i] = math.Floor((val-mean)/stdDev*100000) / 100000
	}
	return normalized
}

func generateFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
