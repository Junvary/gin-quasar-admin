package system

import "gin-quasar-admin/global"

type SysApi struct {
	global.GqaModel
	Group string `json:"group" gorm:"comment:API分组"`
	Path string `json:"path" gorm:"comment:API路径"`
	Method string `json:"method" gorm:"comment:API请求方法"`
}

