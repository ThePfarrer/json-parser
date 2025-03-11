package parser

import (
	"testing"

	"json-parser/types"
)

func TestParseJSONObject(t *testing.T) {
	json := `{"name": "John", "age": 30}`
	parsed, err := ParseJSON(json)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	obj, ok := parsed.(types.JSONObject)
	if !ok {
		t.Fatalf("Expected JSONValue, got %T", parsed)
	}
	if obj["name"] != "John" {
		t.Errorf("Expected name to be John, got %v", obj["name"])
	}
	if obj["age"] != float64(30) {
		t.Errorf("Expected age to be 30, got %v", obj["age"])
	}
}

func TestParseJSONArray(t *testing.T) {
	json := `[1, 2, 3]`
	parsed, err := ParseJSON(json)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	arr, ok := parsed.(types.JSONArray)
	if !ok {
		t.Fatalf("Expected JSONArray, got %T", parsed)
	}
	if len(arr) != 3 {
		t.Errorf("Expected 3 elements, got %v", len(arr))
	}
	if arr[0] != float64(1) {
		t.Errorf("Expected first element to be 1, got %v", arr[0])
	}
	if arr[1] != float64(2) {
		t.Errorf("Expected second element to be 2, got %v", arr[1])
	}
	if arr[2] != float64(3) {
		t.Errorf("Expected third element to be 3, got %v", arr[2])
	}
}
