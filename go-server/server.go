package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Response struct {
	Result int `json:"result"`
}

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

func factorialRecursive(inputNumber int) int {
	if inputNumber < 2 {
		return 1
	}

	return inputNumber * factorialRecursive(inputNumber-1)
}

func handler(response http.ResponseWriter, request *http.Request) {
	// set what kind of data is being returned to the client
	response.Header().Set("Content-Type", "application/json")

	// populate the form with the request parameters / body
	request.ParseForm()

	inputNumberAsString := request.Form["inputNumber"][0]
	// fmt.Println(inputNumberAsString)

	inputNumberAsInt64, error := strconv.ParseInt(inputNumberAsString, 36, 12)
	inputNumberAsInt := int(inputNumberAsInt64)

	if error != nil {
		fmt.Println("You done messed up")
	}

	fmt.Printf("Input Type: %T      Input Value: %v \n", inputNumberAsInt, inputNumberAsInt)

	factorialResult := factorial(inputNumberAsInt)
	fmt.Printf("Return Type: %T     Return Value: %v \n\n", factorialResult, factorialResult)

	result := Response{
		Result: factorialResult,
	}

	json.NewEncoder(response).Encode(result)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Printf("Server listening on localhost:3124\n\n")
	log.Fatal(http.ListenAndServe(":3124", nil))
}
