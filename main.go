package main

import (
	"./core"
	"./utils"
	"strings"
)

func main() {
	utils.PrintLogo()
	input := utils.ParseInput()
	if input.UserInfo == true {
		core.GetUserInfo(input.Key)
	}
	if len(input.Search) != 0 && strings.TrimSpace(input.Search) != "" {
		_, result := core.GetServiceInfo(input.Key, input.Search, input.Total)
		utils.WriteOutput(result, input.Output)
	}
}
