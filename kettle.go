package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Steam client for making Steam API requests
type Client struct {
	sling *sling.Sling

	Store                  *StoreService
	IPlayerService         *IPlayerService
	ISteamAppsService      *ISteamAppsService
	ISteamNewsService      *ISteamNewsService
	ISteamUserService      *ISteamUserService
	ISteamUserStatsService *ISteamUserStatsService
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, key string) *Client {
	b := sling.New().Client(httpClient)

	apiBase := b.New().Base("https://api.steampowered.com/")
	apiBase.QueryStruct(struct {
		Key string `url:"key"`
	}{
		Key: key,
	})

	return &Client{
		sling:                  b,
		Store:                  newStoreService(b.New().Base("https://store.steampowered.com/api/")),
		IPlayerService:         newIPlayerService(apiBase.New()),
		ISteamAppsService:      newISteamAppsService(apiBase.New()),
		ISteamNewsService:      newISteamNewsService(apiBase.New()),
		ISteamUserService:      newISteamUserService(apiBase.New()),
		ISteamUserStatsService: newISteamUserStatsService(apiBase.New()),
	}
}

// BoolAsAnInt is a bool that needs to be an int when transferred to an endpoint
type BoolAsAnInt int

// Option available for BoolAsAnInt
const (
	False = BoolAsAnInt(0)
	True  = BoolAsAnInt(1)
)
