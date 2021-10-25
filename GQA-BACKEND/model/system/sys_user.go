package system

import (
	"gin-quasar-admin/global"
	"github.com/google/uuid"
)

type SysUser struct {
	global.GqaModel
	Uuid     uuid.UUID `json:"uuid" gorm:"comment:uuid;not null;uniqueIndex;"`
	Avatar   string    `json:"avatar" gorm:"comment:头像"`
	Username string    `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Password string    `json:"-"  gorm:"comment:用户登录密码"`
	Nickname string    `json:"nickname" gorm:"comment:用户昵称;default:GQA用户"`
	RealName string    `json:"realName" gorm:"comment:真实姓名"`
	Gender   string    `json:"gender" gorm:"comment:性别;default:u"`
	Mobile   string    `json:"mobile" gorm:"comment:手机号"`
	Email    string    `json:"email" gorm:"comment:邮箱"`
	Role     []SysRole `json:"role" gorm:"many2many:sys_user_role;foreignKey:Id;jointForeignKey:SysUserId;references:RoleCode;joinReferences:SysRoleRoleCode;"`
}

type RequestAddUser struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RealName string `json:"realName" gorm:"default:''"`
}
