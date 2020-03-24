package tvdb

import (
	"net/http"
)

type Client struct {
	Client   *http.Client
	Language string
	Token    string
}

func NewClient(language string) (client *Client) {
	if language == "" {
		language = DefaultLanguage
	}
	client = &Client{
		Language: language,
		Token:    "",
		Client:   &http.Client{},
	}
	return
}
