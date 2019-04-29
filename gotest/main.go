package main

import "regexp"

func main() {
}

func ReplaceString(str string) (string, string) {
	rep := regexp.MustCompile(`\s*\[\s*`)
	result := rep.Split(str, -1)
	if len(result) == 2 {
		return result[0], "[" + result[1]
	}
	return str, ""
}
