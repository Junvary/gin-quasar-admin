package system

type SysDeptUser struct {
	DeptCode string `gorm:"primaryKey;column:sys_dept_dept_code;not null;index;"`
	Username string `gorm:"primaryKey;column:sys_user_username;not null;index;"`
}
