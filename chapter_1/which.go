package chapter_1

import (
	"fmt"
	"os"
	"path/filepath"
)

func Which() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}

	file := arguments[1]
	path := os.Getenv("PATH")
	pathSplit := filepath.SplitList(path)

	for _, directory := range pathSplit {
		fullPath := filepath.Join(directory, file)
		fileInfo, err := os.Stat(fullPath)
		if err != nil {
			continue
		}

		mode := fileInfo.Mode()

		if !mode.IsRegular() {
			continue
		}
		// commented because the comparison always produced ---------- although permissions were there (and I don't
		// want to put effort into fixing simple learning code
		//if mode&0111 != 0 {
		fmt.Println(fullPath)
		//	return
		//}
	}
}
