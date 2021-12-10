package system

type SysUserRole struct {
	RoleCode string `gorm:"primaryKey;column:sys_role_role_code;not null;index;"`
	Username string `gorm:"primaryKey;column:sys_user_username;not null;index;"`
}
