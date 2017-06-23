package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISteamUserStatsServiceGetPlayerAchievements(t *testing.T) {
	t.Parallel()
	const filePath = "./json/isteamuserstats/getplayerachievements.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/ISteamUserStats/GetPlayerAchievements/v1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":     "",
			"steamid": "76561198006575550",
			"appid":   "98800",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	resp, _, err := client.ISteamUserStatsService.GetPlayerAchievements(&GetPlayerAchievementsParams{
		SteamID: 76561198006575550,
		AppID:   98800,
	})

	assert.Nil(t, err)
	assert.Equal(t, "76561198006575550", resp.SteamID)
	assert.Equal(t, "Dungeons of Dredmor", resp.GameName)

	assert.Len(t, resp.Achievements, 5)
	assert.Equal(t, "ACHIEVEMENT_KILL_DREDMOR_EASY", resp.Achievements[0].APIName)
	assert.Equal(t, 1, resp.Achievements[0].Achieved)

	assert.True(t, resp.Success)
}

func TestISteamUserStatsServiceGetGlobalAchievementPercentagesForApp(t *testing.T) {
	t.Parallel()
	const filePath = "./json/isteamuserstats/getglobalachievementpercentagesforapp.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v2/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":    "",
			"gameid": "98800",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	chives, _, err := client.ISteamUserStatsService.GetGlobalAchievementPercentagesForApp(98800)

	assert.Nil(t, err)
	assert.Len(t, chives, 5)
	assert.Equal(t, "ACHIEVEMENT_ORNITHOLOGIST", chives[0].Name)
	assert.Equal(t, float64(66.185623168945313), chives[0].Percent)
}

func TestISteamUserStatsServiceGetSchemaForGame(t *testing.T) {
	t.Parallel()
	const filePath = "./json/isteamuserstats/getschemaforgame.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/ISteamUserStats/GetSchemaForGame/v2/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":   "",
			"appid": "98800",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	gameSchema, _, err := client.ISteamUserStatsService.GetSchemaForGame(98800)

	assert.Nil(t, err)
	assert.Equal(t, "Dungeons of Dredmor", gameSchema.GameName)
	assert.Equal(t, "21", gameSchema.GameVersion)

	assert.Len(t, gameSchema.AvailableGameStats.Achievements, 5)
	assert.Equal(t, "ACHIEVEMENT_KILL_DREDMOR_EASY", gameSchema.AvailableGameStats.Achievements[0].Name)
	assert.Equal(t, 0, gameSchema.AvailableGameStats.Achievements[0].DefaultValue)
	assert.Equal(t, "Dread Less", gameSchema.AvailableGameStats.Achievements[0].DisplayName)
	assert.Equal(t, 0, gameSchema.AvailableGameStats.Achievements[0].Hidden)
	assert.Equal(t, "Kill Lord Dredmor on Elvishly Easy Mode.", gameSchema.AvailableGameStats.Achievements[0].Description)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steamcommunity/public/images/apps/98800/172d00c65ca72db30f9fa38727f7d8c8b70c1bd0.jpg", gameSchema.AvailableGameStats.Achievements[0].Icon)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steamcommunity/public/images/apps/98800/ad5f99ddc91e7b059a0c94f2a80a0309c07f2c3c.jpg", gameSchema.AvailableGameStats.Achievements[0].IconGray)

	assert.Len(t, gameSchema.AvailableGameStats.Stats, 5)
	assert.Equal(t, "STAT_VICTORIES", gameSchema.AvailableGameStats.Stats[0].Name)
	assert.Equal(t, 0, gameSchema.AvailableGameStats.Stats[0].DefaultValue)
	assert.Equal(t, "Victories", gameSchema.AvailableGameStats.Stats[0].DisplayName)
}
