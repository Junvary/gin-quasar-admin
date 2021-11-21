package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysApi struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	ApiGroup  string `json:"apiGroup" gorm:"comment:API分组"`
	ApiPath   string `json:"apiPath" gorm:"comment:API路径"`
	ApiMethod string `json:"apiMethod" gorm:"comment:API请求方法"`
}

type RequestAddApi struct {
	Status    string `json:"status"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	ApiGroup  string `json:"apiGroup"`
	ApiPath   string `json:"apiPath"`
	ApiMethod string `json:"apiMethod"`
}

type RequestApiList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddApi 中的字段
	ApiGroup  string `json:"apiGroup"`
	ApiPath   string `json:"apiPath"`
	ApiMethod string `json:"apiMethod"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysApi
}
