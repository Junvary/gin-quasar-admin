package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysMenu struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	//ParentCode 对应 Name
	ParentCode string    `json:"parentCode" gorm:"comment:父菜单Name;index;"`
	Name       string    `json:"name" gorm:"菜单Name;not null;unique;"`
	Path       string    `json:"path" gorm:"comment:菜单地址;"`
	Component  string    `json:"component" gorm:"comment:前端组件;"`
	Title      string    `json:"title" gorm:"comment:菜单名称"`
	Icon       string    `json:"icon" gorm:"comment:菜单图标"`
	Hidden     string    `json:"hidden" gorm:"comment:是否在列表隐藏;default:no;"`
	KeepAlive  string    `json:"keepAlive" gorm:"comment:是否缓存;default:no;"`
	IsLink     string    `json:"isLink" gorm:"comment:是否外链;default:no;"`
	Role       []SysRole `json:"role" gorm:"many2many:sys_role_menu;"`
}

type RequestAddMenu struct {
	Status     string `json:"status"`
	Sort       uint   `json:"sort"`
	Remark     string `json:"remark"`
	ParentCode string `json:"parentCode"`
	Name       string `json:"name"`
	Path       string `json:"path"`
	Component  string `json:"component"`
	Title      string `json:"title"`
	Icon       string `json:"icon"`
	Hidden     string `json:"hidden"`
	KeepAlive  string `json:"keepAlive"`
	IsLink     string `json:"isLink"`
}

type RequestMenuList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddMenu 中的字段
	Path  string `json:"path"`
	Title string `json:"title"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysMenu
}
