package types

type JSONValue interface {}

type JSONObject map[string]JSONValue
type JSONArray []JSONValue
type JSONString string
type JSONNumber float64
type JSONBool bool
type JSONNull struct {}