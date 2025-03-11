package parser

import (
	"testing"
)

func TestParseJSON(t *testing.T) {
	json := `{"name": "John", "age": 30}`
	parsed, err := ParseJSON(json)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if parsed["name"] != "John" {
		t.Errorf("Expected name to be John, got %v", parsed["name"])
	}
	if parsed["age"] != float64(30) {
		t.Errorf("Expected age to be 30, got %v", parsed["age"])
	}
}
