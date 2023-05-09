package model

type SysRoleApi struct {
	RoleCode  string `json:"role_code" gorm:"comment:角色编码;index;not null;"`
	ApiGroup  string `json:"api_group" gorm:"comment:Api分组;index;not null;"`
	ApiMethod string `json:"api_method" gorm:"comment:请求方法;default:POST;index;not null;"`
	ApiPath   string `json:"api_path" gorm:"comment:Api地址;index;not null;"`
}

type RequestEditRoleApi struct {
	RoleCode string                     `json:"role_code"`
	RoleApi  []RequestEditRoleApiDetail `json:"role_api"`
}

type RequestEditRoleApiDetail struct {
	RoleCode  string `json:"role_code"`
	ApiGroup  string `json:"api_group"`
	ApiMethod string `json:"api_method"`
	ApiPath   string `json:"api_path"`
}
