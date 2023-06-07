package model

type SysRole struct {
	GqaModelWithCreatedByAndUpdatedBy
	RoleCode                 string      `json:"role_code" gorm:"comment:角色编码;not null;uniqueIndex;"`
	RoleName                 string      `json:"role_name" gorm:"comment:角色名称;not null;index;"`
	DeptDataPermissionType   string      `json:"dept_data_permission_type" gorm:"comment:部门数据权限分类;not null;default:deptDataPermissionType_user;"`
	DeptDataPermissionCustom string      `json:"dept_data_permission_custom" gorm:"comment:自定义部门数据权限;type:text;"`
	User                     []SysUser   `json:"user" gorm:"many2many:sys_user_role;foreignKey:RoleCode;joinForeignKey:SysRoleRoleCode;references:Username;joinReferences:SysUserUsername;"`
	Menu                     []SysMenu   `json:"menu" gorm:"many2many:sys_role_menu;foreignKey:RoleCode;joinForeignKey:SysRoleRoleCode;references:Name;joinReferences:SysMenuName;"`
	Button                   []SysButton `json:"button" gorm:"many2many:sys_role_button;foreignKey:RoleCode;joinForeignKey:SysRoleRoleCode;references:ButtonCode;joinReferences:SysButtonButtonCode;"`
	DefaultPage              string      `json:"default_page" gorm:"comment:默认首页;default:dashboard"`
	DefaultPageMenu          SysMenu     `json:"default_page_menu" gorm:"foreignKey:DefaultPage;references:Name"`
}

type RequestGetRoleList struct {
	RequestPageAndSort
	RoleCode string `json:"role_code"`
	RoleName string `json:"role_name"`
}

type RequestAddRole struct {
	RequestAdd
	RoleCode string `json:"role_code"`
	RoleName string `json:"role_name"`
}

type RequestRoleCode struct {
	RoleCode string `json:"role_code"`
}

type RequestRoleUser struct {
	RoleCode string `json:"role_code"`
	Username string `json:"username"`
}

type RequestRoleUserAdd struct {
	RoleCode string   `json:"role_code"`
	Username []string `json:"username"`
}

type RequestRoleMenuEdit struct {
	RoleCode    string              `json:"role_code"`
	RoleMenu    []RequestRoleMenu   `json:"role_menu"`
	RoleButton  []RequestRoleButton `json:"role_button"`
	DefaultPage string              `json:"default_page"`
}

type RequestRoleMenu struct {
	SysRoleRoleCode string `json:"sys_role_role_code"`
	SysMenuName     string `json:"sys_menu_name"`
}

type RequestRoleButton struct {
	SysRoleRoleCode     string `json:"sys_role_role_code"`
	SysButtonButtonCode string `json:"sys_button_button_code"`
}

type RequestRoleDeptDataPermission struct {
	RoleCode                 string `json:"role_code"`
	DeptDataPermissionType   string `json:"dept_data_permission_type"`
	DeptDataPermissionCustom string `json:"dept_data_permission_custom"`
}
