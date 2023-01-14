package middleware

import (
	"embed"
	"encoding/json"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"strings"
)

//go:embed lang/*
var i18nFS embed.FS

type I18nEmbedLoader struct {
	FS embed.FS
}

func (c *I18nEmbedLoader) LoadMessage(path string) ([]byte, error) {
	var newPath string
	if global.GqaOsType == "windows" {
		newPath = strings.Replace(path, "\\", "/", -1)
		return c.FS.ReadFile(newPath)
	}
	return c.FS.ReadFile(path)
}

func I18nHandler() gin.HandlerFunc {
	bundle := ginI18n.WithBundle(defaultBundleConfig)
	handle := ginI18n.WithGetLngHandle(getLangHandler)
	return ginI18n.Localize(bundle, handle)
}

var defaultBundleConfig = &ginI18n.BundleCfg{
	RootPath:         "lang",
	AcceptLanguage:   []language.Tag{language.Chinese, language.English, language.Russian},
	DefaultLanguage:  language.Chinese,
	UnmarshalFunc:    json.Unmarshal,
	FormatBundleFile: "json",
	Loader:           &I18nEmbedLoader{FS: i18nFS},
}

func getLangHandler(c *gin.Context, defaultLng string) string {
	lang := c.Request.Header.Get("Gqa-Lang")
	if c == nil || c.Request == nil || lang == "" {
		return defaultLng
	}
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
