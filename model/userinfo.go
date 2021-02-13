package model

type UserInfo struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		Id   string `json:"id"`
		User struct {
			Username string `json:"username"`
			Fullname string `json:"fullname"`
			Email    string `json:"email"`
		} `json:"user"`
		Baned            bool   `json:"baned"`
		BanStatus        string `json:"ban_status"`
		Credit           int    `json:"email"`
		PersistentCredit int    `json:"persistent_credit"`
		Token            string `json:"token"`
		MobilePhone      string `json:"mobile_phone"`
		Source           string `json:"source"`
		PrivacyLog       struct {
			Status bool   `json:"status"`
			Time   string `json:"time"`
		} `json:"privacy_log"`
		EnterpriseInformation struct {
			Name   string `json:"name"`
			Email  string `json:"email"`
			Status string `json:"status"`
		} `json:"enterprise_information"`
		PersonalInformationStatus bool `json:"personal_information_status"`
		Role                      []struct {
			Fullname string `json:"fullname"`
			Priority int    `json:"priority"`
			Credit   int    `json:"credit"`
		} `json:"Role"`
	} `json:"data"`
	Meta struct{} `json:"meta"`
}
