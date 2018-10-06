package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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

// GetURL GetURL
func (c *Client) GetURL() string {
	return c.url
}

// callAPI callAPI
func (c *Client) callAPI(path string, httpMethod string, headers map[string]string, postData string) []byte {
	if c.GetURL() == "" {
		log.Panic("Client has not Url.")
	}

	urlFull := c.GetURL()
	if path != "" {
		urlFull = urlFull + "/" + path
	}

	// Build headers
	requestHeader := map[string]string{
		"Content-Type": "text/xml",
		"charset": "UTF-8",
	}
	// Override headers
	for k, v := range headers {
		requestHeader[k] = v;
	}

	request, _ := http.NewRequest("POST", urlFull, strings.NewReader(postData))
	for k, v := range requestHeader {
		request.Header.Add(k, v)
	}
	client := new(http.Client)
	resp, err := client.Do(request)
	if err != nil {
		log.Panic("Error occured. | ", err)
	}
	if resp.StatusCode != http.StatusOK {
		m, _ := ioutil.ReadAll(resp.Body)
		log.Print("Error | statusCode: ", resp.StatusCode)
		log.Print("Error | message: ", string(m))
		log.Panic("Error | << Status is not OK.>> ")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	return body
}

// callAPIPost callAPIPost
func (c *Client) callAPIPost(path string, headers map[string]string, postData string) []byte {
	return c.callAPI(path, "POST", headers, postData)
}

// Post Post
func (c *Client) Post(path string, headers map[string]string, postData string) *ResponseResult {
	bodyString := string(c.callAPIPost(path, headers, postData))
	return NewResponseResult(bodyString)
}
