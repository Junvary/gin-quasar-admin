package middleware

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func JwtHandler()gin.HandlerFunc{
	return func(c *gin.Context) {
		token := c.Request.Header.Get("gqa-token")
		if token == "" {
			global.ErrorMessageData(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				global.ErrorMessageData(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			global.ErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + global.GqaConfig.JWT.ExpiresAt
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		}
		c.Set("claims", claims)
		c.Next()
	}
}
