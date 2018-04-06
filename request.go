package soap

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// sendRequest makes new request to the server using the c.Method, c.URL and the body.
// body is enveloped in GetData method
func (c *Client) sendRequest() ([]byte, error) {
	req, err := http.NewRequest("POST", c.WSDL, bytes.NewBuffer(c.payload))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}

	req.ContentLength = int64(len(c.payload))
	req.Header.Add("SOAPAction", fmt.Sprintf("%s/%s", c.URL, c.Method))
	req.Header.Add("Content-Type", "text/xml")
	req.Header.Add("Accept", "text/xml")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println(string(ioutil.ReadAll(resp.Body)))
	return ioutil.ReadAll(resp.Body)
}
