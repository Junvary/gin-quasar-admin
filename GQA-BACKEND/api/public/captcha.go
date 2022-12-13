package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"strconv"
)

type ApiCaptcha struct{}

func (a *ApiCaptcha) GetCaptcha(c *gin.Context) {
	captchaKeyLong := utils.GetConfigBackend("captchaKeyLong")
	if captchaKeyLong == "" {
		global.GqaLogger.Error(utils.GqaI18n("CaptchaConfigError"))
		model.ResponseErrorMessage(utils.GqaI18n("CaptchaConfigError"), c)
		return
	}
	captchaWidth := utils.GetConfigBackend("captchaWidth")
	if captchaWidth == "" {
		global.GqaLogger.Error(utils.GqaI18n("CaptchaConfigError"))
		model.ResponseErrorMessage(utils.GqaI18n("CaptchaConfigError"), c)
		return
	}
	captchaHeight := utils.GetConfigBackend("captchaHeight")
	if captchaHeight == "" {
		global.GqaLogger.Error(utils.GqaI18n("CaptchaConfigError"))
		model.ResponseErrorMessage(utils.GqaI18n("CaptchaConfigError"), c)
		return
	}
	keyLong, _ := strconv.Atoi(captchaKeyLong)
	width, _ := strconv.Atoi(captchaWidth)
	height, _ := strconv.Atoi(captchaHeight)
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.Store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetCaptchaFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetCaptchaFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(model.ResponseCaptcha{
			CaptchaImage: b64s,
			CaptchaId:    id,
		}, utils.GqaI18n("GetCaptchaSuccess"), c)
	}
}
