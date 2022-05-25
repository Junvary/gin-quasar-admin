package middleware

import (
	"errors"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
	"time"
)

type Jwt struct {
	SigningKey []byte
}

func (j *Jwt) ParseToken(tokenString string) (*model.SysJwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.SysJwtClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	//拿到token
	if token != nil {
		if claims, ok := token.Claims.(*model.SysJwtClaims); ok {
			//如果上面的err不为空，判断是否Expired超期，也返回claims，其他情况返回空
			if err != nil && !token.Valid {
				if vError, vOk := err.(*jwt.ValidationError); vOk {
					if vError.Errors&jwt.ValidationErrorExpired != 0 {
						return claims, errors.New("checkRefresh")
					}
				} else {
					return nil, errors.New("身份鉴别失败！")
				}
			} else if token.Valid {
				return claims, nil
			}
		}
		return nil, errors.New("身份鉴别失败！")
	} else {
		return nil, errors.New("身份鉴别失败！")
	}
}

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Gqa-Token")
		if token == "" {
			model.ResponseErrorMessageData(gin.H{"reload": true}, "未登录或非法访问！", c)
			c.Abort()
			return
		}
		j, err := MakeJwt()
		if err != nil {
			model.ResponseErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		claims, err := j.ParseToken(token)
		if err != nil {
			//ExpiresAt已经过期，检查Refresh是否过期
			if err.Error() == "checkRefresh" {
				//如果ExpiresAt在当前时间之前（已过期），但RefreshAt在当前时间之后（还未过期）
				if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt > time.Now().Unix() {
					//重新生成新的token，并插入Header里
					refreshToken := utils.CreateToken(claims.Username)
					if refreshToken != "" {
						_ = global.GqaDb.Model(&model.SysUserOnline{}).Where("username = ?", claims.Username).Update("token", refreshToken).Error
						c.Header("gqa-refresh-token", refreshToken)
						model.ResponseSuccessMessageData(gin.H{"refresh": true}, "以为你刷新令牌！", c)
						c.Abort()
						return
					}
				}
				//双重超期
				if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt < time.Now().Unix() {
					_ = global.GqaDb.Where("username = ?", claims.Username).Delete(&model.SysUserOnline{}).Error
					model.ResponseErrorMessageData(gin.H{"reload": true}, "登录已过期，请重新登录！", c)
					c.Abort()
					return
				}
			}
			model.ResponseErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)

		//判断用户是否已经被踢出
		var userOnline model.SysUserOnline
		err = global.GqaDb.Where("username = ?", claims.Username).First(&userOnline).Error
		if errors.Is(err, gorm.ErrRecordNotFound) || userOnline.Token != token {
			_ = global.GqaDb.Where("username = ?", claims.Username).Delete(&model.SysUserOnline{}).Error
			model.ResponseErrorMessageData(gin.H{"reload": true}, "登录已过期，请重新登录！", c)
			c.Abort()
			return
		}

		c.Next()
	}
}

func MakeJwt() (*Jwt, error) {
	jwtKey := utils.GetConfigBackend("jwtKey")
	if jwtKey == "" {
		return nil, errors.New("没有找到Jwt密钥配置，请重新初始化数据库！")
	}
	return &Jwt{
		[]byte(jwtKey),
	}, nil
}
