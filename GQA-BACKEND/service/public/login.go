package public

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type ServiceLogin struct{}

func (s *ServiceLogin) Login(u *model.SysUser) (err error, sysUser *model.SysUser) {
	var user model.SysUser
	err = global.GqaDb.Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return err, nil
	}
	compare := utils.CompareBcrypt(user.Password, u.Password)
	if compare {
		return nil, &user
	} else {
		return errors.New("login failed"), nil
	}
}

func (s *ServiceLogin) LogLogin(username string, c *gin.Context, status string, detail string) (err error) {
	var loginLog model.SysLogLogin
	ua := user_agent.New(c.Request.UserAgent())

	loginLog.LoginUsername = username
	loginLog.LoginIp = c.ClientIP()
	name, version := ua.Browser()
	loginLog.LoginBrowser = name + " " + version
	loginLog.LoginOs = ua.OS()
	loginLog.LoginPlatform = ua.Platform()
	loginLog.LoginSuccess = status
	loginLog.Memo = detail + c.Request.UserAgent()

	err = global.GqaDb.Create(&loginLog).Error
	return err
}

func (s *ServiceLogin) SaveOnline(username string, token string) error {
	var online = model.SysUserOnline{
		Username: username,
		Token:    token,
	}
	err := global.GqaDb.Where("username = ?", online.Username).Delete(&model.SysUserOnline{}).Error
	err = global.GqaDb.Create(&online).Error
	return err
}
