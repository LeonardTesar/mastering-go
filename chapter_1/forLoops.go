package main

import "fmt"

func main() {
	// Traditional for loop
	for i := 0; i < 10; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()

	// Same but different style
	i := 0
	for ok := true; ok; ok = i != 10 {
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// while loop
	i = 0
	for {
		if i == 10 {
			break
		}
		fmt.Print(i*i, " ")
		i++
	}
	fmt.Println()

	// forEach loop. If the index is not needed, it can be omitted by replacing i with an underscore (_)
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value:", v)
	}
}
