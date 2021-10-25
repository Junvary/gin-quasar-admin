package system

type SysUserRole struct {
	RoleCode string `gorm:"column:sys_role_role_code"`
	UserId   uint   `gorm:"column:sys_user_id"`
}

