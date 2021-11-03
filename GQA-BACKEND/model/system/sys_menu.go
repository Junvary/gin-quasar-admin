package system

import "gin-quasar-admin/global"

type SysMenu struct {
	global.GqaModel
	ParentId  uint      `json:"parentId" gorm:"comment:父菜单ID;"`
	Name      string    `json:"name" gorm:"菜单Name;not null;unique;"`
	Path      string    `json:"path" gorm:"comment:菜单地址;"`
	Component string    `json:"component" gorm:"comment:前端组件;"`
	Title     string    `json:"title" gorm:"comment:菜单名称"`
	Icon      string    `json:"icon" gorm:"comment:菜单图标"`
	Hidden    string    `json:"hidden" gorm:"comment:是否在列表隐藏;default:no;"`
	KeepAlive string    `json:"keepAlive" gorm:"comment:是否缓存;default:no;"`
	IsLink    string    `json:"isLink" gorm:"comment:是否外链;default:no;"`
	Role      []SysRole `json:"role" gorm:"many2many:sys_role_menu;"`
}

type RequestAddMenu struct {
	Status    string `json:"status"`
	Sort      uint   `json:"sort"`
	Remark    string `json:"remark"`
	ParentId  uint   `json:"parentId"`
	Name      string `json:"name"`
	Path      string `json:"path"`
	Component string `json:"component"`
	Title     string `json:"title"`
	Icon      string `json:"icon"`
	Hidden    string `json:"hidden"`
	KeepAlive string `json:"keepAlive"`
	IsLink    string `json:"isLink"`
}

type ResponseMenu struct {
	Menu []SysMenu `json:"menu"`
}
