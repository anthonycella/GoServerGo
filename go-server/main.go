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

func handler(w http.ResponseWriter, r *http.Request) {
	// set what kind of data is being returned to the client
	w.Header().Set("Content-Type", "application/json")

	inputNumberAsString := r.URL.Path[1:]
	// fmt.Println(inputNumberAsString)

	inputNumberAsInt64, error := strconv.ParseInt(inputNumberAsString, 36, 12)
	inputNumberAsInt := int(inputNumberAsInt64)

	if error != nil {
		fmt.Println("You done messed up")
	}

	fmt.Printf("Input Type: %T      Input Value: %v \n", inputNumberAsInt, inputNumberAsInt)

	factorialResult := factorial(inputNumberAsInt)
	fmt.Printf("Return Type: %T     Return Value: %v \n", factorialResult, factorialResult)

	result := Response{
		Result: factorialResult,
	}

	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":3124", nil))
}
