package system

import "gin-quasar-admin/global"

type SysMenu struct {
	global.GqaModel
	ParentId  int       `json:"parentId" gorm:"comment:父菜单ID;"`
	Path      string    `json:"path" gorm:"comment:菜单地址;"`
	Component string    `json:"component" gorm:"comment:前端组件;"`
	Hidden    bool      `json:"hidden" gorm:"comment:是否在列表隐藏"`
	KeepAlive bool      `json:"keepAlive" gorm:"comment:是否缓存"`
	Title     string    `json:"title" gorm:"comment:菜单名称"`
	Icon      string    `json:"icon" gorm:"comment:菜单图标"`
	IsLink    bool      `json:"isLink" gorm:"comment:是否外链"`
	Role      []SysRole `json:"role" gorm:"many2many:sys_role_menu;"`
}

type ResponseMenu struct {
	Menu []SysMenu `json:"menu"`
}

type RequestAddMenu struct {
	ParentId  int    `json:"parentId"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Hidden    bool   `json:"hidden"`
	KeepAlive bool   `json:"keepAlive"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	IsLink    bool   `json:"isLink"`
}
