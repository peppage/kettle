package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISteamNewsServiceGetNewsForApp(t *testing.T) {
	const filePath = "./json/isteamnews/getnewsforapp.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetNewsForApp/v2/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":       "",
			"appid":     "440",
			"maxlength": "300",
			"count":     "3",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	news, _, err := client.ISteamNewsService.GetNewsForApp(&GetNewsForAppParams{
		AppID:     440,
		Count:     3,
		MaxLength: 300,
	})

	assert.Nil(t, err)
	assert.Len(t, news, 3)
	assert.Equal(t, "91593954435862753", news[0].GID)
	assert.Equal(t, "UGC League Winter 2017 Season is Starting!", news[0].Title)
	assert.Equal(t, "http://store.steampowered.com/news/externalpost/tf2_blog/91593954435862753", news[0].URL)
	assert.Equal(t, true, news[0].ExternalURL)
	assert.Equal(t, "", news[0].Author)
	assert.Equal(t, "<a href=\"http://www.ugcleague.com/\"> </a> Prepare yourself, the UGC League Winter Season is about to start! Join the thousands of teams that played last season for some competitive TF2 action! The first week of matches starts <b>January 23rd for Highlander, January 25th for 6v6 and January 27th for 4v4.</b> UGC has divisions across North America...", news[0].Contents)
	assert.Equal(t, "TF2 Blog", news[0].FeedLabel)
	assert.Equal(t, int64(1483470600), news[0].Date)
	assert.Equal(t, "tf2_blog", news[0].FeedName)
}
