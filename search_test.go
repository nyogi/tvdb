package tvdb_test

import (
	"github.com/nyogi/tvdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Setup()
}

func TestClient_SearchSeries(t *testing.T) {
	data, err := client.SearchSeries(&tvdb.SearchSeriesRequest{
		Slug: ChuckSlug,
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data[0].ID, ChuckSeriesId)

	data, err = client.SearchSeries(&tvdb.SearchSeriesRequest{
		Name: ChuckSlug,
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data[0].ID, ChuckSeriesId)
}

func TestClient_SearchSeriesParams(t *testing.T) {
	data, err := client.SearchSeriesParams()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	hasNameParam := false
	for _, d := range data.Data.Params {
		hasNameParam = hasNameParam || d == "name"
	}
	assert.True(t, hasNameParam)
}
