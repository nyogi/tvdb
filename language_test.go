package tvdb_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	Setup()
}

func TestClient_Languages(t *testing.T) {
	data, err := client.Languages()
	assert.Nil(t, err)
	assert.NotNil(t, data)
	hasEnglishLanguage := false
	for _, d := range data.Data {
		hasEnglishLanguage = hasEnglishLanguage || d.Abbreviation == "en"
	}
	assert.True(t, hasEnglishLanguage)
}

func TestClient_LanguagesId(t *testing.T) {
	data, err := client.LanguagesId(7)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, data.Data.Abbreviation, "en")
}
