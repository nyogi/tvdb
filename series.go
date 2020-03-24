package tvdb

import (
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

type SeriesResponse struct {
	Data   SeriesData `json:"data"`
	Errors Errors     `json:"errors,omitEmpty"`
}

type SeriesData struct {
	Added           string   `json:"added"`
	AddedBy         int      `json:"addedBy"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek"`
	AirsTime        string   `json:"airsTime"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	FanArt          string   `json:"fanart"`
	FirstAired      string   `json:"firstAired"`
	Genre           []string `json:"genre"`
	ID              int      `json:"id"`
	ImdbID          string   `json:"imdbId"`
	LastUpdated     int      `json:"lastUpdated"`
	Language        string   `json:"language"`
	Network         string   `json:"network"`
	NetworkID       string   `json:"networkId"`
	Overview        string   `json:"overview"`
	Poster          string   `json:"poster"`
	Rating          string   `json:"rating"`
	Runtime         string   `json:"runtime"`
	Season          string   `json:"season"`
	SeriesID        string   `json:"seriesId"`
	SeriesName      string   `json:"seriesName"`
	SiteRating      float64  `json:"siteRating"`
	SiteRatingCount int      `json:"siteRatingCount"`
	Slug            string   `json:"slug"`
	Status          string   `json:"status"`
	Zap2ItID        string   `json:"zap2itId"`
}

type SeriesActorsResponse struct {
	Data   []SeriesActorsData `json:"data"`
	Errors Errors             `json:"errors,omitEmpty"`
}

type SeriesActorsData struct {
	ID          int    `json:"id"`
	Image       string `json:"image"`
	ImageAdded  string `json:"imageAdded"`
	ImageAuthor int    `json:"imageAuthor"`
	LastUpdated string `json:"lastUpdated"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	SeriesID    int    `json:"seriesId"`
	SortOrder   int    `json:"sortOrder"`
}

type SeriesEpisodesRequest struct {
	Page string `url:"page,omitempty"` // Slug from site URL of series (https://www.thetvdb.com/series/$SLUG)
}

type SeriesEpisodesResponse struct {
	Data   []SeriesEpisodesData `json:"data"`
	Errors Errors               `json:"errors,omitEmpty"`
	Links  Links                `json:"links"`
}

type SeriesEpisodesData struct {
	AbsoluteNumber     int                    `json:"absoluteNumber"`
	AiredEpisodeNumber int                    `json:"airedEpisodeNumber"`
	AiredSeason        int                    `json:"airedSeason"`
	AiredSeasonId      int                    `json:"airedSeasonID"`
	AirsAfterSeason    int                    `json:"airsAfterSeason"`
	AirsBeforeEpisode  int                    `json:"airsBeforeEpisode"`
	AirsBeforeSeason   int                    `json:"airsBeforeSeason"`
	ContentRating      string                 `json:"contentRating"`
	Director           string                 `json:"director"`
	Directors          []string               `json:"directors"`
	DvdChapter         int                    `json:"dvdChapter"`
	DvdDiscid          string                 `json:"dvdDiscid"`
	DvdEpisodeNumber   int                    `json:"dvdEpisodeNumber"`
	DvdSeason          int                    `json:"dvdSeason"`
	EpisodeName        string                 `json:"episodeName"`
	Filename           string                 `json:"filename"`
	FirstAired         string                 `json:"firstAired"`
	GuestStars         []string               `json:"guestStars"`
	ID                 int                    `json:"id"`
	ImdbID             string                 `json:"imdbId"`
	IsMovie            int                    `json:"isMovie"`
	Language           SeriesEpisodesLanguage `json:"language"`
	LastUpdated        int                    `json:"lastUpdated"`
	LastUpdatedBy      int                    `json:"lastUpdatedBy"`
	Overview           string                 `json:"overview"`
	ProductionCode     string                 `json:"productionCode"`
	SeriesID           int                    `json:"seriesId"`
	ShowURL            string                 `json:"showUrl"`
	SiteRating         float64                `json:"siteRating"`
	SiteRatingCount    int                    `json:"siteRatingCount"`
	ThumbAdded         string                 `json:"thumbAdded"`
	ThumbAuthor        int                    `json:"thumbAuthor"`
	ThumbHeight        string                 `json:"thumbHeight"`
	ThumbWidth         string                 `json:"thumbWidth"`
	Writers            []string               `json:"writers"`
}

type SeriesEpisodesLinks struct {
	First    int `json:"first"`
	Last     int `json:"last"`
	Next     int `json:"next"`
	Previous int `json:"previous"`
}

type SeriesEpisodesLanguage struct {
	EpisodeName string `json:"episodeName"`
	Overview    string `json:"overview"`
}

type SeriesEpisodesQueryRequest struct {
	AbsoluteNumber string `url:"absoluteNumber,omitempty"` // Absolute number of the episode
	AiredSeason    string `url:"airedSeason,omitempty"`    // Aired season number
	AiredEpisode   string `url:"airedEpisode,omitempty"`   // Aired episode number
	DvdSeason      string `url:"dvdSeason,omitempty"`      // DVD season number
	DvdEpisode     string `url:"dvdEpisode,omitempty"`     // DVD episode number
	ImdbId         string `url:"imdbId,omitempty"`         // IMDB id of the series
	Page           string `url:"page,omitempty"`           // Page of results to fetch. Defaults to page 1 if not provided.
}

type SeriesEpisodesQueryResponse = SeriesEpisodesResponse

type SeriesEpisodesQueryParamsResponse struct {
	Data []string `json:"data"`
}

type SeriesEpisodesSummaryResponse struct {
	Data SeriesEpisodesSummaryResponseData `json:"data"`
}

type SeriesEpisodesSummaryResponseData struct {
	AiredEpisodes string   `json:"airedEpisodes"`
	AiredSeasons  []string `json:"airedSeasons"`
	DvdEpisodes   string   `json:"dvdEpisodes"`
	DvdSeasons    []string `json:"dvdSeasons"`
}

type SeriesFilterRequest struct {
	Keys string `url:"keys"` // Comma-separated list of keys to filter by (keys in `SeriesFilterData`)
}

type SeriesFilterResponse struct {
	Data   SeriesFilterData `json:"data"`
	Errors Errors           `json:"errors,omitEmpty"`
}

type SeriesFilterData struct {
	Added           string   `json:"added"`
	AddedBy         string   `json:"addedBy"`
	AirsDayOfWeek   string   `json:"airsDayOfWeek"`
	AirsTime        string   `json:"airsTime"`
	Aliases         []string `json:"aliases"`
	Banner          string   `json:"banner"`
	FanArt          string   `json:"fanart"`
	FirstAired      string   `json:"firstAired"`
	Genre           []string `json:"genre"`
	ID              int      `json:"id"`
	ImdbID          string   `json:"imdbId"`
	Language        string   `json:"language"`
	LastUpdated     int      `json:"lastUpdated"`
	Network         string   `json:"network"`
	NetworkID       string   `json:"networkId"`
	Overview        string   `json:"overview"`
	Poster          string   `json:"poster"`
	Rating          string   `json:"rating"`
	Runtime         string   `json:"runtime"`
	Season          string   `json:"season"`
	SeriesID        string   `json:"seriesId"`
	SeriesName      string   `json:"seriesName"`
	SiteRating      int      `json:"siteRating"`
	SiteRatingCount int      `json:"siteRatingCount"`
	Slug            string   `json:"slug"`
	Status          string   `json:"status"`
	Zap2ItID        string   `json:"zap2itId"`
}

type SeriesFilterParamsResponse struct {
	Data SeriesFilterParamsData `json:"data"`
}

type SeriesFilterParamsData struct {
	Params []string `json:"params"`
}

type SeriesImagesResponse struct {
	Data SeriesImagesData `json:"data"`
}

type SeriesImagesData struct {
	Fanart     int `json:"fanart"`
	Poster     int `json:"poster"`
	Season     int `json:"season"`
	SeasonWide int `json:"seasonwide"`
	Series     int `json:"series"`
}

type SeriesImagesQueryRequest struct {
	KeyType    string `url:"keyType,omitempty"`    // Type of image you're querying for (fanart, poster, etc. See ../images/query/params for more details).
	Resolution string `url:"resolution,omitempty"` // Resolution to filter by (1280x1024, for example)
	SubKey     string `url:"subKey,omitempty"`     // Subkey for the above query keys. See /series/{id}/images/query/params for more information
}

type SeriesImagesQueryResponse struct {
	Data   []SeriesImagesQueryData `json:"data"`
	Errors Errors                  `json:"errors,omitEmpty"`
}

type SeriesImagesQueryRatingsInfo struct {
	Average int `json:"average"`
	Count   int `json:"count"`
}

type SeriesImagesQueryData struct {
	FileName    string                       `json:"fileName"`
	ID          int                          `json:"id"`
	KeyType     string                       `json:"keyType"`
	Language    string                       `json:"language"`
	LanguageID  int                          `json:"languageId"`
	RatingsInfo SeriesImagesQueryRatingsInfo `json:"ratingsInfo"`
	Resolution  string                       `json:"resolution"`
	SubKey      string                       `json:"subKey"`
	Thumbnail   string                       `json:"thumbnail"`
}

type SeriesImagesQueryParamsResponse struct {
	Data []SeriesImagesQueryParamsData `json:"data"`
}

type SeriesImagesQueryParamsData struct {
	KeyType    string   `json:"keyType"`
	Resolution []string `json:"resolution"`
	SubKey     []string `json:"subKey"`
}

const Series string = "/series/%d"
const SeriesActors = Series + "/actors"
const SeriesEpisodes = Series + "/episodes"
const SeriesEpisodesQuery = SeriesEpisodes + "/query"
const SeriesEpisodesQueryParams = SeriesEpisodesQuery + "/params"
const SeriesEpisodesSummary = SeriesEpisodes + "/summary"
const SeriesFilter = Series + "/filter"
const SeriesFilterParams = SeriesFilter + "/params"
const SeriesImages = Series + "/images"
const SeriesImagesQuery = SeriesImages + "/query"
const SeriesImagesQueryParams = SeriesImagesQuery + "/params"

func (c *Client) Series(seriesId int) (data *SeriesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(Series, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesActor(seriesId int) (data *SeriesActorsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesActors, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesActorsResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesEpisodes(seriesId int, params *SeriesEpisodesRequest) (data *SeriesEpisodesResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   fmt.Sprintf(SeriesEpisodes, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesEpisodesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesEpisodesQuery(seriesId int, params *SeriesEpisodesQueryRequest) (data *SeriesEpisodesQueryResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   fmt.Sprintf(SeriesEpisodesQuery, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesEpisodesQueryResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesEpisodesQueryParams(seriesId int) (data *SeriesEpisodesQueryParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesEpisodesQueryParams, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesEpisodesQueryParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesEpisodesSummary(seriesId int) (data *SeriesEpisodesSummaryResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesEpisodesSummary, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesEpisodesSummaryResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesFilter(seriesId int, params *SeriesFilterRequest) (data *SeriesFilterResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   fmt.Sprintf(SeriesFilter, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesFilterResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesFilterParams(seriesId int) (data *SeriesFilterParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesFilterParams, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesFilterParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesImages(seriesId int) (data *SeriesImagesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesImages, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesImagesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesImagesQuery(seriesId int, params *SeriesImagesQueryRequest) (data *SeriesImagesQueryResponse, err error) {
	requestParams, err := query.Values(params)
	if err != nil {
		return
	}
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: requestParams,
		Path:   fmt.Sprintf(SeriesImagesQuery, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesImagesQueryResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) SeriesImagesQueryParams(seriesId int) (data *SeriesImagesQueryParamsResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(SeriesImagesQueryParams, seriesId),
	})
	if err != nil {
		return
	}
	data = new(SeriesImagesQueryParamsResponse)
	err = parseResponse(resp.Body, data)
	return
}
