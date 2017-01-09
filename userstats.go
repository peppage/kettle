package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ISteamUserStatsService provides access to information about a user's stats.
type ISteamUserStatsService struct {
	sling *sling.Sling
}

func newISteamUserStatsService(sling *sling.Sling) *ISteamUserStatsService {
	return &ISteamUserStatsService{
		sling: sling.Path("ISteamUserStats"),
	}
}

// GetPlayerAchievementsParams are the parameters for ISteamUserStatsService.GetPlayerAchievements
type GetPlayerAchievementsParams struct {
	SteamID int64  `url:"steamid"`
	AppID   int64  `url:"appid"`
	Lang    string `url:"l,omitempty"`
}

type playerAchievementsResp struct {
	PlayerStats PlayerStats `json:"playerstats"`
}

// PlayerStats are returned from ISteamUserStatsService.GetPlayerAchievements
type PlayerStats struct {
	SteamID      string        `json:"steamID"`
	GameName     string        `json:"gameName"`
	Achievements []Achievement `json:"achievements"`
	Success      bool          `json:"success"`
}

// Achievement is part of the PlayerStats returned from ISteamUserStatsService.GetPlayerAchievements
type Achievement struct {
	APIName  string `json:"apiname"`
	Achieved int    `json:"achieved"`
}

// GetPlayerAchievements Returns a list of achievements for this user by app id
// https://wiki.teamfortress.com/wiki/WebAPI/GetPlayerAchievements
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetPlayerAchievements_.28v0001.29
func (s *ISteamUserStatsService) GetPlayerAchievements(params *GetPlayerAchievementsParams) (*PlayerStats, *http.Response, error) {
	response := new(playerAchievementsResp)

	resp, err := s.sling.New().Get("GetPlayerAchievements/v1/").QueryStruct(params).ReceiveSuccess(response)

	return &response.PlayerStats, resp, err
}
