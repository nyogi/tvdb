package tvdb

import (
	"fmt"
	"net/http"
	"net/url"
)

type MoviesResponse struct {
	Data MoviesData `json:"data"`
}

type MoviesData struct {
	Artworks     []Artworks     `json:"artworks"`
	Genres       []Genres       `json:"genres"`
	ID           int            `json:"id"`
	People       People         `json:"people"`
	ReleaseDates []ReleaseDates `json:"release_dates"`
	RemoteIds    []RemoteIds    `json:"remoteids"`
	Runtime      int            `json:"runtime"`
	Trailers     []Trailers     `json:"trailers"`
	Translations []Translations `json:"translations"`
	URL          string         `json:"url"`
}

type Artworks struct {
	ArtworkType string `json:"artwork_type"`
	Height      int    `json:"height"`
	ID          string `json:"id"`
	IsPrimary   bool   `json:"is_primary"`
	Tags        string `json:"tags"`
	ThumbURL    string `json:"thumb_url"`
	URL         string `json:"url"`
	Width       int    `json:"width"`
}

type Genres struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Actors struct {
	ID              string `json:"id"`
	ImdbID          string `json:"imdb_id"`
	IsFeatured      bool   `json:"is_featured"`
	Name            string `json:"name"`
	PeopleFacebook  string `json:"people_facebook"`
	PeopleID        string `json:"people_id"`
	PeopleImage     string `json:"people_image"`
	PeopleInstagram string `json:"people_instagram"`
	PeopleTwitter   string `json:"people_twitter"`
	Role            string `json:"role"`
	RoleImage       string `json:"role_image"`
}

type Directors struct {
	ID              string `json:"id"`
	ImdbID          string `json:"imdb_id"`
	IsFeatured      bool   `json:"is_featured"`
	Name            string `json:"name"`
	PeopleFacebook  string `json:"people_facebook"`
	PeopleID        string `json:"people_id"`
	PeopleImage     string `json:"people_image"`
	PeopleInstagram string `json:"people_instagram"`
	PeopleTwitter   string `json:"people_twitter"`
	Role            string `json:"role"`
	RoleImage       string `json:"role_image"`
}

type Producers struct {
	ID              string `json:"id"`
	ImdbID          string `json:"imdb_id"`
	IsFeatured      bool   `json:"is_featured"`
	Name            string `json:"name"`
	PeopleFacebook  string `json:"people_facebook"`
	PeopleID        string `json:"people_id"`
	PeopleImage     string `json:"people_image"`
	PeopleInstagram string `json:"people_instagram"`
	PeopleTwitter   string `json:"people_twitter"`
	Role            string `json:"role"`
	RoleImage       string `json:"role_image"`
}

type Writers struct {
	ID              string `json:"id"`
	ImdbID          string `json:"imdb_id"`
	IsFeatured      bool   `json:"is_featured"`
	Name            string `json:"name"`
	PeopleFacebook  string `json:"people_facebook"`
	PeopleID        string `json:"people_id"`
	PeopleImage     string `json:"people_image"`
	PeopleInstagram string `json:"people_instagram"`
	PeopleTwitter   string `json:"people_twitter"`
	Role            string `json:"role"`
	RoleImage       string `json:"role_image"`
}

type People struct {
	Actors    []Actors    `json:"actors"`
	Directors []Directors `json:"directors"`
	Producers []Producers `json:"producers"`
	Writers   []Writers   `json:"writers"`
}

type ReleaseDates struct {
	Country string `json:"country"`
	Date    string `json:"date"`
	Type    string `json:"type"`
}

type RemoteIds struct {
	ID         string `json:"id"`
	SourceID   int    `json:"source_id"`
	SourceName string `json:"source_name"`
	SourceURL  string `json:"source_url"`
	URL        string `json:"url"`
}

type Trailers struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Translations struct {
	IsPrimary    bool   `json:"is_primary"`
	LanguageCode string `json:"language_code"`
	Name         string `json:"name"`
	Overview     string `json:"overview"`
	Tagline      string `json:"tagline"`
}

type MovieUpdatesResponse struct {
	Movies []int `json:"movies"`
}

const Movies string = "/movies/%d"
const MovieUpdates string = "/movieupdates"

func (c *Client) Movie(movieId int) (data *MoviesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Path:   fmt.Sprintf(Movies, movieId),
	})
	if err != nil {
		return
	}
	data = new(MoviesResponse)
	err = parseResponse(resp.Body, data)
	return
}

func (c *Client) MovieUpdates(startDateEpochTime string) (data *MovieUpdatesResponse, err error) {
	resp, err := c.doRequest(RequestArgs{
		Method: http.MethodGet,
		Params: url.Values{"since": {startDateEpochTime}},
		Path:   MovieUpdates,
	})
	if err != nil {
		return
	}
	data = new(MovieUpdatesResponse)
	err = parseResponse(resp.Body, data)
	return
}
