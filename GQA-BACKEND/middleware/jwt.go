package middleware

import (
	"errors"
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"time"
)

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Gqa-Token")
		if token == "" {
			global.ErrorMessageData(gin.H{"reload": true}, "未登录或非法访问！", c)
			c.Abort()
			return
		}
		j, err := MakeJwt()
		if err != nil {
			global.ErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			global.ErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt < time.Now().Unix() {
			global.ErrorMessageData(gin.H{"reload": true}, "登录已过期！", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

type Jwt struct {
	SigningKey []byte
}

func MakeJwt() (*Jwt, error) {
	jwtKey := utils.GetConfig("jwtKey")
	if jwtKey == "" {
		return nil, errors.New("没有找到Jwt密钥配置，请重新初始化数据库！")
	}
	return &Jwt{
		[]byte(jwtKey),
	}, nil
}

func (j *Jwt) ParseToken(tokenString string) (*system.GqaJwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &system.GqaJwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return nil, errors.New("身份鉴别失败！")
	}
	if token != nil {
		if claims, ok := token.Claims.(*system.GqaJwtClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, errors.New("身份鉴别失败！")
	} else {
		return nil, errors.New("身份鉴别失败！")
	}
}
