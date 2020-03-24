package tvdb

import (
	"net/http"
)

// Authentication struct store all data of a summary.
type AuthenticationResponse struct {
	Token string `json:"token"`
}

type AuthenticationRequest struct {
	ApiKey   string `json:"apikey"`
	UserKey  string `json:"userkey"`
	UserName string `json:"username"`
}

const AuthenticationLogin string = "/login"
const AuthenticationRefreshToken string = "/refresh_token"

func (c *Client) Login(requestParams *AuthenticationRequest) (data *AuthenticationResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Body:   requestParams,
		Path:   AuthenticationLogin,
		Method: http.MethodPost,
	})
	if err != nil {
		return
	}
	data, err = c.saveToken(resp)
	defer resp.Body.Close()
	return
}

func (c *Client) RefreshToken() (data *AuthenticationResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Path:   AuthenticationRefreshToken,
		Method: http.MethodGet,
	})
	if err != nil {
		return
	}
	data, err = c.saveToken(resp)
	defer resp.Body.Close()
	return
}

func (c *Client) saveToken(resp *http.Response) (data *AuthenticationResponse, err error) {
	data = new(AuthenticationResponse)
	err = parseResponse(resp.Body, data)
	if err != nil {
		return
	}
	c.Token = data.Token
	return
}
