package easyhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Client is extended stdlib *http.Client
type Client struct {
	// embed *http.Client
	*http.Client
}

// Response is extended *http.Response
type Response struct {
	// embed *http.Response
	*http.Response
	// RawBody holds body data as byte that read from stream
	RawBody []byte
}

// Options are the options available to a request
type Options struct {
	// Header contains the request header fields
	Header map[string]string
}


// EasyGet make request and read body stream
// It returns extended Response that acts as *http.Response
func (c *Client) EasyGet(url string, opts *Options) (*Response, error) {

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	if opts != nil {
		for k, v := range opts.Header {
			req.Header.Add(k, v)
		}
	}

	resp, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		return nil, err
	}

	return &Response{
		resp,
		body,
	}, nil
}

// JSON converts response body to passed value
func (r *Response) JSON(val interface{}) error {

	if err := json.Unmarshal(r.RawBody, val); err != nil {
		return err
	}

	return nil

}

// Text returns response body as string
func (r *Response) Text() string {

	return string(r.RawBody)

}
