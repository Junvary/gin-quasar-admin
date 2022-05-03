package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type ServiceLogin struct{}

func (s *ServiceLogin) Login(u *model.SysUser) (err error, sysUser *model.SysUser) {
	var user model.SysUser
	u.Password = utils.EncodeMD5(u.Password)
	err = global.GqaDb.Where("username=? and password= ?", u.Username, u.Password).First(&user).Error
	return err, &user
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
