package utils

import (
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

func GqaI18n(key string) string {
	return ginI18n.MustGetMessage(key)
}

func GqaI18nWithData(key, data string) string {
	return ginI18n.MustGetMessage(&i18n.LocalizeConfig{
		MessageID: key,
		TemplateData: map[string]string{
			"data": data,
		},
	})
}
