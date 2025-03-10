package chapter_1

import (
	"fmt"
	"os"
	"strconv"
)

func Process() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Not enough arguments")
		return
	}
	var total, nInts, nFloats int
	invalid := make([]string, 0)
	for _, k := range arguments[1:] {

		// Is it an int?
		_, err := strconv.Atoi(k)
		if err == nil {
			total++
			nInts++
			continue
		}

		// or is it a float?
		_, err = strconv.ParseFloat(k, 64)
		if err == nil {
			total++
			nFloats++
			continue
		}

		// No, it's invalid!
		invalid = append(invalid, k)
	}

	fmt.Println("#read:", total, "#ints:", nInts, "#floats:", nFloats)
	if len(invalid) > total {
		fmt.Println("Too much invalid input:", len(invalid))
		for _, s := range invalid {
			fmt.Println(s)
		}
	}
}
