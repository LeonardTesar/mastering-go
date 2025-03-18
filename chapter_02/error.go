package chapter_02

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func customNewError(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("this is a custom error message")
	}
	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d. UserID: %d", a, b, os.Geteuid())
	}
	return nil
}

func CustomErrors() {
	err := customNewError(0, 10)
	if err == nil {
		fmt.Println("customNewError() executed normally")
	} else {
		fmt.Println(err)
	}
	err = customNewError(0, 0)
	// Bad practice. Print something if err != nil.
	if err.Error() == "this is a custom error message" {
		fmt.Println("Custom error detected")
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	}
}
