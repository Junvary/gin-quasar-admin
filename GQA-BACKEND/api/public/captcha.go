package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"strconv"
)

type ApiCaptcha struct{}

func (a *ApiCaptcha) GetCaptcha(c *gin.Context) {
	captchaKeyLong := utils.GetConfigBackend("captchaKeyLong")
	if captchaKeyLong == "" {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "CaptchaConfigError"), c)
		return
	}
	captchaWidth := utils.GetConfigBackend("captchaWidth")
	if captchaWidth == "" {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "CaptchaConfigError"), c)
		return
	}
	captchaHeight := utils.GetConfigBackend("captchaHeight")
	if captchaHeight == "" {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "CaptchaConfigError"), c)
		return
	}
	keyLong, _ := strconv.Atoi(captchaKeyLong)
	width, _ := strconv.Atoi(captchaWidth)
	height, _ := strconv.Atoi(captchaHeight)
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.Store)
	if id, b64s, _, err := cp.Generate(); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetCaptchaFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(model.ResponseCaptcha{
			CaptchaImage: b64s,
			CaptchaId:    id,
		}, utils.GqaI18n(c, "GetCaptchaSuccess"), c)
	}
}
