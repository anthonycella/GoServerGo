package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type RequestParameters struct {
	InputNumber int `json:"inputNumber"`
}

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
			response, error := http.Get("http://localhost:3124/?inputNumber=7")

			if error != nil {
				fmt.Println("Error, unable to retrieve answer from server")
			} else {
				var responseJson map[string]interface{}
				var factorialOfInput int
				errorDecoding := json.NewDecoder(response.Body).Decode(&responseJson)

				if errorDecoding != nil {
					fmt.Println(errorDecoding)
				} else {
					fmt.Printf("Type of response: %T   Response Value: %v", responseJson, responseJson)
					fmt.Println("The factorial of", userInput, "is", factorialOfInput)
				}

			}
		}

		fmt.Println("Enter a positive number to compute factorial or x to exit")
		fmt.Scanln(&userInput)
	}

	fmt.Println("Goodbye")
}
