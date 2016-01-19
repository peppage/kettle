package kettle

import (
	"kettle/steam"
	"kettle/store"
)

func NewSteamApi(key string) *steam.Steam {
	return steam.New(key)
}

func NewStoreApi(key string) *store.Store {
	return store.New()
}
