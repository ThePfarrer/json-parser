package main

import (
	"fmt"
	"os"

	"json-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: json-parser <file>")
		return
	}

	filePath := os.Args[1]
	json, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	parsed, err := parser.ParseJSON(string(json))
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}

	fmt.Println("Parsed JSON object:", parsed)
}
