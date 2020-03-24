package tvdb

import (
	"fmt"
	"net/http"
)

type LanguagesResponse struct {
	Data []LanguagesData `json:"data"`
}

type LanguagesData struct {
	Abbreviation string `json:"abbreviation"`
	EnglishName  string `json:"englishName"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
}

type LanguagesIdResponse struct {
	Data LanguagesData `json:"data"`
}

const Languages string = "/languages"
const LanguagesId string = "/languages/%d"

func (c *Client) Languages() (data *LanguagesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   Languages,
	})
	if err != nil {
		return
	}
	data = new(LanguagesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) LanguagesId(languageId int) (data *LanguagesIdResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(LanguagesId, languageId),
	})
	if err != nil {
		return
	}
	data = new(LanguagesIdResponse)
	err = parseResponse(resp.Body, data)
	return
}
