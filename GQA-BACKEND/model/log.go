package model

type SysLogLogin struct {
	GqaModelWithCreatedByAndUpdatedBy
	LoginUsername string `json:"login_username" gorm:"comment:登录用户名;index"`
	LoginIp       string `json:"login_ip" gorm:"comment:登录IP;index"`
	LoginBrowser  string `json:"login_browser" gorm:"comment:登录浏览器;"`
	LoginOs       string `json:"login_os" gorm:"comment:登录操作系统;"`
	LoginPlatform string `json:"login_platform" gorm:"comment:登录平台;"`
	LoginSuccess  string `json:"login_success" gorm:"是否登录成功;index"`
}

type RequestGetLogLoginList struct {
	RequestPageAndSort
	LoginUsername string `json:"login_username"`
	LoginSuccess  string `json:"login_success"`
}

type SysLogOperation struct {
	GqaModelWithCreatedByAndUpdatedBy
	OperationUsername string `json:"operation_username" gorm:"comment:请求用户名;index"`
	OperationIp       string `json:"operation_ip" gorm:"comment:请求IP;"`
	OperationMethod   string `json:"operation_method" gorm:"comment:请求方法"`
	OperationApi      string `json:"operation_api" gorm:"comment:请求Api;"`
	OperationStatus   int    `json:"operation_status" gorm:"comment:请求状态;"`
	OperationBody     string `json:"operation_body" gorm:"comment:body;type:text;"`
}

type RequestGetLogOperationList struct {
	RequestPageAndSort
	OperationUsername string `json:"operation_username"`
}
