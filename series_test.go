package tvdb_test

import (
	"github.com/nyogi/tvdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Setup()
}

func TestClient_Series(t *testing.T) {
	data, err := client.Series(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.Slug, ChuckSlug)
}

func TestClient_SeriesActor(t *testing.T) {
	data, err := client.SeriesActor(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
	hasChuck := false
	for _, d := range data.Data {
		hasChuck = hasChuck || d.Role == "Chuck Bartowski"
	}
	assert.True(t, hasChuck)
}

func TestClient_SeriesEpisodes(t *testing.T) {
	data, err := client.SeriesEpisodes(ChuckSeriesId, nil)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
	data, err = client.SeriesEpisodes(ChuckSeriesId, &tvdb.SeriesEpisodesRequest{
		Page: "2",
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
}

func TestClient_SeriesEpisodesQuery(t *testing.T) {
	data, err := client.SeriesEpisodesQuery(ChuckSeriesId, &tvdb.SeriesEpisodesQueryRequest{
		AiredSeason: "1",
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
	assert.Equal(t, len(data.Data), 13)
}

func TestClient_SeriesEpisodesQueryParams(t *testing.T) {
	data, err := client.SeriesEpisodesQueryParams(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
	hasSeriesId := false
	for _, d := range data.Data {
		hasSeriesId = hasSeriesId || d == "seriesId"
	}
	assert.True(t, hasSeriesId)
}

func TestClient_SeriesEpisodesSummary(t *testing.T) {
	data, err := client.SeriesEpisodesSummary(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.AiredEpisodes, "102")
}

func TestClient_SeriesFilter(t *testing.T) {
	data, err := client.SeriesFilter(ChuckSeriesId, &tvdb.SeriesFilterRequest{
		Keys: "season",
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
	assert.Equal(t, data.Data.Season, "5")
}

func TestClient_SeriesFilterParams(t *testing.T) {
	data, err := client.SeriesFilterParams(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, len(data.Data.Params) > 0)
}

func TestClient_SeriesImages(t *testing.T) {
	data, err := client.SeriesImages(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.Fanart, 43)
}

func TestClient_SeriesImagesQuery(t *testing.T) {
	data, err := client.SeriesImagesQuery(ChuckSeriesId, &tvdb.SeriesImagesQueryRequest{
		KeyType: "series",
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, len(data.Data) > 0)
}

func TestClient_SeriesImagesQueryParams(t *testing.T) {
	data, err := client.SeriesImagesQueryParams(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, len(data.Data) > 0)
}
