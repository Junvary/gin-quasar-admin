package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiLogin struct{}

func (a *ApiLogin) Login(c *gin.Context) {
	var l model.RequestLogin
	if err := model.RequestShouldBindJSON(c, &l); err != nil {
		return
	}
	if global.Store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &model.SysUser{Username: l.Username, Password: l.Password}
		if err, user := servicePublic.ServiceLogin.Login(u); err != nil {
			global.GqaLogger.Error(l.Username + "登录失败，用户名或密码错误！")
			model.ResponseErrorMessage("用户名或密码错误", c)
			if err = servicePublic.ServiceLogin.LogLogin(l.Username, c, "yesNo_no", "登录失败，用户名或密码错误！"); err != nil {
				global.GqaLogger.Error("登录日志记录错误！", zap.Any("err", err))
			}
		} else {
			a.createToken(*user, c)
		}
	} else {
		model.ResponseErrorMessage("验证码错误！", c)
		if err := servicePublic.ServiceLogin.LogLogin(l.Username, c, "yesNo_no", "验证码错误！"); err != nil {
			global.GqaLogger.Error("登录日志记录错误！", zap.Any("err", err))
		}
	}
}

func (a *ApiLogin) createToken(user model.SysUser, c *gin.Context) {
	ss := utils.CreateToken(user.Username)
	if ss == "" {
		global.GqaLogger.Error("Jwt配置错误，请重新初始化数据库！")
		model.ResponseErrorMessage("Jwt配置错误，请重新初始化数据库！", c)
		return
	}
	if err := servicePublic.ServiceLogin.LogLogin(user.Username, c, "yesNo_yes", "登录成功！"); err != nil {
		global.GqaLogger.Error("登录日志记录错误！", zap.Any("err", err))
	}
	if err := servicePublic.ServiceLogin.SaveOnline(user.Username, ss); err != nil {
		global.GqaLogger.Error("记录在线用户失败！", zap.Any("err", err))
	}
	model.ResponseSuccessMessageData(model.ResponseLogin{
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		RealName: user.RealName,
		Token:    ss,
	}, "登录成功！", c)
}
