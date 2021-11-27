package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"strconv"
	"time"
)

type ApiLogin struct {
}

func (a *ApiLogin) Login(c *gin.Context) {
	var l system.RequestLogin
	if err := c.ShouldBindJSON(&l); err != nil {
		global.ErrorMessage("表单传递不匹配，" + err.Error(), c)
		return
	}
	if global.Store.Verify(l.CaptchaId, l.Captcha, true) {
		u := &system.SysUser{Username: l.Username, Password: l.Password}
		if err, user := ServiceLogin.Login(u); err != nil {
			global.GqaLog.Error(l.Username+" 登录失败，用户名或密码错误！", zap.Any("err", err))
			global.ErrorMessage("用户名或密码错误，" + err.Error(), c)
		} else {
			a.createToken(*user, c)
		}
	} else {
		global.ErrorMessage("验证码错误！", c)
	}
}

func (a *ApiLogin) createToken(user system.SysUser, c *gin.Context) {
	jwtIssuer := utils.GetConfigBackend("jwtIssuer")
	if jwtIssuer == ""{
		global.GqaLog.Error("没有找到Jwt签发者配置，请重新初始化数据库！")
		global.ErrorMessage("没有找到Jwt签发者配置，请重新初始化数据库！", c)
		return
	}
	jwtKey := utils.GetConfigBackend("jwtKey")
	if jwtKey == ""{
		global.GqaLog.Error("没有找到Jwt密钥配置，请重新初始化数据库！")
		global.ErrorMessage("没有找到Jwt密钥配置，请重新初始化数据库！", c)
		return
	}
	jwtExpiresAt := utils.GetConfigBackend("jwtExpiresAt")
	if jwtExpiresAt == ""{
		global.GqaLog.Error("没有找到Jwt过期时间配置，请重新初始化数据库！")
		global.ErrorMessage("没有找到Jwt过期时间配置，请重新初始化数据库！", c)
		return
	}
	jwtExpiresAtInt64, err := strconv.ParseInt(jwtExpiresAt, 10, 64)
	if err!=nil{
		global.GqaLog.Error("Jwt过期时间配置格式错误！", zap.Any("err", err))
		global.ErrorMessage("Jwt过期时间配置格式错误，" + err.Error(), c)
		return
	}
	claims := system.GqaJwtClaims{
		Username:   user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + jwtExpiresAtInt64,
			Issuer:    jwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		global.GqaLog.Error("jwt签发失败！", zap.Any("err", err))
		global.ErrorMessage("jwt签发失败，" + err.Error(), c)
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
