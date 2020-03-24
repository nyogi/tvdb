package some_package

import (
	"github.com/nyogi/tvdb"
	"os"
)

func ExampleLogin(){
	client := tvdb.NewClient("")
	_, err := client.Login(&tvdb.AuthenticationRequest{
		ApiKey:   os.Getenv("TVDB_API_KEY"),
		UserKey:  os.Getenv("TVDB_USER_KEY"),
		UserName: os.Getenv("TVDB_USER_NAME"),
	})

	if err != nil {
		return
	}
}
