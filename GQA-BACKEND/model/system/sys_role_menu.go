package system

type SysRoleMenu struct {
	RoleCode string `gorm:"column:sys_role_role_code"`
	MenuName string `gorm:"column:sys_menu_name"`
}
