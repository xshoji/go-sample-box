package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

func main() {
	// Yaml to object
	obj := make(map[string]any)
	_ = yaml.Unmarshal([]byte(getYaml()), &obj)
	fmt.Println(obj)
	fmt.Println(obj["employees"])
	fmt.Println(obj["employees"].([]any)[0])
	fmt.Println(obj["employees"].([]any)[0].(map[string]any)["martin"])
	fmt.Println(obj["employees"].([]any)[0].(map[string]any)["martin"].(map[string]any)["name"])
	
	// Object to yaml
	obj2 := map[string]any {
		"aaa": "bbb",
		"ccc": map[string]any{
			"ddd": "eee",
			"fff": 1111,
			"ggg": []int64{111,222,333},
			"hhh": []string{"aaa", "bbb", "ccc"},
		},
	}
	bytes, _ := yaml.Marshal(obj2)
	fmt.Println(string(bytes))
	
}

func getYaml() string {
	return `
# Employee records
employees:
-  martin:
     name: Martin D'vloper
     job: Developer
     skills:
       - python
       - perl
       - pascal
-  tabitha:
     name: Tabitha Bitumen
     job: Developer
     skills:
       - lisp
       - fortran
       - erlang
`
}
