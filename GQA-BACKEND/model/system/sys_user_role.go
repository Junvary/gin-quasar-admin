package system

type SysUserRole struct {
	RoleCode string `gorm:"primaryKey;column:sys_role_role_code;"`
	UserId   uint   `gorm:"primaryKey;column:sys_user_id;"`
}

