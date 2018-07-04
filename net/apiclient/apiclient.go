package apiclient

import (
	"io/ioutil"
	"log"
	"net/http"
)

// ApiClient ApiClient
type ApiClient struct {
	Url string
}

// NewApiClient NewApiClient
func NewApiClient(url string) *ApiClient {
	client := new(ApiClient)
	client.Url = url
	return client
}

// Get Get
func (c *ApiClient) Get() string {
	if c.Url == "" {
		log.Panic("Url is empty")
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
