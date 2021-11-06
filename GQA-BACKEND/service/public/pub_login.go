package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
)

type ServiceLogin struct {
}

func (s *ServiceLogin) Login(u *system.SysUser) (err error, sysUser *system.SysUser) {
	var user system.SysUser
	u.Password = utils.EncodeMD5(u.Password)
	err =global.GqaDb.Where("username=? and password= ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

