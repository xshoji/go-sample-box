package main

import (
	"fmt"

	"github.com/xshoji/go-sample-box/httprequestsoap/client"
	"encoding/xml"
)

const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)

// User user
type SoapEnvelope struct {
	XmlnsXsi string `xml:"xmlns:xsi,attr"`
	XmlnsXsd string `xml:"xmlns:xsd,attr"`
	XmlnsSoap string `xml:"xmlns:soap,attr"`
	SoapBody SoapBody `xml:"soap:Body"`
}

type SoapBody struct {
	GetDicList GetDicList `xml:GetDicList`
}

type GetDicList struct {
	XmlnsNs string `xml:"xmlns,attr"`
	AuthTicket string `xml:AuthTicket`
}

func callSoapApi(marshaledXml string) {
	client := client.NewClient("http://public.dejizo.jp/SoapServiceV11.asmx")
	fmt.Println("[Post]")
	resultPost := client.Post(
		"",
		map[string]string{"SOAPAction": "http://MyDictionary.jp/SOAPServiceV11/GetDicList"},
		marshaledXml,
	)
	fmt.Println(resultPost.GetBody())
	fmt.Println("")
}

func main() {
	// > xml serialization - Go XML Marshalling and the Root Element - Stack Overflow
	// > https://stackoverflow.com/questions/12398925/go-xml-marshalling-and-the-root-element
	soapXml := SoapEnvelope{
		XmlnsXsi: "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXsd: "http://www.w3.org/2001/XMLSchema",
		XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		SoapBody: SoapBody{
			GetDicList: GetDicList{
				XmlnsNs: "http://MyDictionary.jp/SOAPServiceV11",
				AuthTicket: "",
			},
		},
	}
	tmp := struct {
		SoapEnvelope
		XMLName struct{} `xml:"soap:Envelope"`
	}{SoapEnvelope: soapXml}
	x, _ := xml.Marshal(tmp)
	// > When generating an XML file with Go, how do you create a doctype declaration? - Stack Overflow
	// > https://stackoverflow.com/questions/26371965/when-generating-an-xml-file-with-go-how-do-you-create-a-doctype-declaration
	marshaledXml := xml.Header + string(x);
	fmt.Println(marshaledXml)
	callSoapApi(marshaledXml)
}
