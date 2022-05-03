package model

type SysUserRole struct {
	SysRoleRoleCode string `gorm:"comment:'角色编码';primaryKey;"`
	SysUserUsername string `gorm:"comment:'用户名';primaryKey;"`
}
