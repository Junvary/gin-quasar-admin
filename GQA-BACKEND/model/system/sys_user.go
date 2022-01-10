package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
)

type SysUser struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	Username string    `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
	Nickname string    `json:"nickname" gorm:"comment:用户昵称;"`
	RealName string    `json:"realName" gorm:"comment:真实姓名"`
	Gender   string    `json:"gender" gorm:"comment:性别;default:u"`
	Mobile   string    `json:"mobile" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Role     []SysRole `json:"role" gorm:"many2many:sys_user_role;foreignKey:Username;joinForeignKey:SysUserUsername;references:RoleCode;joinReferences:SysRoleRoleCode;"`
	Dept     []SysDept `json:"dept" gorm:"many2many:sys_dept_user;foreignKey:Username;joinForeignKey:SysUserUsername;references:DeptCode;joinReferences:SysDeptDeptCode;"`
}

type RequestAddUser struct {
	Status   string `json:"status"`
	Sort     uint   `json:"sort"`
	Remark   string `json:"remark"`
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RealName string `json:"realName"`
	Gender   string `json:"gender"`
	Mobile   string `json:"mobile"`
	Email    string `json:"email"`
}

type RequestUserList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddUser 中的字段
	Username  string `json:"username"`
	RealName  string `json:"realName"`
	WithAdmin bool   `json:"withAdmin"` // 是否显示admin用户
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysUser
}

type RequestChangePassword struct {
	OldPassword  string `json:"oldPassword" binding:"required"`
	NewPassword1 string `json:"newPassword1" binding:"required"`
	NewPassword2 string `json:"newPassword2" binding:"required"`
}

type RequestChangeNickname struct {
	Nickname string `json:"nickname" binding:"required"`
}
