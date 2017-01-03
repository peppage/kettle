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
