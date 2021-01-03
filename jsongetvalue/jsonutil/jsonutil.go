package jsonutil

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// Extract value by specified key from json object typed as interface{}
// Example: object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb"
func Get(object interface{}, keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	var result interface{}
	var exists bool
	for _, key := range keys {
		exists = false
		value, ok := object.(map[string]interface{})
		if ok {
			exists = true
			object = value[key]
			result = object
			continue
		}
		values, ok := object.([]interface{})
		if ok {
			for i, v := range values {
				if strconv.FormatInt(int64(i), 10) == key {
					exists = true
					object = v
					result = object
					continue
				}
			}
		}
	}
	if exists {
		return result
	}
	return nil
}

// Extract and cast value as string
func AsString(object interface{}, keyChain string) *string {
	maybeString := Get(object, keyChain)
	if stringValue, ok := maybeString.(string); ok {
		return &stringValue
	} else {
		return nil
	}
}

// Extract and cast value as int
func AsInt(object interface{}, keyChain string) *int {
	maybeInt := Get(object, keyChain)
	number, ok := maybeInt.(float64)
	if !ok {
		return nil
	}
	intValue := int(number)
	reversedNumber := float64(intValue)
	if number == reversedNumber {
		return &intValue
	} else {
		return nil
	}
}

// Extract and cast value as float
func AsFloat(object interface{}, keyChain string) *float64 {
	maybeFloat := Get(object, keyChain)
	if floatValue, ok := maybeFloat.(float64); ok {
		return &floatValue
	} else {
		return nil
	}
}

// Extract and cast value as slice
func AsSlice(object interface{}, keyChain string) []interface{} {
	maybeSlice := Get(object, keyChain)
	if slice, ok := maybeSlice.([]interface{}); ok {
		return slice
	} else {
		return nil
	}
}

// Convert to json format string
func ToJsonString(v interface{}) string {
	result, _ := json.Marshal(v)
	return string(result)
}

// Convert to json format string as pretty
func ToJsonStringPretty(v interface{}) string {
	var buf bytes.Buffer
	if err := json.Indent(&buf, []byte(ToJsonString(v)), "", "  "); err != nil {
		panic(err)
	}
	return buf.String()
}
