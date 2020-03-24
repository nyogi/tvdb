package tvdb

import (
	"github.com/google/go-querystring/query"
	"net/http"
)

type SearchSeriesRequest struct {
	Name     string `url:"name,omitempty"`     // Name of the series to search for.
	ImdbId   string `url:"imdbId,omitempty"`   // IMDB id of the series
	Zap2ItId string `url:"zap2itId,omitempty"` // Zap2it ID of the series to search for.
	Slug     string `url:"slug,omitempty"`     // Slug from site URL of series (https://www.thetvdb.com/series/$SLUG)
}

type SearchSeriesResponse struct {
	Data []SearchSeriesData `json:"data"`
}

type SearchSeriesData struct {
	Aliases    []string `json:"aliases"`
	Banner     string   `json:"banner"`
	FirstAired string   `json:"firstAired"`
	ID         int      `json:"id"`
	Image      string   `json:"image"`
	Network    string   `json:"network"`
	Overview   string   `json:"overview"`
	Poster     string   `json:"poster"`
	SeriesName string   `json:"seriesName"`
	Slug       string   `json:"slug"`
	Status     string   `json:"status"`
}

// not sure when this is needed as the documented api should also be updated?
type SearchSeriesParamsResponse struct {
	Data SearchSeriesParamsData `json:"data"`
}

type SearchSeriesParamsData struct {
	Params []string `json:"params"`
}

const SearchSeries string = "/search/series"
const SearchSeriesParams string = "/search/series/params"

func (c *Client) SearchSeries(params *SearchSeriesRequest) (data *SearchSeriesResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   SearchSeries,
	})
	if err != nil {
		return
	}
	data = new(SearchSeriesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SearchSeriesParams() (data *SearchSeriesParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   SearchSeriesParams,
	})
	if err != nil {
		return
	}
	data = new(SearchSeriesParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}
