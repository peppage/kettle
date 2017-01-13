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
		sling: sling.Path("ISteamUserStats/"),
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

type gameAchievementResp struct {
	AchievementPercentages achievementPercentages `json:"achievementpercentages"`
}

type achievementPercentages struct {
	Achievements []GameAchievement `json:"achievements"`
}

// GameAchievement is the response for ISteamUserStatsService.GetGlobalAchievementPercentagesForApp
type GameAchievement struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
}

// GetGlobalAchievementPercentagesForApp Statistics showing how much of the player base have unlocked various achievements.
// https://wiki.teamfortress.com/wiki/WebAPI/GetGlobalAchievementPercentagesForApp
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetGlobalAchievementPercentagesForApp_.28v0002.29
func (s *ISteamUserStatsService) GetGlobalAchievementPercentagesForApp(gameid int64) ([]GameAchievement, *http.Response, error) {
	response := new(gameAchievementResp)

	type params struct {
		GameID int64 `url:"gameid"`
	}

	p := &params{
		GameID: gameid,
	}

	resp, err := s.sling.New().Get("GetGlobalAchievementPercentagesForApp/v2/").QueryStruct(p).ReceiveSuccess(response)

	return response.AchievementPercentages.Achievements, resp, err
}

type schemaResp struct {
	Game GameSchema `json:"game"`
}

// GameSchema is the response for ISteamUserStatsService.GetSchemaForGame
type GameSchema struct {
	GameName           string `json:"gameName"`
	GameVersion        string `json:"gameVersion"`
	AvailableGameStats Stats  `json:"availableGameStats"`
}

// Stats are the listed achievements for a GameSchema
type Stats struct {
	Achievements []SchemaAchievement `json:"achievements"`
	Stats        []SchemaStat        `json:"stats"`
}

// SchemaAchievement is an achievement for Stats part of ISteamUserStatsService.GetSchemaForGame
type SchemaAchievement struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Description  string `json:"description"`
	Icon         string `json:"icon"`
	IconGray     string `json:"icongray"`
}

// SchemaStat is a stat for Stats part of ISteamUserStatsService.GetSchemaForGame
type SchemaStat struct {
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
}

// GetSchemaForGame returns gamename, gameversion and availablegamestats
// https://wiki.teamfortress.com/wiki/WebAPI/GetSchemaForGame
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetSchemaForGame_.28v2.29
func (s *ISteamUserStatsService) GetSchemaForGame(appid int64) (*GameSchema, *http.Response, error) {
	response := new(schemaResp)

	type params struct {
		AppID int64 `url:"appid"`
	}

	p := &params{
		AppID: appid,
	}

	resp, err := s.sling.New().Get("GetSchemaForGame/v2/").QueryStruct(p).ReceiveSuccess(response)

	return &response.Game, resp, err
}
