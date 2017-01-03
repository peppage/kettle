package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ISteamUserService provides access information about a steam user
type ISteamUserService struct {
	sling *sling.Sling
}

func newISteamUserService(sling *sling.Sling) *ISteamUserService {
	return &ISteamUserService{
		sling: sling.Path("ISteamUser"),
	}
}

type friendListResponse struct {
	FriendsList friendsList `json:"friendslist"`
}

type friendsList struct {
	Friends []Friend `json:"friends"`
}

// Friend is a steam users's friend.
type Friend struct {
	SteamID      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int64  `json:"friend_since"`
}

// GetFriendListParams are the parameters for ISteamUserService.GetFriendList
// Relatiionship (optional) can be "friend" or "all".
type GetFriendListParams struct {
	SteamID      int64  `url:"steamid"`
	Relationship string `url:"relationship,omitempty"`
}

// GetFriendList Returns friends Steam user if profile is public.
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetFriendList_.28v0001.29
// https://wiki.teamfortress.com/wiki/WebAPI/GetFriendList
func (s *ISteamUserService) GetFriendList(params *GetFriendListParams) ([]Friend, *http.Response, error) {
	response := new(friendListResponse)

	resp, err := s.sling.New().Get("GetFriendList/v1/").QueryStruct(params).ReceiveSuccess(response)

	return response.FriendsList.Friends, resp, err
}

type vanityResponse struct {
	VanityResponse VanityResponse `json:"Response"`
}

// VanityResponse is the response for ISteamUserService.ResolveVanityURL
type VanityResponse struct {
	SteamID string `json:"steamid"`
	Success int    `json:"success"`
	Message string `json:"message"`
}

// ResolveVanityURLParams the parameters for ISteamUserService.ResolveVanityURL
type ResolveVanityURLParams struct {
	VanityURL string     `url:"vanityurl"`
	URLType   VanityType `url:"url_type,omitempty"`
}

//VanityType is the type of vanity url you're trying to resolve
type VanityType int

// The options for VanityType
const (
	Individual        = VanityType(1)
	Group             = VanityType(2)
	OfficialGameGroup = VanityType(3)
)

// ResolveVanityURL resolve a vanity url to a steam user id.
// https://wiki.teamfortress.com/wiki/WebAPI/ResolveVanityURL
func (s *ISteamUserService) ResolveVanityURL(params *ResolveVanityURLParams) (*VanityResponse, *http.Response, error) {
	response := new(vanityResponse)

	resp, err := s.sling.New().Get("ResolveVanityURL/v1/").QueryStruct(params).ReceiveSuccess(response)

	return &response.VanityResponse, resp, err
}
