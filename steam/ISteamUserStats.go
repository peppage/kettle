package steam

import (
	"net/url"
	"strconv"
)

type PlayerAchievementsResp struct {
	PlayerStats PlayerStats
	Success     bool
}

type PlayerStats struct {
	SteamID      string
	GameName     string
	Achievements []Achievement
}

type Achievement struct {
	ApiName  string
	Achieved int
}

func (api *Steam) GetPlayerAchievements(id int64, v url.Values) (resp PlayerAchievementsResp, err error) {
	v = cleanValues(v)
	v.Set("steamid", strconv.FormatInt(id, 10))

	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamUserStats/GetPlayerAchievements/v1",
		form:        v,
		data:        &resp,
		response_ch: response_ch,
	}
	return resp, (<-response_ch).err
}

type GameAchievementResp struct {
	AchievementPercentages AchievementPercentages
}

type AchievementPercentages struct {
	Achievements []GameAchievement
}

type GameAchievement struct {
	Name       string
	Percentage float64
}

func (api *Steam) GetGlobalAchievementPercentagesForApp(id int64) (resp GameAchievementResp, err error) {
	v := url.Values{}
	v.Set("gameid", strconv.FormatInt(id, 10))

	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamUserStats/GetGlobalAchievementPercentagesForApp/v0002",
		form:        v,
		data:        &resp,
		response_ch: response_ch,
	}
	return resp, (<-response_ch).err
}
