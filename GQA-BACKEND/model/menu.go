package model

type SysMenu struct {
	GqaModelWithCreatedByAndUpdatedBy
	//ParentCode 对应 Name
	ParentCode string    `json:"parent_code" gorm:"comment:父菜单Name;index;"`
	Name       string    `json:"name" gorm:"菜单Name;not null;unique;index;"`
	Path       string    `json:"path" gorm:"comment:菜单地址;"`
	Component  string    `json:"component" gorm:"comment:前端组件;"`
	Title      string    `json:"title" gorm:"comment:菜单名称;"`
	Icon       string    `json:"icon" gorm:"comment:菜单图标;"`
	Hidden     string    `json:"hidden" gorm:"comment:是否在列表隐藏;default:no;"`
	KeepAlive  string    `json:"keep_alive" gorm:"comment:是否缓存;default:no;"`
	IsLink     string    `json:"is_link" gorm:"comment:是否外链;default:no;"`
	Role       []SysRole `json:"role" gorm:"many2many:sys_role_menu;"`
}

type RequestGetMenuList struct {
	RequestPageAndSort
	Path  string `json:"path"`
	Title string `json:"title"`
}

type RequestAddMenu struct {
	RequestAdd
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
