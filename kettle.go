package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// Client is a Steam client for making Steam API requests
type Client struct {
	sling *sling.Sling

	Store *StoreService
}

// NewClient returns a new Client
func NewClient(httpClient *http.Client, key string) *Client {
	b := sling.New().Client(httpClient)
	return &Client{
		sling: b,
		Store: newStoreService(b.New().Base("http://store.steampowered.com/api/")),
	}
}
