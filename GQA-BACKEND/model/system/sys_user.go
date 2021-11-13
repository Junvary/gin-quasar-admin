package system

import (
	"gin-quasar-admin/global"
)

type SysUser struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	Username string    `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
	Nickname string    `json:"nickname" gorm:"comment:用户昵称;default:GQA用户"`
	RealName string    `json:"realName" gorm:"comment:真实姓名"`
	Gender   string    `json:"gender" gorm:"comment:性别;default:u"`
	Mobile   string    `json:"mobile" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Role     []SysRole `json:"role" gorm:"many2many:sys_user_role;foreignKey:Id;joinForeignKey:SysUserId;references:RoleCode;joinReferences:SysRoleRoleCode;"`
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
	Username string `json:"username"`
	RealName string `json:"realName"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysUser
}
