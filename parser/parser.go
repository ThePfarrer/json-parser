package parser

import (
	"encoding/json"

	"json-parser/types"
)

func ParseJSON(input string) (types.JSONValue, error) {
	var result types.JSONValue
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseValue(input string) (types.JSONValue, error) {
	var result types.JSONValue
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseObject(input string) (types.JSONObject, error) {
	var result types.JSONObject
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parseArray(input string) (types.JSONArray, error) {
	var result types.JSONArray
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func parserString(input string) (types.JSONString, error) {
	var result types.JSONString
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return "", err
	}
	return result, nil
}

func parseNumber(input string) (types.JSONNumber, error) {
	var result types.JSONNumber
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func parseBool(input string) (types.JSONBool, error) {
	var result types.JSONBool
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return false, err
	}
	return result, nil
}

func parseNull(input string) (types.JSONNull, error) {
	var result types.JSONNull
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return types.JSONNull{}, err
	}
	return result, nil
}
