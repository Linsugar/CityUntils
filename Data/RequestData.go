package Data

type Response struct {
	Data  any
	State int
}

type NetResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    any    `json:"data"`
}

// Login 登录的请求值
type Login struct {
	AccountNumber string `json:"accountNumber"`
	Password      string `json:"password"`
}

// LoginResponse 登录后的响应值
type LoginResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Success bool   `json:"success"`
	Data    struct {
		User struct {
			PageNo         int         `json:"pageNo"`
			PageSize       int         `json:"pageSize"`
			Ids            interface{} `json:"ids"`
			Id             int         `json:"id"`
			AccountId      string      `json:"accountId"`
			AccountNumber  string      `json:"accountNumber"`
			Username       string      `json:"username"`
			Nickname       string      `json:"nickname"`
			Password       interface{} `json:"password"`
			Status         int         `json:"status"`
			Telephone      interface{} `json:"telephone"`
			Email          interface{} `json:"email"`
			SystemName     interface{} `json:"systemName"`
			Lng            interface{} `json:"lng"`
			Lat            interface{} `json:"lat"`
			RegionId       int         `json:"regionId"`
			RoleIdsInfo    string      `json:"roleIdsInfo"`
			CurrentRoleId  int         `json:"currentRoleId"`
			RoleId         interface{} `json:"roleId"`
			SearchRegionId interface{} `json:"searchRegionId"`
			RegionName     interface{} `json:"regionName"`
			RoleIdList     interface{} `json:"roleIdList"`
			RoleVoList     interface{} `json:"roleVoList"`
			Condition      interface{} `json:"condition"`
			Type           interface{} `json:"type"`
			UserVoList     interface{} `json:"userVoList"`
		} `json:"user"`
		Token string `json:"token"`
	} `json:"data"`
}
