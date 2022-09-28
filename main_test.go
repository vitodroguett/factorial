package main

import "testing"

func TestSuccessFactorial(t *testing.T) {
	//arrange
	var n int = 6
	var factorialExpected int = 720

	//act
	factorial, _ := factorial(n)

	//assert
	if factorial != factorialExpected {
		t.Errorf("Factorial expected value %d got %d ", factorialExpected, factorial)
	}
}

func TestFailFactorial(t *testing.T) {
	//arrange
	n := -1

	//act
	factorial, err := factorial(n)

	//assert
	if err == nil {
		t.Errorf("Expected an error, but got a factorial result: %d ", factorial)
	}
}
