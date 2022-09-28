package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("We work")
}

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("cannot be less than 0")
	}

	if n == 0 {
		return 1, nil
	} else {
		factorial, _ := factorial(n - 1)
		return (n * factorial), nil
	}
}
