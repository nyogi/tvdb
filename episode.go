package tvdb

import (
	"fmt"
	"net/http"
)

type EpisodeResponse struct {
	Data   EpisodeData `json:"data"`
	Errors Errors      `json:"errors,omitEmpty"`
}

type EpisodeData struct {
	AbsoluteNumber     int      `json:"absoluteNumber"`
	AiredEpisodeNumber int      `json:"airedEpisodeNumber"`
	AiredSeason        int      `json:"airedSeason"`
	AirsAfterSeason    int      `json:"airsAfterSeason"`
	AirsBeforeEpisode  int      `json:"airsBeforeEpisode"`
	AirsBeforeSeason   int      `json:"airsBeforeSeason"`
	Director           string   `json:"director"`
	Directors          []string `json:"directors"`
	DvdChapter         int      `json:"dvdChapter"`
	DvdDiscId          string   `json:"dvdDiscid"`
	DvdEpisodeNumber   float64  `json:"dvdEpisodeNumber"`
	DvdSeason          int      `json:"dvdSeason"`
	EpisodeName        string   `json:"episodeName"`
	Filename           string   `json:"filename"`
	FirstAired         string   `json:"firstAired"`
	GuestStars         []string `json:"guestStars"`
	Id                 int      `json:"id"`
	ImdbId             string   `json:"imdbId"`
	LastUpdated        int      `json:"lastUpdated"`
	LastUpdatedBy      int      `json:"lastUpdatedBy"`
	Overview           string   `json:"overview"`
	ProductionCode     string   `json:"productionCode"`
	SeriesId           int      `json:"seriesId"`
	ShowURL            string   `json:"showURL"`
	SiteRating         float32  `json:"siteRating"`
	SiteRatingCount    int      `json:"siteRatingCount"`
	ThumbAdded         string   `json:"thumbAdded"`
	ThumbAuthor        int      `json:"thumbAuthor"`
	ThumbHeight        string   `json:"thumbHeight"`
	ThumbWidth         string   `json:"thumbWidth"`
	Writers            []string `json:"writers"`
}

const EpisodesId string = "/episodes/%d"

func (c *Client) EpisodesId(episodeId int) (data *EpisodeResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(EpisodesId, episodeId),
	})
	if err != nil {
		return
	}
	data = new(EpisodeResponse)
	err = parseResponse(resp.Body, data)
	return
}
