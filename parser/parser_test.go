package parser

import (
	"testing"

	"json-parser/types"
)

func TestParseJSONBoolean(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONBool
		err      bool
	}{
		{"true", true, false},
		{"false", false, false},
		{"true,", true, false},
		{"false,", false, false},
		{"true}", true, false},
		{"false}", false, false},
	}

	for _, test := range tests {
		result, _, err := parseBool(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseBool(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && result != test.expected {
			t.Errorf("parseBool(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestParseJSONNumber(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONNumber
		err      bool
	}{
		{"0", 0, false},
		{"123", 123, false},
		{"-456", -456, false},
		{"3.14", 3.14, false},
		{"1.23e4", 12300, false},
		{"invalid", 0, true},
	}

	for _, test := range tests {
		result, _, err := parseNumber(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseNumber(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && result != test.expected {
			t.Errorf("parseNumber(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}

func TestParseString(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONString
		err      bool
	}{
		{`""`, "", false},
		{`"hello"`, "hello", false},
		{`"hello\nworld"`, "hello\nworld", false},
		{`"\""`, "\"", false},
		{`"\u0020"`, " ", false},
		{`"\u0041"`, "A", false},
		{`"invalid`, "", true}, 
	}

	for _, test := range tests {
		result, _, err := parseString(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseString(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && result != test.expected {
			t.Errorf("parseString(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}

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
	if obj["name"] != types.JSONString("John") {
		t.Errorf("Expected name to be John, got %v", obj["name"])
	}
	if obj["age"] != types.JSONNumber(30) {
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
	if arr[0] != types.JSONNumber(1) {
		t.Errorf("Expected first element to be 1, got %v", arr[0])
	}
	if arr[1] != types.JSONNumber(2) {
		t.Errorf("Expected second element to be 2, got %v", arr[1])
	}
	if arr[2] != types.JSONNumber(3) {
		t.Errorf("Expected third element to be 3, got %v", arr[2])
	}
}
