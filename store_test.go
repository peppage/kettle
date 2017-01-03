package kettle

import (
	"net/http"
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestStoreServiceAppDetails(t *testing.T) {
	const filePath = "./json/store/appdetails.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("store.steampowered.com/api/appdetails", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{"appids": "289070"}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	game, _, err := client.Store.AppDetails(289070)
	assert.Nil(t, err)

	assert.Equal(t, "game", game.Type)
	assert.Equal(t, "Sid Meier’s Civilization® VI", game.Name)
	assert.Equal(t, int64(289070), game.SteamAppID)
	assert.Equal(t, false, game.IsFree)

	v, ok := game.RequiredAge.(float64)
	assert.True(t, ok)
	assert.Equal(t, 18, int(v))

	assert.Equal(t, []int64{512033, 512032}, game.Dlc)
	assert.Equal(t, "<h1>D", game.DetailedDescription[0:5])
	assert.Equal(t, "Sid M", game.AboutTheGame[0:5])
	assert.Equal(t, "Civil", game.ShortDescription[0:5])
	assert.Equal(t, "Engli", game.SupportedLanguages[0:5])
	assert.Equal(t, "“I’ll ", game.Reviews[0:10])
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/289070/header.jpg?t=1482279729", game.HeaderImage)
	assert.Equal(t, "http://www.civilization.com/", game.Website)

	pcReq := new(Requirements)
	err = json.Unmarshal(game.PCRequirements, pcReq)
	assert.Nil(t, err)
	assert.Equal(t, "<strong>Minimum:", pcReq.Minimum[0:16])
	assert.Equal(t, "<strong>Recommended:", pcReq.Recommended[0:20])

	macReq := new(Requirements)
	err = json.Unmarshal(game.MacRequirements, macReq)
	assert.Nil(t, err)
	assert.Equal(t, "<strong>Minimum:", macReq.Minimum[0:16])
	assert.Equal(t, "", macReq.Recommended)

	linReq := new(Requirements)
	err = json.Unmarshal(game.LinuxRequirements, linReq)
	assert.NotNil(t, err)

	assert.Equal(t, "©2016 Tak", game.LegalNotice[0:10])

	assert.Len(t, game.Publishers, 2)
	assert.Equal(t, "2K", game.Publishers[0])

	assert.Equal(t, "USD", game.PriceOverview.Currency)
	assert.Equal(t, 5999, game.PriceOverview.Initial)
	assert.Equal(t, 5999, game.PriceOverview.Final)
	assert.Equal(t, 0, game.PriceOverview.DiscountPercent)

	assert.Len(t, game.Packages, 2)
	assert.Equal(t, 123215, game.Packages[0])

	assert.Len(t, game.PackageGroups, 1)
	assert.Equal(t, "default", game.PackageGroups[0].Name)
	assert.Equal(t, "Buy Sid Meier’s Civilization® VI", game.PackageGroups[0].Title)
	assert.Equal(t, "", game.PackageGroups[0].Description)
	assert.Equal(t, "Select a purchase option", game.PackageGroups[0].SelectionText)
	assert.Equal(t, "", game.PackageGroups[0].SaveText)

	v, ok = game.PackageGroups[0].DisplayType.(float64)
	assert.True(t, ok)
	assert.Equal(t, 12, int(v))

	assert.Equal(t, "false", game.PackageGroups[0].IsRecurringSubscription)

	assert.Equal(t, int64(123215), game.PackageGroups[0].Subs[0].PackageID)
	assert.Equal(t, "", game.PackageGroups[0].Subs[0].PercentSavingsText)
	assert.Equal(t, 0, game.PackageGroups[0].Subs[0].PercentSavings)
	assert.Equal(t, "Sid Meier's Civilization VI - $59.99", game.PackageGroups[0].Subs[0].OptionText)
	assert.Equal(t, "", game.PackageGroups[0].Subs[0].OptionDescription)
	assert.Equal(t, "0", game.PackageGroups[0].Subs[0].CanGetFreeLicense)
	assert.Equal(t, false, game.PackageGroups[0].Subs[0].IsFreeLicense)
	assert.Equal(t, 5999, game.PackageGroups[0].Subs[0].PriceInCentsWithDiscount)

	assert.Equal(t, true, game.Platforms.Windows)
	assert.Equal(t, true, game.Platforms.Mac)
	assert.Equal(t, false, game.Platforms.Linux)

	assert.Equal(t, 88, game.MetaCritic.Score)
	assert.Equal(t, "http://www.metacritic.com/game/pc/sid-meiers-civilization-vi?ftag=MCD-06-10aaa1f", game.MetaCritic.URL)

	assert.Len(t, game.Categories, 4)
	assert.Equal(t, 2, game.Categories[0].ID)
	assert.Equal(t, "Single-player", game.Categories[0].Description)

	assert.Len(t, game.Genres, 1)
	assert.Equal(t, "2", game.Genres[0].ID)
	assert.Equal(t, "Strategy", game.Genres[0].Description)

	assert.Len(t, game.Screenshots, 6)
	assert.Equal(t, 0, game.Screenshots[0].ID)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/289070/ss_36c63ebeb006b246cb740fdafeb41bb20e3b330d.600x338.jpg?t=1482279729", game.Screenshots[0].PathThumbnail)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/289070/ss_36c63ebeb006b246cb740fdafeb41bb20e3b330d.1920x1080.jpg?t=1482279729", game.Screenshots[0].PathFull)

	assert.Len(t, game.Movies, 3)
	assert.Equal(t, 256672694, game.Movies[0].ID)
	assert.Equal(t, "Civilization VI Launch Trailer - ESRB", game.Movies[0].Name)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/256672694/movie.293x165.jpg?t=1476736935", game.Movies[0].Thumbnail)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/256672694/movie480.webm?t=1476736935", game.Movies[0].Webm.Low)
	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/256672694/movie_max.webm?t=1476736935", game.Movies[0].Webm.Max)
	assert.Equal(t, true, game.Movies[0].Highlight)

	assert.Equal(t, int64(23249), game.Recomendations.Total)
	assert.Equal(t, 0, game.Achievements.Total)
	assert.Equal(t, false, game.ReleaseDate.ComingSoon)
	assert.Equal(t, "Oct 20, 2016", game.ReleaseDate.Date)

	assert.Equal(t, "http://support.2k.com", game.SupportInfo.URL)
	assert.Equal(t, "", game.SupportInfo.Email)

	assert.Equal(t, "http://cdn.akamai.steamstatic.com/steam/apps/289070/page_bg_generated_v6b.jpg?t=1482279729", game.Background)
}
