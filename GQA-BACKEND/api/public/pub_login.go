package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"time"
)

type ApiLogin struct {
}

func (a *ApiLogin) Login(c *gin.Context) {
	var l system.RequestLogin
	if err := c.ShouldBindJSON(&l); err != nil {
		global.ErrorMessage(err.Error(), c)
		return
	}
	if global.Store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := service.GroupServiceApp.ServiceLogin.Login(u); err != nil {
			global.GqaLog.Error(l.Username+" 登录失败，用户名或密码错误！", zap.Any("err", err))
			global.ErrorMessage("用户名或密码错误！", c)
		} else {
			a.createToken(*user, c)
		}
	} else {
		global.ErrorMessage("验证码错误！", c)
	}
}

func (a *ApiLogin) createToken(user system.SysUser, c *gin.Context) {
	claims := system.JwtClaims{
		Id:         user.Id,
		Username:   user.Username,
		Uuid:       user.Uuid,
		BufferTime: global.GqaConfig.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + global.GqaConfig.JWT.ExpiresAt,
			Issuer:    global.GqaConfig.JWT.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(global.GqaConfig.JWT.SecretKey))
	if err != nil {
		global.GqaLog.Error("jwt签发失败！", zap.Any("err", err))
		global.ErrorMessage("jwt签发失败！", c)
		return
	}
	global.SuccessMessageData(system.ResponseLogin{
		Avatar:   user.Avatar,
		Username: user.Username,
		Nickname: user.Nickname,
		RealName: user.RealName,
		Token:    ss,
	}, "登录成功！", c)
}
