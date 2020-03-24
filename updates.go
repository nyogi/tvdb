package tvdb

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type UpdatedQueryRequest struct {
	FromTime string `url:"fromTime"`         // Epoch time to start your date range.
	ToTime   string `url:"toTime,omitempty"` // Epoch time to end your date range. Must be one week from fromTime.
}

type UpdatedQueryResponse struct {
	Data   []UpdatedQueryData `json:"data"`
	Errors Errors             `json:"errors,omitEmpty"`
}

type UpdatedQueryData struct {
	ID          int `json:"id"`
	LastUpdated int `json:"lastUpdated"`
}

type UpdatedQueryParamsResponse struct {
	Data []string `json:"data"`
}

const UpdatedQuery string = "/updated/query"
const UpdatedQueryParams = UpdatedQuery + "/params"

func (c *Client) UpdatedQuery(params *UpdatedQueryRequest) (data *UpdatedQueryResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   UpdatedQuery,
	})
	if err != nil {
		return
	}
	data = new(UpdatedQueryResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UpdatedQueryParams() (data *UpdatedQueryParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   UpdatedQueryParams,
	})
	if err != nil {
		return
	}
	data = new(UpdatedQueryParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}
