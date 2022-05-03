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
		global.GqaLogger.Error("没有找到验证码字符数配置！")
		model.ResponseErrorMessage("没有找到验证码字符数配置！", c)
		return
	}
	captchaWidth := utils.GetConfigBackend("captchaWidth")
	if captchaWidth == "" {
		global.GqaLogger.Error("没有找到验证码宽度配置！")
		model.ResponseErrorMessage("没有找到验证码宽度配置！", c)
		return
	}
	captchaHeight := utils.GetConfigBackend("captchaHeight")
	if captchaHeight == "" {
		global.GqaLogger.Error("没有找到验证码高度配置！")
		model.ResponseErrorMessage("没有找到验证码高度配置！", c)
		return
	}
	keyLong, _ := strconv.Atoi(captchaKeyLong)
	width, _ := strconv.Atoi(captchaWidth)
	height, _ := strconv.Atoi(captchaHeight)
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.Store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GqaLogger.Error("验证码获取失败!", zap.Any("err", err))
		model.ResponseErrorMessage("验证码获取失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(model.ResponseCaptcha{
			CaptchaImage: b64s,
			CaptchaId:    id,
		}, "验证码获取成功！", c)
	}
}
