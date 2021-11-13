package system

import "gin-quasar-admin/global"

type SysApi struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Group  string `json:"group" gorm:"comment:API分组"`
	Path   string `json:"path" gorm:"comment:API路径"`
	Method string `json:"method" gorm:"comment:API请求方法"`
}

type RequestAddApi struct {
	Status string `json:"status"`
	Sort   uint   `json:"sort"`
	Remark string `json:"remark"`
	Group  string `json:"group"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type RequestApiList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddApi 中的字段
	Group  string `json:"group"`
	Path   string `json:"path"`
	Method string `json:"method"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysApi
}
