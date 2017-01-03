package kettle

import (
	"github.com/dghubble/sling"
)

// IPlayerService provides a method for accessing player information
type IPlayerService struct {
	sling *sling.Sling
}

func newIPlayerService(sling *sling.Sling) *IPlayerService {
	return &IPlayerService{
		sling: sling.Path("IPlayerService"),
	}
}

// GetOwnedGames returns a list of games a player owns along with some playtime information, if the profile is publicly visible.
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetOwnedGames_.28v0001.29
//func (s IPlayerService) GetOwnedGames()
