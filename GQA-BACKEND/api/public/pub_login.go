package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogin struct{}

func (a *ApiLogin) Login(c *gin.Context) {
	var l system.RequestLogin
	if err := c.ShouldBindJSON(&l); err != nil {
		global.ErrorMessage("模型绑定失败！"+err.Error(), c)
		return
	}
	if global.Store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := ServiceLogin.Login(u); err != nil {
			global.GqaLog.Error(l.Username+" 登录失败，用户名或密码错误！", zap.Any("err", err))
			global.ErrorMessage("用户名或密码错误，"+err.Error(), c)
			if err = ServiceLogin.LogLogin(l.Username, c, "no", "登录失败，用户名或密码错误！"); err != nil {
				global.GqaLog.Error("登录日志记录错误！", zap.Any("err", err))
			}
		} else {
			a.createToken(*user, c)
		}
	} else {
		global.ErrorMessage("验证码错误！", c)
		if err := ServiceLogin.LogLogin(l.Username, c, "no", "验证码错误！"); err != nil {
			global.GqaLog.Error("登录日志记录错误！", zap.Any("err", err))
		}
	}
}

func (a *ApiLogin) createToken(user system.SysUser, c *gin.Context) {
	ss := utils.CreateToken(user.Username)
	if ss == "" {
		global.GqaLog.Error("Jwt配置错误，请重新初始化数据库！")
		global.ErrorMessage("Jwt配置错误，请重新初始化数据库！", c)
		return
	}
	if err := ServiceLogin.LogLogin(user.Username, c, "yes", "登录成功！"); err != nil {
		global.GqaLog.Error("登录日志记录错误！", zap.Any("err", err))
	}
	global.SuccessMessageData(system.ResponseLogin{
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		RealName: user.RealName,
		Token:    ss,
	}, "登录成功！", c)
}
