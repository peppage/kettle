package steam

import (
	"errors"
	"net/url"
	"strconv"

	log "github.com/Sirupsen/logrus"
)

type FriendListResponse struct {
	FriendsList FriendsList `json:"friendslist"`
}

type FriendsList struct {
	Friends []Friend `json:"friends"`
}

type Friend struct {
	SteamID      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int64  `json:"friend_since"`
}

func (api *Steam) GetFriendList(id int64, v url.Values) (list FriendListResponse, err error) {
	v = cleanValues(v)
	v.Set("steamid", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamUser/GetFriendList/v1",
		form:        v,
		data:        &list,
		response_ch: response_ch,
	}
	return list, (<-response_ch).err
}

type SummaryResponse struct {
	Response SResponse `json:"response"`
}

type SResponse struct {
	Players []Player `json:"players"`
}

type Player struct {
	SteamID             string `json:"steamid"`
	CommunityVisibility int    `json:"communityvisibilitystate"`
	ProfileState        int    `json:"provilestate"`
	PersonaName         string `json:"personaname"`
	LastLogoff          int64  `json:"lastlogoff"`
	ProfileURL          string `json:"profileurl"`
	Avatar              string `json:"avatar"`
	AvatarMedium        string `json:"avatarmedium"`
	AvatarFull          string `json:"avatarfull"`
	PersonaState        int    `json:"personastate"`
	PrimaryClanID       string `json:"primaryclandid"`
	TimeCreated         int64  `json:"timecreated"`
	PersonaStateFlags   int    `json:"personastateflags"`
	LocCountryCode      string `json:"loccountrycode"`
	LocStateCode        string `json:"locstatecode"`
}

func (api *Steam) GetPlayerSummaries(ids []int64) (users SummaryResponse, err error) {
	if len(ids) > 100 {
		log.WithFields(log.Fields{
			"len": len(ids),
		}).Error("Cannot have more than 100 ids")
		return users, errors.New("Over 100 ids")
	}
	v := url.Values{}
	var pids string
	for w, i := range ids {
		pids += strconv.FormatInt(i, 10)
		if w != len(ids)-1 {
			pids += ","
		}
	}
	v.Set("steamIds", pids)
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamUser/GetPlayerSummaries/v0002",
		form:        v,
		data:        &users,
		response_ch: response_ch,
	}
	return users, (<-response_ch).err
}

type VanityResponse struct {
	Response VResponse `json:"Response"`
}

type VResponse struct {
	SteamID string `json:"steamid"`
	Success int    `json:"success"`
}

func (api *Steam) ResolveVanityURL(vanityName string) (data VanityResponse, err error) {
	v := url.Values{}
	v.Set("vanityurl", vanityName)
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamUser/ResolveVanityURL/v0001",
		form:        v,
		data:        &data,
		response_ch: response_ch,
	}
	return data, (<-response_ch).err
}
