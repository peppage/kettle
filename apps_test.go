package kettle

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestISteamAppsServicegGetAppList(t *testing.T) {
	t.Parallel()
	const filePath = "./json/isteamappservice/getapplist.json"
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/ISteamApps/GetAppList/v2/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)

		assertQuery(t, map[string]string{
			"key": "",
		}, r)

		b, err := getTestFile(filePath)
		if err != nil {
			t.Fatalf("Failed to open testfile %s", filePath)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	})

	client := NewClient(httpClient, "")
	apps, _, err := client.ISteamAppsService.GetAppList()

	assert.Nil(t, err)
	assert.Len(t, apps, 5)
	assert.Equal(t, int64(5), apps[0].AppID)
	assert.Equal(t, "Dedicated Server", apps[0].Name)
}
