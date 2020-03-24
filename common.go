package tvdb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

type RequestArgs struct {
	Body   interface{}
	Method string
	Params url.Values
	Path   string
}

type Errors struct {
	InvalidFilters     []string `json:"invalidFilters"`
	InvalidLanguage    string   `json:"invalidLanguage"`
	InvalidQueryParams []string `json:"invalidQueryParams"`
}

type Links struct {
	First    int `json:"first"`
	Last     int `json:"last"`
	Next     int `json:"next"`
	Previous int `json:"previous"`
}

const ApplicationJson = "application/json"
const BaseURL string = "https://api.thetvdb.com"
const DefaultLanguage = "en" // English

func buildUrlPath(path string) (result string) {
	result = fmt.Sprintf("%s%s", BaseURL, path)
	return
}

func parseResponse(body io.ReadCloser, data interface{}) error {
	return json.NewDecoder(body).Decode(data)
}

func (c *Client) doRequest(args RequestArgs) (resp *http.Response, err error) {
	var body io.Reader
	if args.Body != nil {
		marshal, _ := json.Marshal(args.Body)
		body = bytes.NewBuffer(marshal)
	}
	req, err := http.NewRequest(args.Method, buildUrlPath(args.Path), body)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", ApplicationJson)
	req.Header.Set("Accept", ApplicationJson)
	req.Header.Set("Accept-Language", c.Language)
	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}
	if args.Params != nil {
		req.URL.RawQuery = args.Params.Encode()
	}
	resp, err = c.Client.Do(req)
	if err == nil && resp.StatusCode != http.StatusOK {
		err = errors.Errorf("Received error status code (%d): %s", resp.StatusCode,resp.Status)
		return
	}
	return
}
