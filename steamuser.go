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
func (s *ISteamUserService) GetFriendList(params *GetFriendListParams) ([]Friend, *http.Response, error) {
	response := new(friendListResponse)

	resp, err := s.sling.New().Get("GetFriendList/v1/").QueryStruct(params).ReceiveSuccess(response)

	return response.FriendsList.Friends, resp, err
}
