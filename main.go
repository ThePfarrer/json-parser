package main

import (
	"fmt"

	"json-parser/parser"
)

func main() {
	json := `{"name": "John", "age": 30}`
	parsed, err := parser.ParseJSON(json)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}
	fmt.Println("Parsed JSON:", parsed)
}
