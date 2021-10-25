package system

type SysRoleMenu struct {
	RoleCode string `gorm:"column:sys_role_role_code"`
	MenuId   uint   `gorm:"column:sys_menu_id"`
}
