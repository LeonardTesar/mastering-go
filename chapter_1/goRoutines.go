package main

import (
	"fmt"
	"time"
)

func myPrint(start, finish int) {
	for i := start; i <= finish; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()
	time.Sleep(100 * time.Microsecond)
}

// Incorrect (unsafe) way of synchronizing. Correct one will be covered in chapter_8
// This program will most likely produce a different output each time it is executed, since the execution (and order) of
// go routines is handled by th go scheduler and are initialized (and started) in a random order.
func main() {
	for i := 0; i < 4; i++ {
		go myPrint(i, 5)
	}
	time.Sleep(time.Second)
}
