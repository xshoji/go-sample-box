package jsonutil

import (
	"encoding/json"
	"strconv"
	"strings"
)

// interface{}型のjsonオブジェクトからキー指定で値を取り出す（object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb"）
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

// 値をjson形式の文字列に変換する
func ToJsonString(v interface{}) string {
	result, _ := json.Marshal(v)
	return string(result)
}
