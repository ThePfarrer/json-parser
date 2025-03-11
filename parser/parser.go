package parser

import "encoding/json"

func ParseJSON(input string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err:= json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}