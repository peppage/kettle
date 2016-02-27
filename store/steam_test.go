package store_test

import (
	"net/url"
	"testing"

	"kettle/store"
)

var storeAPI *store.Store

func init() {
	storeAPI = store.New()
}

func Test_AppDetails(t *testing.T) {
	details, err := storeAPI.GetAppDetails(49520, url.Values{})
	if err != nil {
		t.Errorf("Getting app data failed: %s", err.Error())
	}

	if !details["49520"].Success {
		t.Errorf("Expected a successful hit")
	}

	if details["49520"].Name != "Borderlands 2" {
		t.Errorf("Expected title Borderlands 2 got, %s", details["49520"].Name)
	}
}
