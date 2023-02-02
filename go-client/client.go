package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func userWantsToExit(userInput string) bool {
	return userInput == "x"
}

func isValidInput(userInput string) bool {
	inputNumberAsInt64, error := strconv.ParseInt(userInput, 0, 0)

	userDidntInputANumber := error != nil
	userInputANumberOutOfRange := inputNumberAsInt64 < 0 || inputNumberAsInt64 > 20

	if userDidntInputANumber || userInputANumberOutOfRange {
		return false
	} else {
		return true
	}
}

func printFactorialResult(userInput string) {
	factorialResponseFromServer, error := http.PostForm("http://localhost:3124/", url.Values{"inputNumber": {userInput}})

	unableToRetrieveFromServer := error != nil

	if unableToRetrieveFromServer {
		fmt.Print("Error, unable to retrieve answer from server\n\n")
	} else {
		var factorialResponseJson map[string]interface{}
		errorInDecoding := json.NewDecoder(factorialResponseFromServer.Body).Decode(&factorialResponseJson)

		unableToDecodeJson := errorInDecoding != nil

		if unableToDecodeJson {
			fmt.Println("Error, unable to decode data retrieved from server (expecting JSON)")
		} else {
			factorialOfInput := factorialResponseJson["result"]
			fmt.Print("The factorial of ", userInput, " is ", factorialOfInput, "\n\n")
		}
	}
}

func main() {
	fmt.Println("\nEnter a positive number between 0 and 20 inclusive to compute factorial or x to exit")

	var userInput string
	fmt.Scanln(&userInput)

	for !userWantsToExit(userInput) {
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
