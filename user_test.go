package tvdb_test

import (
	"github.com/nyogi/tvdb"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	Setup()
}

func TestClient_User(t *testing.T) {
	data, err := client.User()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.UserName, os.Getenv("TVDB_USER_NAME"))
}

func TestClient_UserFavorites(t *testing.T) {
	// add
	data, err := client.UserFavoritesIdPut(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, verifyFavorite(data.Data.Favorites))

	// get
	data, err = client.UserFavorites()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, verifyFavorite(data.Data.Favorites))

	// delete
	data, err = client.UserFavoritesIdDelete(ChuckSeriesId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.False(t, verifyFavorite(data.Data.Favorites))
}

func verifyFavorite(data []int) (result bool) {
	for _, d := range data {
		result = result || d == ChuckSeriesId
	}
	return
}

func TestClient_UserRatings(t *testing.T) {
	itemType := "series"
	itemId := ChuckSeriesId
	itemRating := 10

	// add
	data, err := client.UserRatingsPut(itemType, itemId, itemRating)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, verifyRatings(data.Data, itemId, itemRating))

	// get
	data, err = client.UserRatings()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, verifyRatings(data.Data, itemId, itemRating))

	// delete
	data, err = client.UserRatingsDelete(itemType, itemId)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.False(t, verifyRatings(data.Data, itemId, itemRating))
}

func verifyRatings(data []tvdb.UserRatingsData, itemId int, itemRating int) (result bool) {
	for _, d := range data {
		result = result || (d.RatingItemID == itemId && d.Rating == itemRating)
	}
	return
}

func TestClient_UserRatingsQuery(t *testing.T) {
	data, err := client.UserRatingsQuery(&tvdb.UserRatingsQueryRequest{
		ItemType: "series",
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
}

func TestClient_UserRatingsQueryParams(t *testing.T) {
	data, err := client.UserRatingsQueryParams()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.True(t, len(data.Data) > 0)
}
