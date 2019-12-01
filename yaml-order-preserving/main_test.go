package main

import (
	"gopkg.in/yaml.v2"
	"testing"
)

func TestGet(t *testing.T) {
	//---------------------
	// init
	yamlString := `
aaa:
  bbb:
    bbb1: bbb
    bbb2: 111
  ccc:
    - 1
    - 2
    - c
`
	var expectedKey interface{}
	var expectedValue interface{}
	var expectedValueList []interface{}
	var actualKey interface{}
	var actualValue interface{}
	var actualValueList []interface{}

	node := &yaml.MapSlice{}
	yaml.Unmarshal([]byte(yamlString), node)
	orderedMapSlice := NewOrderedMapSlice(node)

	//---------------------
	// success 1
	expectedKey = `bbb1`
	expectedValue = `bbb`
	actualKey = orderedMapSlice.Get(`aaa`).Get(`bbb`).Get(`bbb1`).Key().(string)
	actualValue = orderedMapSlice.Get(`aaa`).Get(`bbb`).Get(`bbb1`).Value().(string)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Fatalf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	}

	//---------------------
	// success 1
	expectedKey = `ccc`
	expectedValueList = []interface{}{1, 2, `c`}
	actualKey = orderedMapSlice.Get(`aaa`).Get(`ccc`).Key().(string)
	actualValueList = orderedMapSlice.Get(`aaa`).Get(`ccc`).Value().([]interface{})
	if actualKey != expectedKey || !compareSlice(actualValueList, expectedValueList) {
		t.Fatalf("actualKey:%v, expectedKey:%v | actualValueList:%v, expectedValueList:%v\n", actualKey, expectedKey, actualValueList, expectedValueList)
	}

	//---------------------
	// success 1
	expectedKey = `bbb2`
	expectedValue = 111
	actualKey = orderedMapSlice.Get(`aaa`).Get(`bbb`).Get(`bbb2`).Key().(string)
	actualValue = orderedMapSlice.Get(`aaa`).Get(`bbb`).Get(`bbb2`).Value().(int)
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Fatalf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	}

	//---------------------
	// success 2
	expectedKey = nil
	expectedValue = nil
	actualKey = orderedMapSlice.Get(`xxx`).Key()
	actualValue = orderedMapSlice.Get(`xxx`).Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Fatalf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	}

	//---------------------
	// success 3
	expectedKey = nil
	expectedValue = nil
	actualKey = orderedMapSlice.Get(`xxx`).Get(`yyy`).Get(`zzz`).Key()
	actualValue = orderedMapSlice.Get(`xxx`).Get(`yyy`).Get(`zzz`).Value()
	if actualKey != expectedKey || actualValue != expectedValue {
		t.Fatalf("actualKey:%v, expectedKey:%v | actualValue:%v, expectedValue:%v\n", actualKey, expectedKey, actualValue, expectedValue)
	}
}

func compareSlice(a []interface{}, b []interface{}) bool {
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
