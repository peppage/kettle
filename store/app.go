package store

import (
	"net/url"
	"strconv"
)

type AppDetailsResponse map[string]AppDetails

type AppDetails struct {
	Success bool `json:"success"`
	AppData `json:"data"`
}

type GenreID string

const (
	EarlyAccess = GenreID("70")
)

func (details AppDetails) HasGenre(id GenreID) bool {
	for _, g := range details.Genres {
		if g.ID == string(id) {
			return true
		}
	}
	return false
}

type AppData struct {
	Type                string         `json:"type"`
	Name                string         `json:"name"`
	SteamAppId          int64          `json:"steam_appid"`
	IsFree              bool           `json:"is_free"`
	ControllerSupport   string         `json:"controller_support"`
	Dlc                 []int64        `json:"dlc"`
	DetailedDescription string         `json:"detailed_description"`
	AboutTheGame        string         `json:"about_the_game"`
	SupportedLanguages  string         `json:"supported_languages"`
	Reviews             string         `json:"reviews"`
	HeaderImage         string         `json:"header_image"`
	Website             string         `json:"website"`
	LegalNotice         string         `json:"legal_notice"`
	Developers          []string       `json:"developers"`
	Publishers          []string       `json:"publishers"`
	PriceOverview       Price          `json:"price_overview"`
	Packages            []interface{}  `json:"packages"`
	PackageGroups       []PackageGroup `json:"package_groups"`
	Platforms           Platform       `json:"platforms"`
	MetaCritic          `json:"metacritic,omitempty"`
	Categories          []Category   `json:"categories"`
	Genres              []Genre      `json:"genres"`
	Screenshots         []Screenshot `json:"screenshots"`
	Movies              []Movie      `json:"movies"`
	Recomendations      `json:"recommendations"`
	Achievements        `json:"achievements"`
	ReleaseDate         `json:"release_date"`
	SupportInfo         `json:"support_info"`
	Background          string `json:"background"`
	//PcRequirements      Requirements   `json:"pc_requirements"` //This can be empty for demos
	//RequredAge          string         `json:"required_age"` // Can be string or int
	//MacRequirements     Requirements   `json:"mac_requirements,omitempty"`
	//LinuxRequirements   Requirements   `json:"linux_requirements,omitempty"`
	// These 2 fields are sometimes arrays and sometimes values :/
}

type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

type Price struct {
	Currency        string `json:"currency"`
	Initial         int    `json:"initial"`
	Final           int    `json:"final"`
	DiscountPercent int    `json:"discount_percent"`
}

type PackageGroup struct {
	Name                    string `json:"name"`
	Title                   string `json:"title"`
	Description             string `json:"description"`
	SelectionText           string `json:"selection_text"`
	SaveText                string `json:"save_text"`
	IsRecurringSubscription string `json:"is_recurring_subscription"`
	Subs                    []Sub  `json:"subs"`
	//DisplayType             int    `json:"display_type"` This can be a string or number
}

type Sub struct {
	PackageID                interface{} `json:"packageid"`
	PercentSavingsText       string      `json:"percent_savings_text"`
	PercentSavings           int         `json:"percent_savings"`
	OptionText               string      `json:"option_text"`
	OptionDescription        string      `json:"option_description"`
	CanGetFreeLicense        string      `json:"can_get_free_license"`
	IsFreeLicense            bool        `json:"is_free_license"`
	PriceInCentsWithDiscount int         `json:"price_in_cents_with_discount"`
}

type Platform struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

type MetaCritic struct {
	Score int    `json:"score"`
	URL   string `json:"url"`
}

type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type Screenshot struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

type Movie struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Thumbnail  string `json:"thumbnail"`
	Webm       `json:"webm"`
	Hightlight bool `json:"highlight"`
}

type Webm struct {
	Low string `json:"480"`
	Max string `json:"max"`
}

type Recomendations struct {
	Total int64 `json:"total"`
}

type Achievements struct {
	Total       int                     `json:"total"`
	Highlighted []HighlightedAchivement `json:"highlighted"`
}

type HighlightedAchivement struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

func (api *Store) GetAppDetails(id int64, v url.Values) (data AppDetailsResponse, err error) {
	v = cleanValues(v)
	v.Set("appids", strconv.FormatInt(id, 10))
	response_ch := make(chan response)
	api.queryQueue <- query{
		url:         api.baseUrl + "/appdetails",
		form:        v,
		data:        &data,
		response_ch: response_ch,
	}
	return data, (<-response_ch).err
}
