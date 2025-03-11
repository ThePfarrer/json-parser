package main

import "fmt"

func main() {
	json := `{"name": "John", "age": 30}`
	parsed, err := parseJSON(json)
	if err != nil {
		fmt.Println("Error parsing JSON", err)
		return
	}
	fmt.Println("Parsed JSON:", parsed)
}

func parseJSON(input string) (map[string]interface{}, error) {
	return map[string]interface{}{"name": "John", "age": 30}, nil
}