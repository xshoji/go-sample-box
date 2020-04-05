package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const (
	HttpContentTypeJson = "application/json;charset=utf-8"
	HttpContentTypeForm = "application/x-www-form-urlencoded;charset=utf-8"
)

func main() {

	urlBase := "http://httpbin.org"

	responseBody := DoHttpGet(urlBase + "/get?name=taro&age=20")
	log.Println(ToJsonString(responseBody))
	log.Printf("args.name               => %v\n", Get(responseBody, "args.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(responseBody, "headers.X-Amzn-Trace-Id"))
	log.Printf("\n\n\n")

	responseBody = DoHttpPost(urlBase+"/post", HttpContentTypeForm, "name=taro&age=20")
	log.Println(responseBody)
	log.Println(ToJsonString(responseBody))
	log.Printf("form.name               => %v\n", Get(responseBody, "form.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(responseBody, "headers.X-Amzn-Trace-Id"))
	log.Printf("\n\n\n")

	responseBody = DoHttpPost(urlBase+"/post", HttpContentTypeJson, `{"name":"taro", "age":20}`)
	log.Println(responseBody)
	log.Println(ToJsonString(responseBody))
	log.Printf("data                    => %v\n", Get(responseBody, "data"))
	log.Printf("json.name               => %v\n", Get(responseBody, "json.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(responseBody, "headers.X-Amzn-Trace-Id"))

}

//=======================================
// HTTP Utils
//=======================================

// Get request
func DoHttpGet(url string) interface{} {
	// GET
	log.Println(">---------- Get request start ---------->")
	log.Printf("url : %v\n", url)
	resp, err := http.Get(url)
	r := handleResponse(resp, err)
	log.Printf("responseBody : %v\n", r)
	log.Println("<----------  Get request end  ----------<")
	return r
}

// Post request
func DoHttpPost(url string, contentType string, requestBody string) interface{} {
	// POST
	log.Println(">---------- Post request start ---------->")
	log.Printf("url : %v\n", url)
	log.Printf("contentType : %v\n", contentType)
	log.Printf("requestBody : %v\n", requestBody)
	resp, err := http.Post(url, contentType, strings.NewReader(requestBody))
	r := handleResponse(resp, err)
	log.Printf("responseBody : %v\n", r)
	log.Println("<----------  Post request end  ----------<")
	return r
}

func handleResponse(resp *http.Response, err error) interface{} {
	if err != nil {
		log.Fatal(err)
		return nil
	}
	result, err := readBody(resp)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return result
}

func readBody(resp *http.Response) (interface{}, error) {
	// Response handling
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result interface{}
	json.Unmarshal(body, &result)
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Panic("resp.Body.Close() failed.")
		}
	}()
	return result, nil
}

//=======================================
// Json Utils
//=======================================

// json bytes to interface{} object
func ToJsonObject(body []byte) interface{} {
	var jsonObject interface{}
	json.Unmarshal(body, &jsonObject)
	return jsonObject
}

// get value in interface{} object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object interface{}, keyChain string) interface{} {
	keys := strings.Split(keyChain, ".")
	var result interface{}
	var exists bool
	for _, key := range keys {
		exists = false
		if _, ok := object.(map[string]interface{}); ok {
			exists = true
			object = object.(map[string]interface{})[key]
			result = object
			continue
		}
		if values, ok := object.([]interface{}); ok {
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

// to map
func ToMap(v interface{}, keys []string) map[string]interface{} {
	resultMap := make(map[string]interface{}, len(keys))
	for _, key := range keys {
		resultMap[key] = Get(v, key)
	}
	return resultMap
}

// to json string
func ToJsonString(v interface{}) string {
	result, _ := json.Marshal(v)
	return string(result)
}
