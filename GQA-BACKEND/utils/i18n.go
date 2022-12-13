package utils

import ginI18n "github.com/gin-contrib/i18n"

func GqaI18n(text string) string {
	return ginI18n.MustGetMessage(text)
}
