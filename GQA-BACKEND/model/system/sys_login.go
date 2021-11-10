package system

import (
	"github.com/golang-jwt/jwt"
)

type GqaJwtClaims struct {
	Username   string `json:"username"`
	jwt.StandardClaims
}

type RequestLogin struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captchaId" binding:"required"`
}

type ResponseLogin struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RealName string `json:"realName"`
	Token    string `json:"token"`
}
