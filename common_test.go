package tvdb_test

import (
	"fmt"
	"github.com/nyogi/tvdb"
	"os"
	"strconv"
	"time"
)

const ChuckSeriesId int = 80348
const ChuckSlug string = "chuck"

var client *tvdb.Client

func Setup() {
	if client == nil {
		client = tvdb.NewClient("")
		_, _ = client.Login(&tvdb.AuthenticationRequest{
			ApiKey:   os.Getenv("TVDB_API_KEY"),
			UserKey:  os.Getenv("TVDB_USER_KEY"),
			UserName: os.Getenv("TVDB_USER_NAME"),
		})
		fmt.Println("initialized...")
	}
}

func LastWeekEpochStr() (result string) {
	lastWeekEpochTime := time.Now().AddDate(0, 0, -7).Unix()
	result = strconv.FormatInt(lastWeekEpochTime, 10)
	return
}
