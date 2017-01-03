package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIPlayerServiceGetOwnedGames(t *testing.T) {
	const filePath = "./json/iplayerservice/ownedgames.simple.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetOwnedGames/v1/", func(w http.ResponseWriter, r *http.Request) {
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
