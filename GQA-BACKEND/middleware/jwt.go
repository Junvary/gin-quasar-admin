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
	if token != nil {
		if claims, ok := token.Claims.(*model.SysJwtClaims); ok {
			if err != nil && !token.Valid {
				var vError *jwt.ValidationError
				if errors.As(err, &vError) {
					if vError.Errors&jwt.ValidationErrorExpired != 0 {
						return claims, errors.New("checkRefresh")
					}
				}
			} else if token.Valid {
				return claims, nil
			}
		}
		return nil, errors.New(utils.GqaI18n(nil, "AuthFailed"))
	} else {
		return nil, errors.New(utils.GqaI18n(nil, "AuthFailed"))
	}
}

func JwtHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Gqa-Token")
		if token == "" {
			model.ResponseErrorMessageData(gin.H{"reload": true}, utils.GqaI18n(c, "AccessDisabled"), c)
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
			// ExpiresAt has expired. Check whether Refresh has expired
			if err.Error() == "checkRefresh" {
				// If ExpiresAt is before the current time (expired), but RefreshAt is after the current time (not expired)
				if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt > time.Now().Unix() {
					// Regenerate a new token and insert it into the header
					refreshToken := utils.CreateToken(claims.Username)
					if refreshToken != "" {
						_ = global.GqaDb.Model(&model.SysUserOnline{}).Where("username = ?", claims.Username).Update("token", refreshToken).Error
						c.Header("gqa-refresh-token", refreshToken)
						model.ResponseSuccessMessageData(gin.H{"refresh": true}, utils.GqaI18n(c, "RefreshToken"), c)
						c.Abort()
						return
					}
				}
				// ALL Expired
				if claims.ExpiresAt < time.Now().Unix() && claims.RefreshAt < time.Now().Unix() {
					_ = global.GqaDb.Where("username = ?", claims.Username).Delete(&model.SysUserOnline{}).Error
					model.ResponseErrorMessageData(gin.H{"reload": true}, utils.GqaI18n(c, "LoginExpired"), c)
					c.Abort()
					return
				}
			}
			model.ResponseErrorMessageData(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		c.Set("claims", claims)

		// Judge whether the user has been kicked out
		var userOnline model.SysUserOnline
		err = global.GqaDb.Where("username = ?", claims.Username).First(&userOnline).Error
		if errors.Is(err, gorm.ErrRecordNotFound) || userOnline.Token != token {
			_ = global.GqaDb.Where("username = ?", claims.Username).Delete(&model.SysUserOnline{}).Error
			model.ResponseErrorMessageData(gin.H{"reload": true}, utils.GqaI18n(c, "LoginExpired"), c)
			c.Abort()
			return
		}
		c.Next()
	}
}

func MakeJwt() (*Jwt, error) {
	jwtKey := utils.GetConfigBackend("jwtKey")
	if jwtKey == "" {
		return nil, errors.New(utils.GqaI18n(nil, "JwtConfigError"))
	}
	return &Jwt{
		[]byte(jwtKey),
	}, nil
}
