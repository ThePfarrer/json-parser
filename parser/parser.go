package parser

import (
	"errors"
	"strconv"
	"strings"
	"unicode"

	"json-parser/types"
)

func ParseJSON(input string) (types.JSONValue, error) {
	input = strings.TrimSpace(input)
	if len(input) == 0 {
		return nil, errors.New("empty input")
	}

	if input[0] != '[' && input[0] != '{' {
		return nil, errors.New("JSON must start with '[' or '{'")
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

func parseNumber(input string) (types.JSONNumber, string, error) {
	if len(input) > 1 && input[0] == '0' && unicode.IsDigit(rune(input[1])) {
		return 0, input, errors.New("numbers cannot have leading zeroes")
	}
	end := 0
	for end < len(input) && isNumberChar(input[end]) {
		end++
	}
	if end == 0 {
		return 0, input, errors.New("invalid number")
	}
	numStr := input[:end]
	num, err := strconv.ParseFloat(numStr, 64)
	if err != nil {
		return 0, input, err
	}
	return types.JSONNumber(num), input[end:], nil
}

func isNumberChar(b byte) bool {
	return unicode.IsDigit(rune(b)) || b == '.' || b == '-' || b == 'e' || b == 'E' || b == '+'
}

func parseNull(input string) (types.JSONNull, string, error) {
	if strings.HasPrefix(input, "null") {
		return types.JSONNull{}, input[4:], nil
	}
	return types.JSONNull{}, input, errors.New("invalid null")
}

func parseBool(input string) (types.JSONBool, string, error) {
	if strings.HasPrefix(input, "true") {
		return types.JSONBool(true), input[4:], nil
	} else if strings.HasPrefix(input, "false") {
		return types.JSONBool(false), input[5:], nil
	}
	return types.JSONBool(false), input, errors.New("invalid boolean")
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
	unescaped, err := unescapeString(raw)
	if err != nil {
		return "", input, err
	}
	return types.JSONString(unescaped), input[end+1:], nil
}

func unescapeString(raw string) (string, error) {
	var result strings.Builder
	for i := 0; i < len(raw); i++ {
		if raw[i] == '\\' {
			i++
			if i < len(raw) {
				switch raw[i] {
				case '"':
					result.WriteByte('"')
				case '\\':
					result.WriteByte('\\')
				case '/':
					result.WriteByte('/')
				case 'b':
					result.WriteByte('\b')
				case 'f':
					result.WriteByte('\f')
				case 'n':
					result.WriteByte('\n')
				case 'r':
					result.WriteByte('\r')
				case 't':
					result.WriteByte('\t')
				case 'u':
					if i+4 < len(raw) {
						hex := raw[i+1 : i+5]
						codePoint, err := strconv.ParseUint(hex, 16, 32)
						if err == nil {
							result.WriteRune(rune(codePoint))
							i += 4
						} else {
							return "", errors.New("invalid escape sequence")
						}
					} else {
						return "", errors.New("invalid escape sequence")
					}
				default:
					return "", errors.New("invalid escape sequence")
				}
			}
		} else {
			result.WriteByte(raw[i])
		}
	}
	return result.String(), nil
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
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == ']' {
			return nil, input, errors.New("trailing comma in array")
		}
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
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			return nil, input, errors.New("unexpected end of input")
		}
		if input[0] == '}' {
			return nil, input, errors.New("trailing comma in array")
		}
	}
}
