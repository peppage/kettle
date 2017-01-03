package kettle

import (
	"errors"
	"net/http"
	"strconv"

	"encoding/json"

	"github.com/dghubble/sling"
)

// StoreService provies a method for accessing Steam store endpoints
type StoreService struct {
	sling *sling.Sling
}

func newStoreService(sling *sling.Sling) *StoreService {
	return &StoreService{
		sling: sling,
	}
}

type appDetails struct {
	Success bool    `json:"success"`
	AppData AppData `json:"data"`
}

// AppData holds the data for StoreService.AppDetails
// https://wiki.teamfortress.com/wiki/User:RJackson/StorefrontAPI#appdetails
type AppData struct {
	Type                string          `json:"type"`
	Name                string          `json:"name"`
	SteamAppID          int64           `json:"steam_appid"`
	RequiredAge         interface{}     `json:"required_age"` // Can be string or int
	IsFree              bool            `json:"is_free"`
	ControllerSupport   string          `json:"controller_support"`
	Dlc                 []int64         `json:"dlc"`
	DetailedDescription string          `json:"detailed_description"`
	AboutTheGame        string          `json:"about_the_game"`
	ShortDescription    string          `json:"short_description"`
	SupportedLanguages  string          `json:"supported_languages"`
	Reviews             string          `json:"reviews"`
	HeaderImage         string          `json:"header_image"`
	Website             string          `json:"website"`
	PCRequirements      json.RawMessage `json:"pc_requirements,omitempty"`
	MacRequirements     json.RawMessage `json:"mac_requirements,omitempty"`
	LinuxRequirements   json.RawMessage `json:"linux_requirements,omitempty"`
	LegalNotice         string          `json:"legal_notice"`
	Developers          []string        `json:"developers"`
	Publishers          []string        `json:"publishers"`
	PriceOverview       Price           `json:"price_overview"`
	Packages            []int           `json:"packages"`
	PackageGroups       []PackageGroup  `json:"package_groups"`
	Platforms           Platform        `json:"platforms"`
	MetaCritic          MetaCritic      `json:"metacritic,omitempty"`
	Categories          []Category      `json:"categories"`
	Genres              []Genre         `json:"genres"`
	Screenshots         []Screenshot    `json:"screenshots"`
	Movies              []Movie         `json:"movies"`
	Recomendations      Recomendations  `json:"recommendations"`
	Achievements        Achievements    `json:"achievements"`
	ReleaseDate         ReleaseDate     `json:"release_date"`
	SupportInfo         SupportInfo     `json:"support_info"`
	Background          string          `json:"background"`
}

// Requirements is possibly used for pc/mac/linux requirements. If it's empty
// it will be an empty array so you will need to cast this yourself.
type Requirements struct {
	Minimum     string `json:"minimum"`
	Recommended string `json:"recommended"`
}

// Price holds the current and sale price for an app
type Price struct {
	Currency        string `json:"currency"`
	Initial         int    `json:"initial"`
	Final           int    `json:"final"`
	DiscountPercent int    `json:"discount_percent"`
}

// PackageGroup are the packages the app is part of
type PackageGroup struct {
	Name                    string      `json:"name"`
	Title                   string      `json:"title"`
	Description             string      `json:"description"`
	SelectionText           string      `json:"selection_text"`
	SaveText                string      `json:"save_text"`
	DisplayType             interface{} `json:"display_type"` //This can be a string or number
	IsRecurringSubscription string      `json:"is_recurring_subscription"`
	Subs                    []Sub       `json:"subs"`
}

// Sub is part of the PackageGroup, details about a package
type Sub struct {
	PackageID                int64  `json:"packageid"`
	PercentSavingsText       string `json:"percent_savings_text"`
	PercentSavings           int    `json:"percent_savings"`
	OptionText               string `json:"option_text"`
	OptionDescription        string `json:"option_description"`
	CanGetFreeLicense        string `json:"can_get_free_license"`
	IsFreeLicense            bool   `json:"is_free_license"`
	PriceInCentsWithDiscount int    `json:"price_in_cents_with_discount"`
}

// Platform lists what platforms this app works on
type Platform struct {
	Windows bool `json:"windows"`
	Mac     bool `json:"mac"`
	Linux   bool `json:"linux"`
}

// MetaCritic information about the app
type MetaCritic struct {
	Score int    `json:"score"`
	URL   string `json:"url"`
}

// Category associated to an AppData
type Category struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

// Genre associated with an AppData
type Genre struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

// Screenshot is a screenshot of an app
type Screenshot struct {
	ID            int    `json:"id"`
	PathThumbnail string `json:"path_thumbnail"`
	PathFull      string `json:"path_full"`
}

// Movie are trailers and videos associated with AppData
type Movie struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Webm      Webm   `json:"webm"`
	Highlight bool   `json:"highlight"`
}

// Webm are links to the different quality level of a Movie
type Webm struct {
	Low string `json:"480"`
	Max string `json:"max"`
}

// Recomendations are how many people recommended the app
type Recomendations struct {
	Total int64 `json:"total"`
}

// Achievements are the achievements associated with AppData
type Achievements struct {
	Total       int                     `json:"total"`
	Highlighted []HighlightedAchivement `json:"highlighted"`
}

// HighlightedAchivement are name/photo of achievements on an App
type HighlightedAchivement struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// ReleaseDate is the release date for an app
type ReleaseDate struct {
	ComingSoon bool   `json:"coming_soon"`
	Date       string `json:"date"`
}

// SupportInfo holds contact info for support for an app
type SupportInfo struct {
	URL   string `json:"url"`
	Email string `json:"email"`
}

// AppDetails gets detailed information about a game
// Not supported: multiple appids with &filters=price_overview
// https://wiki.teamfortress.com/wiki/User:RJackson/StorefrontAPI#appdetails
func (s *StoreService) AppDetails(id int64) (*AppData, *http.Response, error) {
	response := make(map[string]appDetails)

	resp, err := s.sling.New().Path("appdetails").QueryStruct(struct {
		AppIDs int64 `url:"appids"`
	}{
		AppIDs: id,
	}).Receive(&response, &response)

	i := strconv.FormatInt(id, 10)
	a := response[i].AppData

	if !response[i].Success && err != nil {
		err = errors.New("API request failed with Success = false")
	}

	return &a, resp, err
}
