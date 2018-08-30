# Kettle [![GoDoc](https://godoc.org/github.com/peppage/kettle?status.svg)](https://godoc.org/github.com/peppage/kettle) [![Go Report Card](https://goreportcard.com/badge/github.com/peppage/kettle)](https://goreportcard.com/report/github.com/peppage/kettle)
Kettle is a simple Go package for accessing the[Steam API](https://developer.valvesoftware.com/wiki/Steam_Web_API), or [xPaw documentation](https://lab.xpaw.me/steam_api_documentation.html) and the [Storefront API](https://wiki.teamfortress.com/wiki/User:RJackson/StorefrontAPI).

If you find any errors please create an issue or I'll accept your pull request :)

## Install
    go get -u github.com/peppage/kettle

## Usage
```go

    httpClient := http.DefaultClient
    steamClient = kettle.NewClient(httpClient, "steamkey")
    
    // Get the full app list
    games, _, err := steamClient.ISteamAppsService.GetAppList()
    
    // Get reviews for a game
    d, _, err := steamClient.Store.AppReviews(&kettle.AppReviewsParams{
		AppID:    gameID,
		Language: "all",
	})
	
	// Get details about a game
	d, _, err := steamClient.Store.AppDetails(game.ID)
```

## License
[MIT License](LICENSE.md)
