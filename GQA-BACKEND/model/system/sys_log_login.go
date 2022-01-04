package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysLogLogin struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	LoginUsername string `json:"loginUsername" gorm:"comment:登录用户名;index"`
	LoginIp       string `json:"loginIp" gorm:"comment:登录IP;index"`
	LoginBrowser  string `json:"loginBrowser" gorm:"comment:登录浏览器;"`
	LoginOs       string `json:"loginOs" gorm:"comment:登录操作系统;"`
	LoginPlatform string `json:"loginPlatform" gorm:"comment:登录平台;"`
	LoginSuccess  string `json:"loginSuccess" gorm:"是否登录成功;index"`
}

type RequestLogLoginList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 SysLogLogin 中的字段
	LoginUsername string `json:"loginUsername"`
	LoginSuccess  string `json:"loginSuccess"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysLogLogin
}
