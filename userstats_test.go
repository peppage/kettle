package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISteamUserStatsServiceGetPlayerAchievements(t *testing.T) {
	const filePath = "./json/isteamuserstats/getplayerachievements.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetPlayerAchievements/v1/", func(w http.ResponseWriter, r *http.Request) {
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
	const filePath = "./json/isteamuserstats/getglobalachievementpercentagesforapp.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetGlobalAchievementPercentagesForApp/v2/", func(w http.ResponseWriter, r *http.Request) {
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
