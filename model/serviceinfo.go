package model

import "time"

type ServiceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Time      time.Time `json:"time"`
		Transport string    `json:"transport"`
		Service   struct {
			HTTP struct {
				HTMLHash string `json:"html_hash"`
				Favicon  struct {
					Hash     string `json:"hash"`
					Location string `json:"location"`
					Data     string `json:"data"`
				} `json:"favicon"`
				Robots          string `json:"robots"`
				SitemapHash     string `json:"sitemap_hash"`
				Server          string `json:"server"`
				Body            string `json:"body"`
				XPoweredBy      string `json:"x_powered_by"`
				MetaKeywords    string `json:"meta_keywords"`
				RobotsHash      string `json:"robots_hash"`
				Sitemap         string `json:"sitemap"`
				Path            string `json:"path"`
				Title           string `json:"title"`
				Host            string `json:"host"`
				SecurityText    string `json:"security_text"`
				StatusCode      int    `json:"status_code"`
				ResponseHeaders string `json:"response_headers"`
			} `json:"http"`
			Version  string `json:"version"`
			Name     string `json:"name"`
			Product  string `json:"product"`
			Banner   string `json:"banner"`
			Response string `json:"response"`
		} `json:"service"`
		Images     []interface{} `json:"images"`
		OsName     string        `json:"os_name"`
		Components []interface{} `json:"components"`
		Location   struct {
			DistrictCn  string    `json:"district_cn"`
			ProvinceCn  string    `json:"province_cn"`
			Gps         []float64 `json:"gps"`
			ProvinceEn  string    `json:"province_en"`
			CityEn      string    `json:"city_en"`
			CountryCode string    `json:"country_code"`
			CountryEn   string    `json:"country_en"`
			Radius      float64   `json:"radius"`
			DistrictEn  string    `json:"district_en"`
			Isp         string    `json:"isp"`
			StreetEn    string    `json:"street_en"`
			Owner       string    `json:"owner"`
			CityCn      string    `json:"city_cn"`
			CountryCn   string    `json:"country_cn"`
			StreetCn    string    `json:"street_cn"`
		} `json:"location"`
		Asn       int    `json:"asn"`
		Hostname  string `json:"hostname"`
		Org       string `json:"org"`
		OsVersion string `json:"os_version"`
		IsIpv6    bool   `json:"is_ipv6"`
		IP        string `json:"ip"`
		Port      int    `json:"port"`
	} `json:"data"`
	Meta struct {
		Total        int    `json:"total"`
		PaginationID string `json:"pagination_id"`
	} `json:"meta"`
}
