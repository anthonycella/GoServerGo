package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func isValidInput(userInput string) bool {
	inputNumberAsInt64, error := strconv.ParseInt(userInput, 0, 0)

	if error != nil {
		return false
	}

	if inputNumberAsInt64 < 0 || inputNumberAsInt64 > 20 {
		return false
	}

	return true
}

func printFactorialResult(userInput string) {
	response, error := http.PostForm("http://localhost:3124/", url.Values{"inputNumber": {userInput}})

	if error != nil {
		fmt.Print("Error, unable to retrieve answer from server\n\n")
	} else {
		var responseJson map[string]interface{}
		errorInDecoding := json.NewDecoder(response.Body).Decode(&responseJson)

		if errorInDecoding != nil {
			fmt.Println(errorInDecoding)
		} else {
			factorialOfInput := responseJson["result"]
			fmt.Print("The factorial of ", userInput, " is ", factorialOfInput, "\n\n")
		}
	}
}

func main() {
	fmt.Println("\nEnter a positive number between 0 and 20 inclusive to compute factorial or x to exit")

	var userInput string
	fmt.Scanln(&userInput)

	for userInput != "x" {
		if !isValidInput(userInput) {
			fmt.Print("Error: please input either a positive number between 0 and 20 inclusive or x\n\n")
		} else {
			printFactorialResult(userInput)
		}

		fmt.Println("Enter a positive number between 0 and 20 inclusive to compute factorial or x to exit")
		fmt.Scanln(&userInput)
	}

	fmt.Println("Goodbye")
}
