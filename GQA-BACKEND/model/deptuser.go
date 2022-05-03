package model

type SysDeptUser struct {
	SysDeptDeptCode string `gorm:"comment:'部门编码';primaryKey;not null;"`
	SysUserUsername string `gorm:"comment:'用户名';primaryKey;not null;"`
}
