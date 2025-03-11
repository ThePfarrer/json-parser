package parser

import (
	"errors"
	"strings"

	"json-parser/types"
)

func ParseJSON(input string) (types.JSONValue, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil, errors.New("empty input")
	}

	value, rest, err := parseValue(input)
	if err != nil {
		return nil, err
	}

	rest = strings.TrimSpace(rest)
	if len(rest) != 0 {
		return nil, errors.New("unexpected trailing characters")
	}
	return value, nil
}

func parseValue(input string) (types.JSONValue, string, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil, input, errors.New("unexpected end of input")
	}

	switch input[0] {
	case '{':
		return parseObject(input)
	case '[':
		return parseArray(input)
	case '"':
		return parseString(input)
	case 't', 'f':
		return parseBool(input)
	case 'n':
		return parseNull(input)
	default:
		return parseNumber(input)
	}
}

func parseNumber(input string) (types.JSONValue, string, error) {
	panic("unimplemented")
}

func parseNull(input string) (types.JSONValue, string, error) {
	if strings.HasPrefix(input, "null") {
		return types.JSONNull{}, input[4:], nil
	}
	return nil, input, errors.New("invalid null")
}

func parseBool(input string) (types.JSONValue, string, error) {
	if strings.HasPrefix(input, "true") {
		return types.JSONBool(true), input[4:], nil
	} else if strings.HasPrefix(input, "false") {
		return types.JSONBool(false), input[5:], nil
	}
	return nil, input, errors.New("invalid boolean")
}

func parseString(input string) (types.JSONString, string, error) {
	if input[0] != '"' {
		return "", input, errors.New("invalid string")
	}
	end := 1
	for end < len(input) && input[end] != '"' {
		if input[end] == '\\' {
			end++
		}
		end++
	}
	if end >= len(input) {
		return "", input, errors.New("unterminated string")
	}
	raw := input[1:end]
	unescaped := raw
	return types.JSONString(unescaped), input[end+1:], nil
}

func parseArray(input string) (types.JSONArray, string, error) {
	if input[0] != '[' {
		return nil, input, errors.New("invalid array")
	}
	arr := types.JSONArray{}
	input = input[1:]
	for {
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == ']' {
			return arr, input[1:], nil
		}
		value, rest, err := parseValue(input)
		if err != nil {
			return nil, input, err
		}
		arr = append(arr, value)
		input = strings.TrimSpace(rest)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == ']' {
			return arr, input[1:], nil
		}
		if input[0] != ',' {
			return nil, input, errors.New("expected ',' or ']'")
		}
		input = input[1:]
	}
}

func parseObject(input string) (types.JSONObject, string, error) {
	if input[0] != '{' {
		return nil, input, errors.New("invalid object")
	}
	obj := types.JSONObject{}
	input = input[1:]
	for {
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == '}' {
			return obj, input[1:], nil
		}
		key, rest, err := parseString(input)
		if err != nil {
			return nil, input, err
		}
		input = strings.TrimSpace(rest)
		if len(input) == 0 || input[0] != ':' {
			return nil, input, errors.New("expected ':' after key")
		}
		input = strings.TrimSpace(input[1:])
		value, rest, err := parseValue(input)
		if err != nil {
			return nil, input, err
		}

		obj[string(key)] = value
		input = strings.TrimSpace(rest)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == '}' {
			return obj, input[1:], nil
		}
		if input[0] != ',' {
			return nil, input, errors.New("expected ',' or '}'")
		}
		input = input[1:]
	}
}
