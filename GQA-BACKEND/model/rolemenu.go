package model

type SysRoleMenu struct {
	SysRoleRoleCode string `gorm:"primaryKey;" json:"sys_role_role_code"`
	SysMenuName     string `gorm:"primaryKey;" json:"sys_menu_name"`
}
