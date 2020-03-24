# tvdb
TheTVDB API v3

https://api.thetvdb.com/swagger

## How to use
```
    client := tvdb.NewClient("")
    _, err := client.Login(&tvdb.AuthenticationRequest{
        ApiKey:   os.Getenv("TVDB_API_KEY"),
        UserKey:  os.Getenv("TVDB_USER_KEY"),
        UserName: os.Getenv("TVDB_USER_NAME"),
    })
    
    if err != nil {
        return
    }

    data, err = client.SearchSeries(&tvdb.SearchSeriesRequest{
        Name: "<name_of_series>",
    })
```

## More examples
Refer to `example/example.go` or any `_test.go` file to see how the client is used.
