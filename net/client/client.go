package client

import (
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
