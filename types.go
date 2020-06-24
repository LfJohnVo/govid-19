package main

import "time"

type Sum struct {
	Global struct {
		NewConfirmed   int `json:"NewConfirmed"`
		TotalConfirmed int `json:"TotalConfirmed"`
		NewDeaths      int `json:"NewDeaths"`
		TotalDeaths    int `json:"TotalDeaths"`
		NewRecovered   int `json:"NewRecovered"`
		TotalRecovered int `json:"TotalRecovered"`
	} `json:"Global"`
	Countries []struct {
		Country        string    `json:"Country"`
		CountryCode    string    `json:"CountryCode"`
		Slug           string    `json:"Slug"`
		NewConfirmed   int       `json:"NewConfirmed"`
		TotalConfirmed int       `json:"TotalConfirmed"`
		NewDeaths      int       `json:"NewDeaths"`
		TotalDeaths    int       `json:"TotalDeaths"`
		NewRecovered   int       `json:"NewRecovered"`
		TotalRecovered int       `json:"TotalRecovered"`
		Date           time.Time `json:"Date"`
	} `json:"Countries"`
	Date time.Time `json:"Date"`
}

// struct Countries
type Countries []struct {
	Country string `json:"Country"`
	Slug    string `json:"Slug"`
	ISO2    string `json:"ISO2"`
}

type Country []struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Province    string    `json:"Province"`
	City        string    `json:"City"`
	CityCode    string    `json:"CityCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Confirmed   int       `json:"Confirmed"`
	Deaths      int       `json:"Deaths"`
	Recovered   int       `json:"Recovered"`
	Active      int       `json:"Active"`
	Date        time.Time `json:"Date"`
}

type Alldatastruct []struct {
	Country     string    `json:"Country"`
	CountryCode string    `json:"CountryCode"`
	Lat         string    `json:"Lat"`
	Lon         string    `json:"Lon"`
	Confirmed   int       `json:"Confirmed"`
	Deaths      int       `json:"Deaths"`
	Recovered   int       `json:"Recovered"`
	Active      int       `json:"Active"`
	Date        time.Time `json:"Date"`
	LocationID  string    `json:"LocationID"`
}

// struct Stats
type Stats struct {
	Total                             int    `json:"Total"`
	All                               int    `json:"All"`
	AllUpdated                        string `json:"AllUpdated"`
	Countries                         int    `json:"Countries"`
	CountriesUpdated                  string `json:"CountriesUpdated"`
	ByCountry                         int    `json:"ByCountry"`
	ByCountryUpdated                  string `json:"ByCountryUpdated"`
	ByCountryLive                     int    `json:"ByCountryLive"`
	ByCountryLiveUpdated              string `json:"ByCountryLiveUpdated"`
	ByCountryTotal                    int    `json:"ByCountryTotal"`
	ByCountryTotalUpdated             string `json:"ByCountryTotalUpdated"`
	DayOne                            int    `json:"DayOne"`
	DayOneUpdated                     string `json:"DayOneUpdated"`
	DayOneLive                        int    `json:"DayOneLive"`
	DayOneLiveUpdated                 string `json:"DayOneLiveUpdated"`
	DayOneTotal                       int    `json:"DayOneTotal"`
	DayOneTotalUpdated                string `json:"DayOneTotalUpdated"`
	LiveCountryStatus                 int    `json:"LiveCountryStatus"`
	LiveCountryStatusUpdated          string `json:"LiveCountryStatusUpdated"`
	LiveCountryStatusAfterDate        int    `json:"LiveCountryStatusAfterDate"`
	LiveCountryStatusAfterDateUpdated string `json:"LiveCountryStatusAfterDateUpdated"`
	Stats                             int    `json:"Stats"`
	StatsUpdated                      string `json:"StatsUpdated"`
	Default                           int    `json:"Default"`
	DefaultUpdated                    string `json:"DefaultUpdated"`
	SubmitWebhook                     int    `json:"SubmitWebhook"`
	SubmitWebhookUpdated              string `json:"SubmitWebhookUpdated"`
	Summary                           int    `json:"Summary"`
	SummaryUpdated                    string `json:"SummaryUpdated"`
}
