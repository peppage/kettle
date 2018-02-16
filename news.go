package kettle

import (
	"net/http"

	"github.com/dghubble/sling"
)

// ISteamNewsService provides a method for accessing news about an app
type ISteamNewsService struct {
	sling *sling.Sling
}

func newISteamNewsService(sling *sling.Sling) *ISteamNewsService {
	return &ISteamNewsService{
		sling: sling.Path("ISteamNews/"),
	}
}

// GetNewsForAppParams are the paremeters for ISteamNewsService.GetNewsForApp
type GetNewsForAppParams struct {
	AppID     int64  `url:"appid"`
	MaxLength int    `url:"maxlength,omitempty"`
	EndDate   int64  `url:"enddate,omitempty"`
	Count     int    `url:"count,omitempty"`
	FeedName  string `url:"feedname,omitempty"`
}

type newsResponse struct {
	AppNews appNews `json:"appnews"`
}

type appNews struct {
	AppID     int64      `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
}

// NewsItem is news about an app from ISteamNewsService.GetNewsForApp
type NewsItem struct {
	GID         string `json:"gid"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ExternalURL bool   `json:"is_external_url"`
	Author      string `json:"author"`
	Contents    string `json:"contents"`
	FeedLabel   string `json:"feedlabel"`
	Date        int64  `json:"date"`
	FeedName    string `json:"feeds"`
}

// GetNewsForApp returns the latest of a game specified by its appID.
// https://developer.valvesoftware.com/wiki/Steam_Web_API#GetNewsForApp_.28v0002.29
// https://wiki.teamfortress.com/wiki/WebAPI/GetNewsForApp
func (s *ISteamNewsService) GetNewsForApp(params *GetNewsForAppParams) ([]NewsItem, *http.Response, error) {
	response := new(newsResponse)

	resp, err := s.sling.New().Get("GetNewsForApp/v2/").QueryStruct(params).ReceiveSuccess(response)

	return response.AppNews.NewsItems, resp, err
}
