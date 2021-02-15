package utils

import (
	"QuakeAPI/model"
	"flag"
)

func ParseInput() model.Input {
	var userInfo bool
	var key string
	var search string
	var help bool
	var output string
	var total int
	result := model.Input{}
	flag.StringVar(&key, "key", "", "Input Your API Key.")
	flag.IntVar(&total, "total", 100, "Number Of Queries You Want.")
	flag.StringVar(&search, "search", "", "Input Search String.")
	flag.StringVar(&output, "output", "result.txt", "Output File.")
	flag.BoolVar(&userInfo, "userinfo", false, "Show Your User Information.")
	flag.BoolVar(&help, "help", false, "Show Help Information.")
	flag.Parse()
	if key == "" || help == true {
		flag.PrintDefaults()
		return result
	}
	result.UserInfo = userInfo
	result.Key = key
	result.Search = search
	result.Output = output
	result.Total = total
	return result
}
