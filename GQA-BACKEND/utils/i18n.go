package utils

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GqaI18n(c *gin.Context, key string) string {
	return ginI18n.MustGetMessage(c, key)
}

func GqaI18nWithData(c *gin.Context, key string, data string) string {
	return ginI18n.MustGetMessage(c, &i18n.LocalizeConfig{
		MessageID: key,
		TemplateData: map[string]string{
			"data": data,
		},
	})
}
