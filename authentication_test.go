package tvdb_test

import (
	"github.com/nyogi/tvdb"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestClient_Login(t *testing.T) {
	c := tvdb.NewClient("")
	resp, err := c.Login(&tvdb.AuthenticationRequest{
		ApiKey:   os.Getenv("TVDB_API_KEY"),
		UserKey:  os.Getenv("TVDB_USER_KEY"),
		UserName: os.Getenv("TVDB_USER_NAME"),
	})
	assert.Nil(t, err)
	assert.NotEmpty(t, resp.Token)
	assert.NotEmpty(t, c.Token)
	assert.Equal(t, resp.Token, c.Token)
}

func TestClient_RefreshToken(t *testing.T) {
	c := tvdb.NewClient("")
	_, _ = c.Login(&tvdb.AuthenticationRequest{
		ApiKey:   os.Getenv("TVDB_API_KEY"),
		UserKey:  os.Getenv("TVDB_USER_KEY"),
		UserName: os.Getenv("TVDB_USER_NAME"),
	})
	oldToken := c.Token
	// there needs to a period of time between creation of token and refresh or else the token will be the same.
	// not sure what the time period is set to as it's not documented.
	time.Sleep(5 * time.Second)

	resp, err := c.RefreshToken()
	assert.Nil(t, err)
	assert.NotEmpty(t, resp.Token)
	assert.NotEmpty(t, c.Token)
	assert.Equal(t, resp.Token, c.Token)
	assert.NotEqual(t, oldToken, c.Token)
}
