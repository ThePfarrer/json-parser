package parser

import (
	"encoding/json"
	"errors"

	"json-parser/types"
)

func ParseJSON(input string) (types.JSONValue, error) {
	var result interface{}
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return convertToJSONValue(result)
}

func convertToJSONValue(result interface{}) (types.JSONValue, error) {
	switch v := result.(type) {
	case map[string]interface{}:
		obj := types.JSONObject{}
		for key, value := range v {
			convertedValue, err := convertToJSONValue(value)
			if err != nil {
				return nil, err
			}
			obj[key] = convertedValue
		}
		return obj, nil
	case []interface{}:
		arr := types.JSONArray{}
		for _, value := range v {
			convertedValue, err := convertToJSONValue(value)
			if err != nil {
				return nil, err
			}
			arr = append(arr, convertedValue)
		}
		return arr, nil
	case string:
		return types.JSONString(v), nil
	case float64:
		return types.JSONNumber(v), nil
	case bool:
		return types.JSONBool(v), nil
	case nil:
		return types.JSONNull{}, nil
	default:
		return nil, errors.New("unknown JSON value type")
	}
}

// func parseValue(input string) (types.JSONValue, error) {
// 	var result types.JSONValue
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func parseObject(input string) (types.JSONObject, error) {
// 	var result types.JSONObject
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func parseArray(input string) (types.JSONArray, error) {
// 	var result types.JSONArray
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

// func parserString(input string) (types.JSONString, error) {
// 	var result types.JSONString
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return "", err
// 	}
// 	return result, nil
// }

// func parseNumber(input string) (types.JSONNumber, error) {
// 	var result types.JSONNumber
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return 0, err
// 	}
// 	return result, nil
// }

// func parseBool(input string) (types.JSONBool, error) {
// 	var result types.JSONBool
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return false, err
// 	}
// 	return result, nil
// }

// func parseNull(input string) (types.JSONNull, error) {
// 	var result types.JSONNull
// 	err := json.Unmarshal([]byte(input), &result)
// 	if err != nil {
// 		return types.JSONNull{}, err
// 	}
// 	return result, nil
// }
