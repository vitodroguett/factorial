package main

import "testing"

func TestSuccessFactorial(t *testing.T) {
	//arrange
	var n int = 6
	var factorialExpected int = 720

	//act
	factorial := factorial(n)

	//assert
	if factorial != factorialExpected {
		t.Errorf("Factorial expected value %d got %d ", factorialExpected, factorial)
	}
}

func TestFailFactorial(t *testing.T) {
	//arrange
	n := 6
	var factorialExpected int = 720

	//act
	factorial := factorial(n)

	//assert
	if factorial != factorialExpected {
		t.Errorf("Factorial expected value %d got %d ", factorialExpected, factorial)
	}
}
