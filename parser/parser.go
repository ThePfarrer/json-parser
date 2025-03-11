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
	return parseValue(input)
}

func parseValue(input string) (types.JSONValue, error) {
	switch {
	case input[0] == '{':
		return parseObject(input)
	case input[0] == '[':
		return parseArray(input)
	case input[0] == '"':
		return parserString(input)
	case input[0] == 't' || input[0] == 'f':
		return parseBool(input)
	case input[0] == 'n':
		return parseNull(input)
	default:
		return parseNumber(input)
	}
}

func parseNumber(input string) (types.JSONValue, error) {
	panic("unimplemented")
}

func parseNull(input string) (types.JSONValue, error) {
	panic("unimplemented")
}

func parseBool(input string) (types.JSONValue, error) {
	panic("unimplemented")
}

func parserString(input string) (types.JSONValue, error) {
	panic("unimplemented")
}

func parseArray(input string) (types.JSONValue, error) {
	panic("unimplemented")
}

func parseObject(input string) (types.JSONValue, error) {
	panic("unimplemented")
}
