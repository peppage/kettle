package kettle_test

import (
	"net/url"
	"os"
	"strconv"
	"testing"
	"time"

	"kettle"
	"kettle/steam"
	"kettle/store"
)

var KEY = os.Getenv("STEAM_KEY")

var api *steam.Steam
var storeApi *store.Store

const steamID = int64(76561198006575550)
const vanityName = "Peppage"

func init() {
	api = kettle.NewSteamApi(KEY)
	api.EnableThrottling(2*time.Second, 1)
	storeApi = kettle.NewStoreApi()
}

func Test_GetApps(t *testing.T) {
	games, err := api.GetAppList()
	if err != nil {
		t.Errorf("Getting all the games failed: %s", err.Error())
	}

	totalNewsItems := len(games.Applist.Apps)
	if totalNewsItems < 100 {
		t.Fatalf("Expected more than 100 apps, found %d", totalNewsItems)
	}
}

func Test_GetNews(t *testing.T) {
	news, err := api.GetNewsForApp(242760, url.Values{})
	if err != nil {
		t.Errorf("Getting news failed: %s", err.Error())
	}

	totalNewsItems := len(news.AppNews.NewsItems)
	if totalNewsItems < 2 {
		t.Fatalf("Expecting more than 2 news items, found %d", totalNewsItems)
	}
}

func Test_GetFriends(t *testing.T) {
	friends, err := api.GetFriendList(steamID, url.Values{})
	if err != nil {
		t.Errorf("Getting friends list failed: %s", err.Error())
	}

	totalFriends := len(friends.FriendsList.Friends)
	if totalFriends < 2 {
		t.Fatalf("Expecting more than 2 friends, found %d", totalFriends)
	}
}

func Test_PlayerSummaries(t *testing.T) {
	summary, err := api.GetPlayerSummaries([]int64{steamID})
	if err != nil {
		t.Errorf("Getting player summary failed: %s", err.Error())
	}

	if summary.Players[0].SteamID != strconv.FormatInt(steamID, 10) {
		t.Fatalf("Expecting steam IDs to match, got %s", summary.Players[0].SteamID)
	}
}

func Test_PlayerVanity(t *testing.T) {
	vanityResp, err := api.ResolveVanityURL(vanityName)
	if err != nil {
		t.Errorf("Getting player vanity failed: %s", err.Error())
	}

	if vanityResp.SteamID != strconv.FormatInt(steamID, 10) {
		t.Fatalf("Expecting my Steam ID, got %s", vanityResp.SteamID)
	}
}

func Test_GetOwnedGames(t *testing.T) {
	ownedResp, err := api.GetOwnedGames(steamID, url.Values{})
	if err != nil {
		t.Errorf("Getting owned games failed: %s", err.Error())
	}

	if ownedResp.GameCount < 263 {
		t.Fatalf("Expected owned games larger or equal to 263, got %d", ownedResp.GameCount)
	}
}

func Test_AppDetails(t *testing.T) {
	details, err := storeApi.GetAppDetails(49520, url.Values{})
	if err != nil {
		t.Errorf("Getting app data failed: %s", err.Error())
	}

	if !details["49520"].Success {
		t.Fatal("Expected a successful hit")
	}

	if details["49520"].Name != "Borderlands 2" {
		t.Fatalf("Expected title Borderlands 2 got, %s", details["49520"].Name)
	}
}
