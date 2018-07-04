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

// GetAsString GetAsString
func (c *Client) GetAsString() string {
	if c.Url == "" {
		log.Panic("Url is")
	}

	resp, err := http.Get(c.Url)
	if err != nil {
		log.Panic("Error occured. | ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Panic("Status is not OK. | ", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	bodyString := string(body)
	return bodyString
}

// - [golang は ゆるふわに JSON を扱えまぁす! — KaoriYa](https://www.kaoriya.net/blog/2016/06/25/)
// GetAsObject GetAsObject
func (c *Client) GetAsObject() interface{} {
	if c.Url == "" {
		log.Panic("Url is")
	}

	resp, err := http.Get(c.Url)
	if err != nil {
		log.Panic("Error occured. | ", err)
	}
	if resp.StatusCode != http.StatusOK {
		log.Panic("Status is not OK. | ", resp.StatusCode)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	var result interface{}
	json.Unmarshal(body, &result)
	return result
}
