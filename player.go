package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// IPlayerService provides a method for accessing player information
type IPlayerService struct {
	sling *sling.Sling
}

func newIPlayerService(sling *sling.Sling) *IPlayerService {
	return &IPlayerService{
		sling: sling.Path("IPlayerService/"),
	}
}

type ownedGameResponse struct {
	OwnedGameResponse ogresp `json:"response"`
}

type ogresp struct {
	GameCount int        `json:"game_count"`
	Games     []UserGame `json:"games"`
}

// UserGame is a game listing specific to a user for IPlayerService.GetOwnedGames
type UserGame struct {
	AppID           int64  `json:"appid"`
	Name            string `json:"name,omitempty"`
	Playtime2Weeks  int    `json:"playtime_2weeks,omitempty"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconURL      string `json:"img_icon_url,omitempty"`
	ImgLogoURL      string `json:"img_logo_url,omitempty"`
	VisibleStats    bool   `json:"has_community_visible_stats,omitempty"`
}

// OwnedGamesParams are the parameters for IPlayerService.GetOwnedGames
type OwnedGamesParams struct {
	SteamID        string      `url:"steamid"`
	IncludeAppInfo BoolAsAnInt `url:"include_appinfo,omitempty"`
	IncludeFree    BoolAsAnInt `url:"include_played_free_games,omitempty"`
	//AppIdFilter too complicated right now
}

type playedGameResponse struct {
	PlayedGameResponse pgresp `json:"response"`
}

type pgresp struct {
	TotalCount int          `json:"total_count"`
	Games      []PlayedGame `json:"games"`
}

// PlayedGame is a struct for a game from IPlayerService.GetRecentlyPlayedGames
type PlayedGame struct {
	AppID           int64  `json:"appid"`
	Name            string `json:"name,omitempty"`
	Playtime2Weeks  int    `json:"playtime_2weeks,omitempty"`
	PlaytimeForever int    `json:"playtime_forever"`
	ImgIconURL      string `json:"img_icon_url,omitempty"`
	ImgLogoURL      string `json:"img_logo_url,omitempty"`
}

// GetOwnedGames returns a list of games a player owns along with some playtime information, if the profile is publicly visible.
// If IncludeAppInfo is not set in the params then the response will only have
// AppID, Playtime2Seeks, and PlaytimeForever
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetOwnedGames_.28v0001.29
// https://lab.xpaw.me/steam_api_documentation.html#IPlayerService_GetOwnedGames_v1
func (s IPlayerService) GetOwnedGames(params *OwnedGamesParams) ([]UserGame, *http.Response, error) {
	response := new(ownedGameResponse)

	resp, err := s.sling.New().Get("GetOwnedGames/v1/").QueryStruct(params).Receive(response, response)

	return response.OwnedGameResponse.Games, resp, err
}

// RecentGamesParams are the parameters for IPlayerService.GetRecentlyPlayedGames
type RecentGamesParams struct {
	SteamID string `url:"steamid"`
	Count   int    `url:"count,omitempty"`
}

// GetRecentlyPlayedGames returns a list of games a player has played in the last two weeks, if the profile is publicly visible.
// Private, friends-only, and other privacy settings are not supported unless you are asking for your own personal details
// (ie the WebAPI key you are using is linked to the steamid you are requesting).
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetRecentlyPlayedGames_.28v0001.29
func (s IPlayerService) GetRecentlyPlayedGames(params *RecentGamesParams) ([]PlayedGame, *http.Response, error) {
	response := new(playedGameResponse)

	resp, err := s.sling.New().Get("GetRecentlyPlayedGames/v0001/").QueryStruct(params).Receive(response, response)

	return response.PlayedGameResponse.Games, resp, err
}
