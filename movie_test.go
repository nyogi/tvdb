package tvdb_test

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func init() {
	Setup()
}

func TestClient_Movie(t *testing.T) {
	data, err := client.Movie(190)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, strings.Contains(data.Data.URL, "the-shawshank-redemption"))
}

func TestClient_MovieUpdates(t *testing.T) {
	data, err := client.MovieUpdates(LastWeekEpochStr())
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Movies)
}
