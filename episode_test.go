package tvdb_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Setup()
}

func TestClient_Episodes(t *testing.T) {
	data, err := client.EpisodesId(6840191)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.SeriesId, 70366)
}
