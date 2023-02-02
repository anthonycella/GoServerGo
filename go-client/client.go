package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type RequestParameters struct {
	InputNumber int `json:"inputNumber"`
}

func isValidInput(userInput string) bool {
	inputNumberAsInt64, error := strconv.ParseInt(userInput, 36, 12)

	if error != nil {
		return false
	}

	if inputNumberAsInt64 < 0 {
		return false
	}

	return true
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
		if !isValidInput(userInput) {
			fmt.Println("Error: please input either a positive number or x")
		} else {
			response, error := http.PostForm("http://localhost:3124/", url.Values{"inputNumber": {userInput}})

			if error != nil {
				fmt.Println("Error, unable to retrieve answer from server")
			} else {
				var responseJson map[string]interface{}
				errorDecoding := json.NewDecoder(response.Body).Decode(&responseJson)

				if errorDecoding != nil {
					fmt.Println(errorDecoding)
				} else {
					factorialOfInput := responseJson["result"]
					fmt.Println("The factorial of", userInput, "is", factorialOfInput)
				}

			}
		}

		fmt.Println("Enter a positive number to compute factorial or x to exit")
		fmt.Scanln(&userInput)
	}

	fmt.Println("Goodbye")
}
