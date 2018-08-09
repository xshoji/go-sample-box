package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

// ResponseResult ResponseResult
type ResponseResult struct {
	responseBody string
}

// NewResponseResult NewResponseResult
func NewResponseResult(responseBody string) *ResponseResult {
	result := new(ResponseResult)
	result.responseBody = responseBody
	return result
}

// GetBody GetBody
func (r *ResponseResult) GetBody() string {
	return r.responseBody
}

// GetBodyAsObject GetBodyAsObject
func (r *ResponseResult) GetBodyAsObject() interface{} {
	var result interface{}
	json.Unmarshal([]byte(r.responseBody), &result)
	return result
}

// Client Client
type Client struct {
	url string
}

// NewClient NewClient
func NewClient(url string) *Client {
	client := new(Client)
	client.url = url
	return client
}

// GetUrl GetUrl
func (c *Client) GetUrl() string {
	return c.url
}

// callApi callApi
func (c *Client) callApi(path string, httpMethod string, postData map[string][]string) []byte {
	if c.GetUrl() == "" {
		log.Panic("Client has not Url.")
	}

	urlFull := c.GetUrl()
	if path != "" {
		urlFull = urlFull + "/" + path
	}

	var resp *http.Response
	var err error
	if httpMethod == "GET" {
		resp, err = http.Get(urlFull)
	} else if httpMethod == "POST" {
		form := url.Values(postData)
		resp, err = http.Post(urlFull, "text/plain", strings.NewReader(form.Encode()))
	} else {
		log.Panic("httpMethod:" + httpMethod + " is unkown.")
	}

	if err != nil {
		log.Panic("Error occured. | ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Panic("Status is not OK. | ", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

// callApiGet callApiGet
func (c *Client) callApiGet(path string) []byte {
	return c.callApi(path, "GET", nil)
}

// callApiPost callApiPost
func (c *Client) callApiPost(path string, postData map[string][]string) []byte {
	return c.callApi(path, "POST", postData)
}

// Post Post
func (c *Client) Post(path string, postData map[string][]string) *ResponseResult {
	bodyString := string(c.callApiPost(path, postData))
	return NewResponseResult(bodyString)
}

// Get Get
func (c *Client) Get(path string) *ResponseResult {
	bodyString := string(c.callApiGet(path))
	return NewResponseResult(bodyString)
}
