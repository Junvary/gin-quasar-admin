package utils

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"strconv"
	"time"
)

func GetUsername(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		global.GqaLog.Error("解析token负载失败！")
		return ""
	} else {
		waitUse := claims.(*system.GqaJwtClaims)
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
	claims := system.GqaJwtClaims{
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
