package model

type SysRole struct {
	GqaModelWithCreatedByAndUpdatedBy
	RoleCode                 string    `json:"role_code" gorm:"comment:角色编码;not null;uniqueIndex;"`
	RoleName                 string    `json:"role_name" gorm:"comment:角色名称;not null;index;"`
	DeptDataPermissionType   string    `json:"dept_data_permission_type" gorm:"comment:部门数据权限分类;not null;default:user;"`
	DeptDataPermissionCustom string    `json:"dept_data_permission_custom" gorm:"comment:自定义部门数据权限;type:text;"`
	User                     []SysUser `json:"user" gorm:"many2many:sys_user_role;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Username;joinReferences:SysUserUsername;"`
	Menu                     []SysMenu `json:"menu" gorm:"many2many:sys_role_menu;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Name;joinReferences:SysMenuName;"`
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
	RoleCode string            `json:"role_code"`
	RoleMenu []RequestRoleMenu `json:"role_menu"`
}

type RequestRoleMenu struct {
	RoleCode string `json:"role_code"`
	MenuName string `json:"menu_name"`
}

type RequestRoleDeptDataPermission struct {
	RoleCode                 string `json:"role_code"`
	DeptDataPermissionType   string `json:"dept_data_permission_type"`
	DeptDataPermissionCustom string `json:"dept_data_permission_custom"`
}
