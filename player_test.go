package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPlayerServiceGetOwnedGames(t *testing.T) {
	t.Parallel()
	const filePath = "./json/iplayerservice/ownedgames.complete.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/IPlayerService/GetOwnedGames/v1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"steamid":         "76561198006575550",
			"include_appinfo": "1",
			"key":             "",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	games, _, err := client.IPlayerService.GetOwnedGames(&OwnedGamesParams{
		SteamID:        "76561198006575550",
		IncludeAppInfo: True,
	})

	assert.Nil(t, err)
	assert.True(t, len(games) > 0)

	assert.Equal(t, int64(220), games[0].AppID)
	assert.Equal(t, "Half-Life 2", games[0].Name)
	assert.Equal(t, 391, games[0].Playtime2Weeks)
	assert.Equal(t, 1436, games[0].PlaytimeForever)
	assert.Equal(t, "fcfb366051782b8ebf2aa297f3b746395858cb62", games[0].ImgIconURL)
	assert.Equal(t, "e4ad9cf1b7dc8475c1118625daf9abd4bdcbcad0", games[0].ImgLogoURL)
	assert.Equal(t, true, games[0].VisibleStats)
}

func TestIplayerServiceGetRecentlyPlayed(t *testing.T) {
	t.Parallel()
	const filePath = "./json/iplayerservice/recentlyplayedgames.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/IPlayerService/GetRecentlyPlayedGames/v0001/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"steamid": "76561198006575550",
			"key":     "",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	games, _, err := client.IPlayerService.GetRecentlyPlayedGames(&RecentGamesParams{
		SteamID: "76561198006575550",
	})

	assert.Nil(t, err)
	assert.True(t, len(games) > 0)

	assert.Equal(t, int64(468670), games[0].AppID)
	assert.Equal(t, "Speed Brawl", games[0].Name)
	assert.Equal(t, 420, games[0].Playtime2Weeks)
	assert.Equal(t, 421, games[0].PlaytimeForever)
	assert.Equal(t, "8f48f2265746c08cd1fd7b8ce5f310172eb4fa12", games[0].ImgIconURL)
	assert.Equal(t, "ba0d065302833a4851093410ced0e1c82df30ba4", games[0].ImgLogoURL)
}
