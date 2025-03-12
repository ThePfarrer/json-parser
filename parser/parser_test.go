package parser

import (
	"reflect"
	"testing"

	"json-parser/types"
)

func TestParseJSONBool(t *testing.T) {
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

func TestParseJSON(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONValue
		err      bool
	}{
		{`{"name": "John", "age": 30}`, types.JSONObject{"name": types.JSONString("John"), "age": types.JSONNumber(30)}, false},
		{`[1, 2, 3]`, types.JSONArray{types.JSONNumber(1), types.JSONNumber(2), types.JSONNumber(3)}, false},
		{`"hello"`, types.JSONString("hello"), false},
		{`123`, types.JSONNumber(123), false},
		{`true`, types.JSONBool(true), false},
		{`false`, types.JSONBool(false), false},
		{`null`, types.JSONNull{}, false},
		{``, nil, true},
		{`invalid`, nil, true},
		{`{"name": "John", "age": 30`, nil, true},
		{`[1, 2, 3`, nil, true},
	}

	for _, test := range tests {
		result, err := ParseJSON(test.input)
		if (err != nil) != test.err {
			t.Errorf("ParseJSON(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !reflect.DeepEqual(result, test.expected) {
			t.Errorf("ParseJSON(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
func TestIsNumberChar(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{'0', true},
		{'9', true},
		{'.', true},
		{'-', true},
		{'e', true},
		{'E', true},
		{'+', true},
		{'a', false},
		{'Z', false},
		{' ', false},
		{'$', false},
	}

	for _, test := range tests {
		result := isNumberChar(test.input)
		if result != test.expected {
			t.Errorf("isNumberChar(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
func TestParseNull(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONNull
		err      bool
	}{
		{"null", types.JSONNull{}, false},
		{"null,", types.JSONNull{}, false},
		{"null}", types.JSONNull{}, false},
		{"nul", types.JSONNull{}, true},
		{"invalid", types.JSONNull{}, true},
	}

	for _, test := range tests {
		result, _, err := parseNull(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseNull(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && result != test.expected {
			t.Errorf("parseNull(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
func TestUnescapeString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`hello`, "hello"},
		{`hello\nworld`, "hello\nworld"},
		{`hello\\world`, "hello\\world"},
		{`hello\"world`, "hello\"world"},
		{`hello\/world`, "hello/world"},
		{`hello\bworld`, "hello\bworld"},
		{`hello\fworld`, "hello\fworld"},
		{`hello\rworld`, "hello\rworld"},
		{`hello\tworld`, "hello\tworld"},
		{`hello\u0020world`, "hello world"},
		{`hello\u0041world`, "helloAworld"},
		{`hello\u`, "hello\\u"},
		{`hello\u123`, "hello\\u123"},
	}

	for _, test := range tests {
		result := unescapeString(test.input)
		if result != test.expected {
			t.Errorf("unescapeString(%q) = %q, want %q", test.input, result, test.expected)
		}
	}
}
func TestParseArray(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONArray
		err      bool
	}{
		{`[]`, types.JSONArray{}, false},
		{`[1, 2, 3]`, types.JSONArray{types.JSONNumber(1), types.JSONNumber(2), types.JSONNumber(3)}, false},
		{`["a", "b", "c"]`, types.JSONArray{types.JSONString("a"), types.JSONString("b"), types.JSONString("c")}, false},
		{`[true, false, null]`, types.JSONArray{types.JSONBool(true), types.JSONBool(false), types.JSONNull{}}, false},
		{`[1, "a", true, null]`, types.JSONArray{types.JSONNumber(1), types.JSONString("a"), types.JSONBool(true), types.JSONNull{}}, false},
		{`[1, 2, 3`, nil, true},
		{`[1, 2, 3,]`, nil, true},
		{`[1, "a", true, null,]`, nil, true},
		{`[`, nil, true},
		{`invalid`, nil, true},
	}

	for _, test := range tests {
		result, _, err := parseArray(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseArray(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseArray(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}
func TestParseObject(t *testing.T) {
	tests := []struct {
		input    string
		expected types.JSONObject
		err      bool
	}{
		{`{}`, types.JSONObject{}, false},
		{`{"name": "John"}`, types.JSONObject{"name": types.JSONString("John")}, false},
		{`{"age": 30}`, types.JSONObject{"age": types.JSONNumber(30)}, false},
		{`{"name": "John", "age": 30}`, types.JSONObject{"name": types.JSONString("John"), "age": types.JSONNumber(30)}, false},
		{`{"name": "John", "age": 30, "isStudent": true}`, types.JSONObject{"name": types.JSONString("John"), "age": types.JSONNumber(30), "isStudent": types.JSONBool(true)}, false},
		{`{"name": "John", "address": {"city": "New York", "zip": "10001"}}`, types.JSONObject{"name": types.JSONString("John"), "address": types.JSONObject{"city": types.JSONString("New York"), "zip": types.JSONString("10001")}}, false},
		{`{"name": "John", "hobbies": ["reading", "swimming"]}`, types.JSONObject{"name": types.JSONString("John"), "hobbies": types.JSONArray{types.JSONString("reading"), types.JSONString("swimming")}}, false},
		{`{"name": "John", "age": 30`, nil, true},
		{`{"name": "John", "age": 30,}`, nil, true},
		{`{"name": "John", "age": 30, "isStudent": true,}`, nil, true},
		{`{`, nil, true},
		{`invalid`, nil, true},
	}

	for _, test := range tests {
		result, _, err := parseObject(test.input)
		if (err != nil) != test.err {
			t.Errorf("parseObject(%q) error: %v, expected error: %v", test.input, err, test.err)
		}
		if !test.err && !reflect.DeepEqual(result, test.expected) {
			t.Errorf("parseObject(%q) = %v, want %v", test.input, result, test.expected)
		}
	}
}





