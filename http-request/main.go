package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

// Debugging HTTP Client requests with Go · Jamie Tanna | Software Engineer
// https://www.jvt.me/posts/2023/03/11/go-debug-http/
type loggingTransport struct{}

func (s *loggingTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	bytes, _ := httputil.DumpRequestOut(r, true)

	resp, err := http.DefaultTransport.RoundTrip(r)

	respBytes, _ := httputil.DumpResponse(resp, true)
	bytes = append(bytes, []byte("\n\n")...)
	bytes = append(bytes, respBytes...)

	fmt.Printf("%s\n", bytes)

	return resp, err
}

const (
	HttpContentTypeHeader = "Content-Type"
)

var HttpHeaderEmptyMap = make(map[string]string, 0)
var HttpHeaderContentTypeForm = map[string]string{HttpContentTypeHeader: "application/x-www-form-urlencoded;charset=utf-8"}
var HttpHeaderContentTypeJson = map[string]string{HttpContentTypeHeader: "application/json;charset=utf-8"}

func init() {
	// log config
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
}

func main() {

	urlBase := "http://httpbin.org"
	client := http.Client{
		Transport: &loggingTransport{},
	}

	log.Printf("\n<< HttpGet >>")
	resp, err := HttpGet(client, urlBase+"/get?name=taro&age=20")
	body := handleResponse(resp, err)
	jsonBody := ToJsonObject(body)
	log.Println(ToJsonString(jsonBody))
	log.Printf("args.name               => %v\n", Get(jsonBody, "args.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(jsonBody, "headers.X-Amzn-Trace-Id"))
	log.Printf("\n\n\n")

	log.Printf("\n<< HttpGetWithHeaders >>")
	resp, err = HttpGetWithHeaders(client, urlBase+"/get?name=taro&age=20", map[string]string{
		"X-Test-Custom-Header": "test",
	})
	body = handleResponse(resp, err)
	jsonBody = ToJsonObject(body)
	log.Println(ToJsonString(jsonBody))
	log.Printf("args.name               => %v\n", Get(jsonBody, "args.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(jsonBody, "headers.X-Amzn-Trace-Id"))
	log.Printf("\n\n\n")

	log.Printf("\n<< HttpPostWithHeaders >>")
	resp, err = HttpPostWithHeaders(client, urlBase+"/post", HttpHeaderContentTypeForm, "name=taro&age=20")
	body = handleResponse(resp, err)
	jsonBody = ToJsonObject(body)
	log.Println(ToJsonString(jsonBody))
	log.Printf("form.name               => %v\n", Get(jsonBody, "form.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(jsonBody, "headers.X-Amzn-Trace-Id"))
	log.Printf("\n\n\n")

	log.Printf("\n<< HttpPostWithHeaders >>")
	resp, err = HttpPostWithHeaders(client, urlBase+"/post", HttpHeaderContentTypeJson, `{"name":"taro", "age":20}`)
	body = handleResponse(resp, err)
	jsonBody = ToJsonObject(body)
	log.Println(ToJsonString(jsonBody))
	log.Printf("data                    => %v\n", Get(jsonBody, "data"))
	log.Printf("json.name               => %v\n", Get(jsonBody, "json.name"))
	log.Printf("headers.X-Amzn-Trace-Id => %v\n", Get(jsonBody, "headers.X-Amzn-Trace-Id"))

}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// HttpGet Get =======================================
// HTTP Utils
// =======================================
// HttpGet Get request
func HttpGet(client http.Client, url string) (*http.Response, error) {
	return client.Get(url)
}

// HttpGetWithHeaders Get request with headers
func HttpGetWithHeaders(client http.Client, url string, headers map[string]string) (resp *http.Response, err error) {
	return DoHttpRequest(client, "GET", url, headers, nil)
}

// HttpPost POST request
func HttpPost(client http.Client, url string, requestBody string) (*http.Response, error) {
	return client.Post(url, HttpContentTypeHeader, strings.NewReader(requestBody))
}

// HttpPostWithHeaders POST request
func HttpPostWithHeaders(client http.Client, url string, headers map[string]string, requestBody string) (*http.Response, error) {
	return DoHttpRequest(client, "POST", url, headers, strings.NewReader(requestBody))
}

func DoHttpRequest(client http.Client, method string, url string, headers map[string]string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	handleError(err)
	for key, value := range headers {
		log.Printf("header [%s] : %s\n", key, value)
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

func handleResponse(resp *http.Response, err error) []byte {
	handleError(err)
	responseBodyBytes, err := ioutil.ReadAll(resp.Body)
	handleError(err)
	return responseBodyBytes
}

//=======================================
// Json Utils
//=======================================

// json bytes to any object
func ToJsonObject(body []byte) any {
	var jsonObject any
	json.Unmarshal(body, &jsonObject)
	return jsonObject
}

// get value in any object [ example : object["aaa"][0]["bbb"] -> keyChain: "aaa.0.bbb" ]
func Get(object any, keyChain string) any {
	keys := strings.Split(keyChain, ".")
	var result any
	var exists bool
	for _, key := range keys {
		exists = false
		if _, ok := object.(map[string]any); ok {
			exists = true
			object = object.(map[string]any)[key]
			result = object
			continue
		}
		if values, ok := object.([]any); ok {
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
func ToMap(v any, keys []string) map[string]any {
	resultMap := make(map[string]any, len(keys))
	for _, key := range keys {
		resultMap[key] = Get(v, key)
	}
	return resultMap
}

// to json string
func ToJsonString(v any) string {
	result, _ := json.Marshal(v)
	return string(result)
}
