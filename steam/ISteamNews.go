package steam

import (
	"net/url"
	"strconv"
)

type NewsResponse struct {
	AppNews `json:"appnews"`
}

type AppNews struct {
	AppID     int64      `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
}

type NewsItem struct {
	GID         string `json:"gid"`
	Title       string `json:"title"`
	URL         string `json:"url"`
	ExternalUrl bool   `json:"is_external_url"`
	Author      string `json:"author"`
	Contents    string `json:"contents"`
	FeedLabel   string `json:"feedlabel"`
	Date        int    `json:"date"`
	FeedName    string `json:"feedname"`
}

func (api *Steam) GetNewsForApp(id int64, v url.Values) (news NewsResponse, err error) {
	v = cleanValues(v)
	v.Set("appid", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamNews/GetNewsForApp/v0002",
		form:        v,
		data:        &news,
		response_ch: response_ch,
	}
	return news, (<-response_ch).err
}
