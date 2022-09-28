package main

import "fmt"

func main() {
	n := 6
	factorial := factorial(n)
	fmt.Printf("factorial de %d es => %d", n, factorial)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	} else {
		return n * factorial(n-1)
	}
}
