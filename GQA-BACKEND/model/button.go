package model

type SysButton struct {
	MenuName   string `json:"menu_name" gorm:"comment:菜单Name;not null;primaryKey;"`
	ButtonName string `json:"button_name" gorm:"comment:按钮名称;not null;"`
	ButtonCode string `json:"button_code" gorm:"comment:按钮编码;not null;primaryKey;unique"`
}
