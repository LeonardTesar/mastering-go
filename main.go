package main

import (
	"Mastering_Go/chapter_02"
	"Mastering_Go/chapter_03"
	"Mastering_Go/chapter_1"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	for {

		fmt.Println("Please enter the function you want to execute or 'exit' to quit.")
		fmt.Println("Schema: functionCall argument1 argument2 ... argumentN")
		fmt.Print("Call: ")
		var functionCall string

		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			fmt.Println("Aborting due to error:", scanner.Err().Error())
			return
		}

		functionCall = scanner.Text()
		functionCallParts := strings.Split(functionCall, " ")
		functionName := functionCallParts[0]
		os.Args = os.Args[:1]
		for i := 1; i < len(functionCallParts); i++ {
			os.Args = append(os.Args, functionCallParts[i])
		}

		switch functionName {
		case "which":
			chapter_1.Which()
		case "variables":
			chapter_1.Variables()
		case "userInput":
			chapter_1.UserInput()
		case "process":
			chapter_1.Process()
		case "hw":
			chapter_1.Hw()
		case "goRoutines":
			chapter_1.GoRoutines()
		case "forLoops":
			chapter_1.ForLoops()
		case "control":
			chapter_1.Control()
		case "cla":
			chapter_1.Cla()
		case "systemLog":
			chapter_1.SystemLog()
		case "multiLog":
			chapter_1.MultiLog()
		case "stats":
			chapter_1.Stats()
		case "customErrors":
			chapter_02.CustomErrors()
		case "numbers":
			chapter_02.Numbers()
		case "stats2":
			chapter_02.Stats()
		case "csvReadWrite":
			chapter_03.ReadWriteCsv()
		case "exit":
			os.Exit(0)
		default:
			fmt.Printf("Function %s is not yet implemented.\n", functionName)
		}
	}
}
