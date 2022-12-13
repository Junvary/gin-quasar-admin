package middleware

import (
	"encoding/json"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func I18nHandler() gin.HandlerFunc {
	bundle := ginI18n.WithBundle(defaultBundleConfig)
	handle := ginI18n.WithGetLngHandle(getLangHandler)
	return ginI18n.Localize(bundle, handle)
}

var defaultBundleConfig = &ginI18n.BundleCfg{
	RootPath:         "config/localize",
	AcceptLanguage:   []language.Tag{language.Chinese, language.English, language.Russian},
	DefaultLanguage:  language.Chinese,
	UnmarshalFunc:    json.Unmarshal,
	FormatBundleFile: "json",
}

func getLangHandler(c *gin.Context, defaultLng string) string {
	if c == nil || c.Request == nil {
		return defaultLng
	}
	lang := c.Query("lang")
	switch lang {
	case "zh-CN":
		return "zh"
	case "en-US":
		return "en"
	case "ru":
		return "ru"
	default:
		return defaultLng
	}
}
