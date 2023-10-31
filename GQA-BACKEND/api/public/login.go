package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
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
			model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "LoginFailed"), c)
			if err = servicePublic.ServiceLogin.LogLogin(l.Username, c, "yesNo_no", utils.GqaI18n(c, "LoginFailed")); err != nil {
				global.GqaSLogger.Error(utils.GqaI18n(c, "LoginLogError"), "err", err)
			}
		} else {
			a.createToken(*user, c)
		}
	} else {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "CaptchaError"), c)
		if err := servicePublic.ServiceLogin.LogLogin(l.Username, c, "yesNo_no", utils.GqaI18n(c, "CaptchaError")); err != nil {
			global.GqaSLogger.Error(utils.GqaI18n(c, "LoginLogError"), "err", err)
		}
	}
}

func (a *ApiLogin) createToken(user model.SysUser, c *gin.Context) {
	ss := utils.CreateToken(user.Username)
	if ss == "" {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "JwtConfigError"), c)
		return
	}
	if err := servicePublic.ServiceLogin.LogLogin(user.Username, c, "yesNo_yes", utils.GqaI18n(c, "LoginSuccess")); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "LoginLogError"), "err", err)
	}
	if err := servicePublic.ServiceLogin.SaveOnline(user.Username, ss); err != nil {
		global.GqaSLogger.Error(utils.GqaI18n(c, "UserOnlineFailed"), "err", err)
	}
	model.ResponseSuccessMessageDataWithLog(model.ResponseLogin{
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		RealName: user.RealName,
		Token:    ss,
	}, utils.GqaI18n(c, "LoginSuccess"), c)
}
