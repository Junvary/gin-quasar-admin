package utils

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func GetUsername(c *gin.Context) string {
	if claims, exist := c.Get("claims"); !exist {
		global.GqaSLogger.Error("解析token负载失败！")
		return ""
	} else {
		waitUse := claims.(*model.SysJwtClaims)
		return waitUse.Username
	}
}

func CreateToken(username string) (ss string) {
	jwtIssuer := GetConfigBackend("jwtIssuer")
	if jwtIssuer == "" {
		return ""
	}
	jwtKey := GetConfigBackend("jwtKey")
	if jwtKey == "" {
		return ""
	}
	jwtExpiresAt := GetConfigBackend("jwtExpiresAt")
	if jwtExpiresAt == "" {
		return ""
	}
	jwtRefreshAt := GetConfigBackend("jwtRefreshAt")
	if jwtRefreshAt == "" {
		return ""
	}
	jwtExpiresAtInt64, err := strconv.ParseInt(jwtExpiresAt, 10, 64)
	if err != nil {
		return ""
	}
	jwtRefreshAtInt64, err := strconv.ParseInt(jwtRefreshAt, 10, 64)
	if err != nil {
		return ""
	}
	claims := model.SysJwtClaims{
		Username:  username,
		RefreshAt: time.Now().Unix() + jwtRefreshAtInt64,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,
			ExpiresAt: time.Now().Unix() + jwtExpiresAtInt64,
			Issuer:    jwtIssuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err = token.SignedString([]byte(jwtKey))
	if err != nil {
		return ""
	}
	return ss
}
