package main

import (
	"fmt"
	"os"
	"path/filepath"

	"json-parser/parser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: json-parser <file>")
		return
	}

	path := os.Args[1]
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error accessing path:", err)
		return
	}
	if fileInfo.IsDir() {
		err = processDirectory(path)
	} else {
		err = processFile(path)
	}

	if err != nil {
		fmt.Println("Error:", err)
	}
}

func processFile(path string) error {
	json, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	parsed, err := parser.ParseJSON(string(json))
	if err != nil {
		return fmt.Errorf("error parsing JSON: %v", err)
	}

	fmt.Println("Parsed JSON object from", path, ":")
	fmt.Println(parsed)
	return nil
}

func processDirectory(path string) error {
	files, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			err := processFile(filepath.Join(path, file.Name()))
			if err != nil {
				fmt.Println("Error processing file", file.Name(), err)
			}
		}
	}
	return nil
}
