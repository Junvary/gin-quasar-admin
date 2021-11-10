package system

import "gin-quasar-admin/global"

type SysApi struct {
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
