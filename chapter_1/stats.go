package chapter_1

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func Stats() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}
	var min, max float64
	var initialized = 0

	nValues := 0
	var sum float64
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}

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
}
