package main

import (
	"fmt"
)

func factorial(inputNumber int) int {
	if inputNumber < 2 {
		return 1
	}

	factorialResult := 1

	for inputNumber > 0 {
		factorialResult *= inputNumber
		inputNumber -= 1
	}

	return factorialResult
}

func main() {
	// fmt.Println("I am Batman")
	// fmt.Println("I am going to kill you", time.Now())
	// fmt.Println("My favorite number is", rand.Intn(10))
	// fmt.Printf("Now you have %g problems. \n", math.Sqrt(4))
	// fmt.Println(math.Pi)
	fmt.Println(factorial(5))
}
