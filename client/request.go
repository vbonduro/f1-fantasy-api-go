package client

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
)

// Request encodes data for an HTTP Request.
type Request struct {
	Endpoint string
	Payload  io.Reader
	Header   map[string][]string
}

func (req Request) makeHttpRequest(method string) (*http.Request, error) {
	httpReq, err := http.NewRequest(method, req.Endpoint, req.Payload)
	if err != nil {
		return nil, err
	}

	for key, values := range req.Header {
		for _, value := range values {
			httpReq.Header.Add(key, value)
		}
	}

	return httpReq, nil
}

// Post will send an HTTP request to the server and return back a bytes array with the result.
func (req Request) Post() (http.Header, []byte, error) {
	httpReq, err := req.makeHttpRequest("POST")
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	res, err := client.Do(httpReq)
	if err != nil {
		return nil, nil, err
	}

	if res.StatusCode != 200 {
		return nil, nil, errors.New(res.Status)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}

	return res.Header, body, nil
}
