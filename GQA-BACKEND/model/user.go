package model

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type GqaModelWithCreatedByAndUpdatedBy struct {
	global.GqaModel
	CreatedByUser *SysUser `json:"created_by_user" gorm:"foreignKey:CreatedBy;references:Username"`
	UpdatedByUser *SysUser `json:"updated_by_user" gorm:"foreignKey:UpdatedBy;references:Username"`
}

type SysUser struct {
	GqaModelWithCreatedByAndUpdatedBy
	Username string    `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Nickname string    `json:"nickname" gorm:"comment:用户昵称;"`
	RealName string    `json:"real_name" gorm:"comment:真实姓名;"`
	Password string    `json:"-"  gorm:"comment:用户密码;"`
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	Gender   string    `json:"gender" gorm:"comment:性别;default:gender_unknown;"`
	Mobile   string    `json:"mobile" gorm:"comment:手机号;"`
	Email    string    `json:"email" gorm:"comment:邮箱;"`
	Role     []SysRole `json:"role" gorm:"many2many:sys_user_role;foreignKey:Username;joinForeignKey:SysUserUsername;references:RoleCode;joinReferences:SysRoleRoleCode;"`
	Dept     []SysDept `json:"dept" gorm:"many2many:sys_dept_user;foreignKey:Username;joinForeignKey:SysUserUsername;references:DeptCode;joinReferences:SysDeptDeptCode;"`
}

type RequestGetUserList struct {
	RequestPageAndSort
	DeptCode  string `json:"dept_code"`
	Username  string `json:"username"`
	RealName  string `json:"real_name"`
	WithAdmin bool   `json:"with_admin"` // 是否显示admin用户
}

type RequestAddUser struct {
	RequestAdd
	Avatar   string    `json:"avatar"`
	Username string    `json:"username"`
	Nickname string    `json:"nickname"`
	RealName string    `json:"real_name"`
	Gender   string    `json:"gender"`
	Mobile   string    `json:"mobile"`
	Email    string    `json:"email"`
	Dept     []SysDept `json:"dept"`
}

type RequestChangePassword struct {
	OldPassword  string `json:"old_password" binding:"required"`
	NewPassword1 string `json:"new_password_1" binding:"required"`
	NewPassword2 string `json:"new_password_2" binding:"required"`
}

type RequestChangeNickname struct {
	Nickname string `json:"nickname" binding:"required"`
}
