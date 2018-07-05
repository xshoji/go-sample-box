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

// GetAsObject GetAsObject
func (c *Client) GetAsObject() LatestBlock {
	// - [The Go Playground](https://play.golang.com/p/QXdlVsi166)
	bytes := c.getBodyBytesWithPath("")
	var latestBlock LatestBlock
	json.Unmarshal(bytes, &latestBlock)
	return latestBlock
}
