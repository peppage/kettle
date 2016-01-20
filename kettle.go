package kettle

import (
	"kettle/steam"
	"kettle/store"
)

func NewSteamApi(key string) *steam.Steam {
	return steam.New(key)
}

func NewStoreApi() *store.Store {
	return store.New()
}
