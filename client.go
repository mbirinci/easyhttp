package easyhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Client struct {
	*http.Client
}

type Response struct {
	*http.Response
	RawBody []byte
}

func(c *Client) EasyGet(url string) (*Response, error) {

	request, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.Do(request)

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


func (r *Response) JSON(val interface{}) error {

	if err := json.Unmarshal(r.RawBody, val); err != nil {
		return err
	}

	return nil

}
