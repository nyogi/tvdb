package tvdb_test

import (
	"github.com/nyogi/tvdb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Setup()
}

func TestClient_UpdatedQuery(t *testing.T) {
	data, err := client.UpdatedQuery(&tvdb.UpdatedQueryRequest{
		FromTime: LastWeekEpochStr(),
	})
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)

	data, err = client.UpdatedQuery(&tvdb.UpdatedQueryRequest{})
	assert.NotNil(t, data.Errors)
}

func TestClient_UpdatedQueryParams(t *testing.T) {
	data, err := client.UpdatedQueryParams()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data.Data)
}
