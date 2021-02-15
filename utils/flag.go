package utils

import "flag"
import "QuakeAPI/model"

func ParseInput() model.Input {
	var userInfo bool
	var key string
	var search string
	var help bool
	var output string
	var total int
	flag.StringVar(&key, "key", "", "Input Your API Key.")
	flag.IntVar(&total, "total", 100, "Number Of Queries You Want.")
	flag.StringVar(&search, "search", "", "Input Search String.")
	flag.StringVar(&output, "output", "result.txt", "Output File.")
	flag.BoolVar(&userInfo, "userinfo", false, "Show Your User Information.")
	flag.BoolVar(&help, "help", false, "Show Help Information.")
	flag.Parse()
	if help == true {
		flag.PrintDefaults()
	}
	result := model.Input{}
	result.UserInfo = userInfo
	result.Key = key
	result.Search = search
	result.Output = output
	result.Total = total
	return result
}
