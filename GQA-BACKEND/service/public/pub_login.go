package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mssola/user_agent"
)

type ServiceLogin struct{}

func (s *ServiceLogin) Login(u *system.SysUser) (err error, sysUser *system.SysUser) {
	var user system.SysUser
	u.Password = utils.EncodeMD5(u.Password)
	err = global.GqaDb.Where("username=? and password= ?", u.Username, u.Password).First(&user).Error
	return err, &user
}

func (s *ServiceLogin) LogLogin(username string, c *gin.Context, status string, detail string) (err error) {
	var loginLog system.SysLogLogin
	ua := user_agent.New(c.Request.UserAgent())

	loginLog.LoginUsername = username
	loginLog.LoginIp = c.ClientIP()
	name, version := ua.Browser()
	loginLog.LoginBrowser = name + " " + version
	loginLog.LoginOs =ua.OS()
	loginLog.LoginPlatform = ua.Platform()
	loginLog.LoginSuccess = status
	loginLog.Remark = detail + c.Request.UserAgent()

	err = global.GqaDb.Create(&loginLog).Error
	return err
}
