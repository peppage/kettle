package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISteamUserServiceGetFriendList(t *testing.T) {
	const filePath = "./json/isteamuser/getfriendlist.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetFriendList/v1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":     "",
			"steamid": "76561198006575550",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	friends, _, err := client.ISteamUserService.GetFriendList(&GetFriendListParams{
		SteamID: 76561198006575550,
	})

	assert.Nil(t, err)
	assert.Len(t, friends, 5)

	assert.Equal(t, "76561197960412202", friends[0].SteamID)
	assert.Equal(t, "friend", friends[0].Relationship)
	assert.Equal(t, int64(1379557878), friends[0].FriendSince)
}

func TestISteamUserServiceResolveVanityURL(t *testing.T) {
	const filePath = "./json/isteamuser/resolvevanityurl.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/ResolveVanityURL/v1/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":       "",
			"vanityurl": "nelipot",
			"url_type":  "2",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	resp, _, err := client.ISteamUserService.ResolveVanityURL(&ResolveVanityURLParams{
		VanityURL: "nelipot",
		URLType:   Group,
	})

	assert.Nil(t, err)
	assert.Equal(t, "103582791431962114", resp.SteamID)
	assert.Equal(t, 1, resp.Success)
}

func TestISteamUserServiceGetPlayerSummaries(t *testing.T) {
	const filePath = "./json/isteamuser/getplayersummaries.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/GetPlayerSummaries/v2/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key":      "",
			"steamids": "76561197960435530,76561198006575550,76561197977122693",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	ids := []int64{
		76561197960435530,
		76561198006575550,
		76561197977122693,
	}
	summaries, _, err := client.ISteamUserService.GetPlayerSummaries(ids)

	assert.Nil(t, err)
	assert.Len(t, summaries, 3)

	assert.Equal(t, "76561197977122693", summaries[0].SteamID)
	assert.Equal(t, 3, summaries[0].CommunityVisibility)
	assert.Equal(t, 1, summaries[0].ProfileState)
	assert.Equal(t, "Viking", summaries[0].PersonaName)
	assert.Equal(t, int64(1483736691), summaries[0].LastLogoff)
	assert.Equal(t, "http://steamcommunity.com/id/Viking/", summaries[0].ProfileURL)
	assert.Equal(t, "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/8f/8fa7f1d5b92270783d6632a43cb0594f592839fa.jpg", summaries[0].Avatar)
	assert.Equal(t, "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/8f/8fa7f1d5b92270783d6632a43cb0594f592839fa_medium.jpg", summaries[0].AvatarMedium)
	assert.Equal(t, "https://steamcdn-a.akamaihd.net/steamcommunity/public/images/avatars/8f/8fa7f1d5b92270783d6632a43cb0594f592839fa_full.jpg", summaries[0].AvatarFull)
	assert.Equal(t, 1, summaries[0].PersonaState)
	assert.Equal(t, "Viking", summaries[0].RealName)
	assert.Equal(t, "103582791429523489", summaries[0].PrimaryClanID)
	assert.Equal(t, int64(1122128079), summaries[0].TimeCreated)
	assert.Equal(t, 0, summaries[0].PersonaStateFlags)
	assert.Equal(t, "HITMAN™", summaries[0].GameTitle)
	assert.Equal(t, "236870", summaries[0].GameID)
	assert.Equal(t, "SE", summaries[0].LocCountryCode)
}
