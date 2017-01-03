package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ISteamAppsService provides a method for accessing steam app info
type ISteamAppsService struct {
	sling *sling.Sling
}

func newISteamAppsService(sling *sling.Sling) *ISteamAppsService {
	return &ISteamAppsService{
		sling: sling.Path("ISteamApps"),
	}
}

type appListResponse struct {
	AppList appList `json:"applist"`
}

type appList struct {
	Apps []App `json:"apps"`
}

// App is an app on steam from ISteamAppsService.GetAppList
type App struct {
	AppID int64  `json:"appid"`
	Name  string `json:"name"`
}

// GetAppList returns a list of all apps on steam
// https://wiki.teamfortress.com/wiki/WebAPI/GetAppList
func (s *ISteamAppsService) GetAppList() ([]App, *http.Response, error) {
	response := new(appListResponse)

	resp, err := s.sling.New().Get("GetAppList/v2/").Receive(response, response)

	return response.AppList.Apps, resp, err
}
