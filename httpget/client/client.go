package client

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// Client Client
type Client struct {
	Url string
}

// NewClient NewClient
func NewClient(url string) *Client {
	client := new(Client)
	client.Url = url
	return client
}

// GetWithPathAsString GetWithPathAsString
func (c *Client) getBodyBytesWithPath(path string) []byte {
	if c.Url == "" {
		log.Panic("Url is")
	}

	url := c.Url
	if path != "" {
		url = url + "/" + path
	}

	resp, err := http.Get(url)
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

// GetWithPathAsString GetWithPathAsString
func (c *Client) GetWithPathAsString(path string) string {
	bodyString := string(c.getBodyBytesWithPath(path))
	return bodyString
}

// GetAsString GetAsString
func (c *Client) GetAsString() string {
	return c.GetWithPathAsString("")
}

// GetWithPathAsObject GetWithPathAsObject
// - [golang は ゆるふわに JSON を扱えまぁす! — KaoriYa](https://www.kaoriya.net/blog/2016/06/25/)
func (c *Client) GetWithPathAsObject(path string) interface{} {
	var result interface{}
	json.Unmarshal(c.getBodyBytesWithPath(path), &result)
	return result
}

// GetAsObject GetAsObject
func (c *Client) GetAsObject() interface{} {
	return c.GetWithPathAsObject("")
}
