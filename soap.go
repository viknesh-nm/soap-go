package soap

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// Params type is used to set the params in soap request
type Params map[string]interface{}

// Client struct hold all the informations about WSDL,
// request and response of the server
type Client struct {
	WSDL        string
	URL         string
	Method      string
	Params      Params
	Definitions *WSDLDefinitions
	Body        []byte
	payload     []byte
}

// SEnvelope -
type SEnvelope struct {
	XMLName struct{} `xml:"Envelope"`
	Body    SBody
}

// SBody -
type SBody struct {
	XMLName  struct{} `xml:"Body"`
	Contents []byte   `xml:",innerxml"`
}

// NewClient return new *Client to handle the requests with the WSDL
func NewClient(wsdlURL string) (*Client, error) {
	d, err := GetWsdlDefinitions(wsdlURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		WSDL:        wsdlURL,
		URL:         strings.TrimSuffix(d.TargetNamespace, "/"),
		Definitions: d,
	}

	return c, nil
}

// GetData call's the method m with Params p
func (c *Client) GetData(m string, p Params) (result string, err error) {
	c.Method = m
	c.Params = p

	c.payload, err = xml.MarshalIndent(c, "", "")
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	b, err := c.sendRequest()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var soap SEnvelope
	err = xml.Unmarshal(b, &soap)
	if err != nil {
		return "", err
	}
	return string(soap.Body.Contents), nil
}
