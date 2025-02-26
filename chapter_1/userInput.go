package main

import "fmt"

func main() {
	// Get user input
	fmt.Print("Please give me your name: ")
	var name string
	// &name -> Pointer to name var
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("Aborting due to error", err.Error())
		return
	}
	fmt.Println("Your name is", name)
}
