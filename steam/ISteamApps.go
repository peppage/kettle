package steam

import (
	"net/url"
)

type AppListResponse struct {
	Applist `json:"applist"`
}

type Applist struct {
	Apps []App `json:"apps"`
}

type App struct {
	AppID int64  `json:"appid"`
	Name  string `json:"name"`
}

func (api *Steam) GetAppList() (applist AppListResponse, err error) {
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/ISteamApps/GetAppList/v2",
		form:        url.Values{},
		data:        &applist,
		response_ch: response_ch,
	}
	return applist, (<-response_ch).err
}
