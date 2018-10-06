package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/httprequestsoap/client"
)

func main() {
	client := client.NewClient("http://public.dejizo.jp/SoapServiceV11.asmx")
	fmt.Println("[Post]")
	resultPost := client.Post(
	"",
	map[string]string{"SOAPAction": "http://MyDictionary.jp/SOAPServiceV11/GetDicList"},
	`<?xml version="1.0" encoding="utf-8"?>
	<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
	  <soap:Body>
	    <GetDicList xmlns="http://MyDictionary.jp/SOAPServiceV11">
	      <AuthTicket></AuthTicket>
	    </GetDicList>
	  </soap:Body>
	</soap:Envelope>`)
	fmt.Println(resultPost.GetBody())
	fmt.Println("")
}
