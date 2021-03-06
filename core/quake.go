package core

import (
	"QuakeAPI/log"
	"QuakeAPI/model"
	"QuakeAPI/utils"
	"bytes"
	"encoding/json"
	"strconv"
)

type QuakeInterface interface {
	GetUserInfo(key string)
	GetServiceInfo(key string, query string, total int, pid string) (string, string)
}

type Core struct {
}

var httpClient utils.HttpClient

func init() {
	httpClient = utils.HttpClient{}
}

func (c Core) GetUserInfo(key string) {
	url := "https://quake.360.cn/api/v3/user/info"
	data := make(map[string]string)
	headers := make(map[string]string)
	headers["X-QuakeToken"] = key
	headers["Content-Type"] = "application/json"
	res := httpClient.DoGet(url, data, headers)
	var userInfo model.UserInfo
	err := json.Unmarshal(res, &userInfo)
	if err != nil {
		log.Log("unmarshal error:"+err.Error(), log.ERROR)
		return
	}
	if userInfo.Code != 0 {
		log.Log("Error API Key", log.ERROR)
		return
	}
	log.Log("Connect Success", log.INFO)
	log.Log("Your Name Is "+userInfo.Data.User.Username, log.INFO)
	log.Log("Your Email Is "+userInfo.Data.User.Email, log.INFO)
	log.Log("Your Phone Is "+userInfo.Data.MobilePhone, log.INFO)
	var roles bytes.Buffer
	for _, role := range userInfo.Data.Role {
		roles.WriteString(role.Fullname + " ")
	}
	log.Log("Your Role Is "+roles.String(), log.INFO)
}

func (c Core) GetServiceInfo(key string, query string, total int, pid string) (string, string) {
	url := "https://quake.360.cn/api/v3/scroll/quake_service"
	data := make(map[string]string)
	data["query"] = query
	data["size"] = strconv.Itoa(total)
	data["start_time"] = "2020-01-01 00:00:00"
	data["end_time"] = "2020-02-01 00:00:00"
	data["ignore_cache"] = "true"
	if pid != "" {
		data["pagination_id"] = pid
	}
	headers := make(map[string]string)
	headers["X-QuakeToken"] = key
	headers["Content-Type"] = "application/json"
	log.Log("Please Wait......", log.INFO)
	res := httpClient.DoPost(url, data, headers)
	var serviceInfo model.ServiceInfo
	err := json.Unmarshal(res, &serviceInfo)
	if err != nil {
		log.Log("unmarshal error:"+err.Error(), log.ERROR)
		return "", ""
	}
	if serviceInfo.Code != 0 {
		log.Log("Error API Key", log.ERROR)
		return "", ""
	}
	log.Log("Parsing Data......", log.INFO)
	result := bytes.Buffer{}
	for _, value := range serviceInfo.Data {
		output := value.IP + ":" + strconv.Itoa(value.Port)
		result.WriteString(output + "\n")
	}
	paginationID := serviceInfo.Meta.PaginationID
	return paginationID, result.String()
}
