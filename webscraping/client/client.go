package client

import (
	"encoding/xml"
	"fmt"
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
func (r *ResponseResult) GetBodyAsObject() *Rss {
	rss := new(Rss)
	if err := xml.Unmarshal([]byte(r.responseBody), &rss); err != nil {
		fmt.Println("XML Unmarshal error:", err)
		return nil
	}
	return rss
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

// GetURL GetURL
func (c *Client) GetURL() string {
	return c.url
}

// callAPI callAPI
func (c *Client) callAPI(path string, httpMethod string, postData map[string][]string) []byte {
	if c.GetURL() == "" {
		log.Panic("Client has not Url.")
	}

	urlFull := c.GetURL()
	if path != "" {
		urlFull = urlFull + "/" + path
	}

	var resp *http.Response
	var err error
	if httpMethod == "GET" {
		resp, err = http.Get(urlFull)
	} else if httpMethod == "POST" {
		// - [How to send a POST request in Go? - Stack Overflow](https://stackoverflow.com/questions/24493116/how-to-send-a-post-request-in-go)
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

// callAPIGet callAPIGet
func (c *Client) callAPIGet(path string) []byte {
	return c.callAPI(path, "GET", nil)
}

// callAPIPost callAPIPost
func (c *Client) callAPIPost(path string, postData map[string][]string) []byte {
	return c.callAPI(path, "POST", postData)
}

// Post Post
func (c *Client) Post(path string, postData map[string][]string) *ResponseResult {
	bodyString := string(c.callAPIPost(path, postData))
	return NewResponseResult(bodyString)
}

// Get Get
func (c *Client) Get(path string) *ResponseResult {
	bodyString := string(c.callAPIGet(path))
	return NewResponseResult(bodyString)
}
