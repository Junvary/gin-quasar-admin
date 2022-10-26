package model

type SysRoleButton struct {
	SysRoleRoleCode     string `gorm:"primaryKey;" json:"sys_role_role_code"`
	SysButtonButtonCode string `gorm:"primaryKey;" json:"sys_button_button_code"`
}
