package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	// First argument is always the file path (temp file path in case of go run)
	if len(arguments) == 1 {
		fmt.Println("Need one or more arguments!")
		return
	}

	var smallest, biggest float64
	var initialized = 0
	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			continue
		}
		if initialized == 0 {
			smallest = n
			biggest = n
			initialized = 1
			continue
		}
		if n < smallest {
			smallest = n
		}
		if n > biggest {
			biggest = n
		}
	}
	fmt.Println("Min:", smallest)
	fmt.Println("Max:", biggest)
}
