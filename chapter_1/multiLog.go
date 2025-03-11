package chapter_1

import (
	"fmt"
	"io"
	"log"
	"os"
)

func MultiLog() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("app.log", flag, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	defer file.Close()

	// This function is variadic - meaning it can take as many writers as you want. os.Stdout or os.Stderr would be the
	// basic ones (usually console). Keep in mind that container software usually handles the stdout to persist logs!
	w := io.MultiWriter(file, os.Stdout)
	logger := log.New(w, "MultiLog: ", log.LstdFlags)
	logger.Printf("Book %d", os.Getpid())
}
