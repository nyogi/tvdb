package tvdb

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type UserResponse struct {
	Data UserData `json:"data"`
}

type UserData struct {
	FavoritesDisplayMode string `json:"favoritesDisplaymode"`
	Language             string `json:"language"`
	UserName             string `json:"userName"`
}

type UserFavoritesResponse struct {
	Data   UserFavoritesData `json:"data"`
	Errors Errors            `json:"errors,omitEmpty"`
}

type UserFavoritesData struct {
	Favorites []int `json:"favorites"`
}

type UserFavoritesIdDeleteResponse = UserFavoritesResponse
type UserFavoritesIdPutResponse = UserFavoritesResponse

type UserRatingsResponse struct {
	Data   []UserRatingsData `json:"data,omitEmpty"`
	Errors Errors            `json:"errors,omitEmpty"`
	Links  Links             `json:"links,omitEmpty"`
}

type UserRatingsData struct {
	Rating       int    `json:"rating"`
	RatingItemID int    `json:"ratingItemId"`
	RatingType   string `json:"ratingType"`
}

type UserRatingsQueryRequest struct {
	ItemType string `url:"itemType"`
}

type UserRatingsQueryResponse = UserRatingsResponse

type UserRatingsQueryParamsResponse struct {
	Data []string `json:"data"`
}

type UserRatingsDeleteResponse = UserRatingsResponse
type UserRatingsPutResponse = UserRatingsResponse

const User string = "/user"
const UserFavorites = User + "/favorites"
const UserFavoritesId = UserFavorites + "/%d"
const UserRatings = User + "/ratings"
const UserRatingsQuery = UserRatings + "/query"
const UserRatingsQueryParams = UserRatingsQuery + "/params"
const UserRatingsDelete = UserRatings + "/%s/%d"
const UserRatingsPut = UserRatingsDelete + "/%d"

func (c *Client) User() (data *UserResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   User,
	})
	if err != nil {
		return
	}
	data = new(UserResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserFavorites() (data *UserFavoritesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   UserFavorites,
	})
	if err != nil {
		return
	}
	data = new(UserFavoritesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserFavoritesIdDelete(seriesId int) (data *UserFavoritesIdDeleteResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf(UserFavoritesId, seriesId),
	})
	if err != nil {
		return
	}
	data = new(UserFavoritesIdDeleteResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserFavoritesIdPut(seriesId int) (data *UserFavoritesIdPutResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodPut,
		Path:   fmt.Sprintf(UserFavoritesId, seriesId),
	})
	if err != nil {
		return
	}
	data = new(UserFavoritesIdPutResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserRatings() (data *UserRatingsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   UserRatings,
	})
	if err != nil {
		return
	}
	data = new(UserRatingsResponse)
	err = parseResponse(resp.Body, data)
	return
}

// as of 3/28/20 this is not working.
func (c *Client) UserRatingsQuery(params *UserRatingsQueryRequest) (data *UserRatingsQueryResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   UserRatingsQuery,
	})
	if err != nil {
		return
	}
	data = new(UserRatingsQueryResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserRatingsQueryParams() (data *UserRatingsQueryParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   UserRatingsQueryParams,
	})
	if err != nil {
		return
	}
	data = new(UserRatingsQueryParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserRatingsDelete(itemType string, itemId int) (data *UserRatingsDeleteResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodDelete,
		Path:   fmt.Sprintf(UserRatingsDelete, itemType, itemId),
	})
	if err != nil {
		return
	}
	data = new(UserRatingsDeleteResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) UserRatingsPut(itemType string, itemId int, itemRating int) (data *UserRatingsPutResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodPut,
		Path:   fmt.Sprintf(UserRatingsPut, itemType, itemId, itemRating),
	})
	if err != nil {
		return
	}
	data = new(UserRatingsPutResponse)
	err = parseResponse(resp.Body, data)
	return
}
