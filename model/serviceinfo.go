package model

type ServiceInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    []struct {
		Service struct {
			Http struct {
				Favicon struct {
					Hash     string `json:"hash"`
					Data     string `json:"data"`
					Location string `json:"location"`
				} `json:"favicon"`
				SecurityText    string `json:"security_text"`
				RobotsText      string `json:"robots_hash"`
				ResponseHeaders string `json:"response_headers"`
				Path            string `json:"path"`
				SitemapHash     string `json:"sitemap_hash"`
				Body            string `json:"body"`
				Title           string `json:"title"`
				XPoweredBy      string `json:"x_powered_by"`
				Host            string `json:"host"`
				StatusCode      int    `json:"status_code"`
				Robots          string `json:"robots"`
				MetaKeywords    string `json:"meta_keywords"`
				HtmlHash        string `json:"html_hash"`
				Sitemap         string `json:"sitemap"`
				Server          string `json:"server"`
			} `json:"http"`
			Banner   string `json:"banner"`
			Response string `json:"response"`
			Name     string `json:"name"`
			Version  string `json:"version"`
			Product  string `json:"product"`
		} `json:"service"`
		Port     int `json:"port"`
		Location struct {
			Gps         []float64 `json:"gps"`
			CountryEn   string    `json:"country_en"`
			CountryCode string    `json:"country_code"`
			ProvinceEn  string    `json:"province_en"`
			DistrictEn  string    `json:"district_en"`
			DistrictCn  string    `json:"district_cn"`
			CityEn      string    `json:"city_en"`
			ProvinceCn  string    `json:"province_cn"`
			CityCn      string    `json:"city_cn"`
			StreetEn    string    `json:"street_en"`
			Isp         string    `json:"isp"`
			StreetCN    string    `json:"street_cn"`
			Owner       string    `json:"owner"`
			Radius      float64   `json:"radius"`
			CountryCN   string    `json:"country_cn"`
		} `json:"location"`
		Images     []string   `json:"images"`
		OsVersion  string     `json:"os_version"`
		Components []struct{} `json:"components"`
		Org        string     `json:"org"`
		Asn        int        `json:"asn"`
		Hostname   string     `json:"hostname"`
		IsIpv6     bool       `json:"is_ipv6"`
		Ip         string     `json:"ip"`
		OsName     string     `json:"os_name"`
		Time       string     `json:"time"`
		Transport  string     `json:"transport"`
	} `json:"data"`
	Meta struct {
		Total        int    `json:"total"`
		PaginationId string `json:"pagination_id"`
	} `json:"meta"`
}
