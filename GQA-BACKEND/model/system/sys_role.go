package system

import "gin-quasar-admin/global"

type SysRole struct {
	global.GqaModel
	RoleCode string    `json:"roleCode" gorm:"comment:角色编码;not null;uniqueIndex;"`
	RoleName string    `json:"roleName" gorm:"comment:角色名称;not null;unique;"`
	User     []SysUser `json:"user" gorm:"many2many:sys_user_role;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Id;joinReferences:SysUserId;"`
	Menu     []SysMenu `json:"menu" gorm:"many2many:sys_role_menu;foreignKey:RoleCode;jointForeignKey:SysRoleRoleCode;references:Id;joinReferences:SysMenuId;"`
}

type RequestAddRole struct {
	RoleCode string `json:"roleCode"`
	RoleName string `json:"roleName"`
}