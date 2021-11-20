package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
)

type ServiceLogin struct {
}

func (s *ServiceLogin) Login(u *system.SysUser) (err error, sysUser *system.SysUser) {
	var user system.SysUser
	u.Password = utils.EncodeMD5(u.Password)
	err =global.GqaDb.Where("username=? and password= ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

