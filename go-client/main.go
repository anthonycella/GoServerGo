package main

import (
	"fmt"
	"strconv"
)

func inputAsInteger(userInput string) int {
	inputNumberAsInt64, error := strconv.ParseInt(userInput, 36, 12)

	if error != nil {
		return -1
	}

	inputNumberAsInt := int(inputNumberAsInt64)

	if inputNumberAsInt < 0 {
		return -1
	}

	return inputNumberAsInt
}

func factorial(inputNumber int) int {
	factorialOfInputNumber := 1

	for inputNumber > 0 {
		factorialOfInputNumber *= inputNumber
		inputNumber--
	}

	return factorialOfInputNumber
}

func main() {
	fmt.Println("Enter a positive number to compute factorial or x to exit")

	var userInput string
	fmt.Scanln(&userInput)

	for userInput != "x" {
		numberFormattedInput := inputAsInteger(userInput)

		if numberFormattedInput == -1 {
			fmt.Println("Error: please input either a positive number or x")
		} else {
			factorialOfInput := factorial(numberFormattedInput)
			fmt.Println("The factorial of", userInput, "is", factorialOfInput)
		}

		fmt.Println("Enter a positive number to compute factorial or x to exit")
		fmt.Scanln(&userInput)
	}

	fmt.Println("Goodbye")
}
