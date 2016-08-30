package steam

import (
	"net/url"
	"strconv"
)

type OwnedGameResponse struct {
	OGResponse `json:"response"`
}

type OGResponse struct {
	GameCount int        `json:"game_count"`
	Games     []UserGame `json:"games"`
}

type UserGame struct {
	AppID           int64  `json:"appid"`
	Name            string `json:"name,omitempty"`
	Playtime2Weeks  int32  `json:"playtime_2weeks"`
	PlaytimeForever int32  `json:"playtime_forever"`
	ImgIconUrl      string `json:"img_icon_url,omitempty"`
	ImgLogoUrl      string `json:"img_logo_url,omitempty"`
	VisibleStats    bool   `json:"has_community_visible_stats,omitempty"`
}

// https://wiki.teamfortress.com/wiki/WebAPI/GetOwnedGames
func (api *Steam) GetOwnedGames(id int64, v url.Values) (data OwnedGameResponse, err error) {
	v = cleanValues(v)
	v.Set("steamid", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/IPlayerService/GetOwnedGames/v1",
		form:        v,
		data:        &data,
		response_ch: response_ch,
	}
	return data, (<-response_ch).err
}
