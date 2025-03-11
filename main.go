package main

import (
	"fmt"

	"json-parser/parser"
	"json-parser/types"
)

func main() {
	json := `{"name": "John", "age": 30}`
	parsed, err := parser.ParseJSON(json)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}

	switch v := parsed.(type) {
	case types.JSONObject:
		fmt.Println("Parsed JSON object:", v)
	case types.JSONArray:
		fmt.Println("Parsed JSON array:", v)
	case types.JSONString:
		fmt.Println("Parsed JSON string:", v)
	case types.JSONNumber:
		fmt.Println("Parsed JSON number:", v)
	case types.JSONBool:
		fmt.Println("Parsed JSON boolean:", v)
	case types.JSONNull:
		fmt.Println("Parsed JSON null:", v)
	default:
		fmt.Println("Parsed JSONv value:", v)
	}
}
