package model

import "github.com/golang-jwt/jwt"

type SysJwtClaims struct {
	Username  string `json:"username"`
	RefreshAt int64  `json:"refresh_at"`
	jwt.StandardClaims
}

type RequestLogin struct {
	Username  string `json:"username" binding:"required"`
	Password  string `json:"password" binding:"required"`
	Captcha   string `json:"captcha" binding:"required"`
	CaptchaId string `json:"captcha_id" binding:"required"`
}

type ResponseLogin struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	RealName string `json:"real_name"`
	Token    string `json:"token"`
}
