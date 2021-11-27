package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
	"strconv"
)

type ApiCaptcha struct{}

func (a *ApiCaptcha) Captcha(c *gin.Context) {
	captchaKeyLong := utils.GetConfigBackend("captchaKeyLong")
	if captchaKeyLong == "" {
		global.GqaLog.Error("没有找到验证码字符数配置！")
		global.ErrorMessage("没有找到验证码字符数配置！", c)
		return
	}
	captchaWidth := utils.GetConfigBackend("captchaWidth")
	if captchaWidth == "" {
		global.GqaLog.Error("没有找到验证码宽度配置！")
		global.ErrorMessage("没有找到验证码宽度配置！", c)
		return
	}
	captchaHeight := utils.GetConfigBackend("captchaHeight")
	if captchaHeight == "" {
		global.GqaLog.Error("没有找到验证码高度配置！")
		global.ErrorMessage("没有找到验证码高度配置！", c)
		return
	}
	keyLong, _ := strconv.Atoi(captchaKeyLong)
	width, _ := strconv.Atoi(captchaWidth)
	height, _ := strconv.Atoi(captchaHeight)
	driver := base64Captcha.NewDriverDigit(height, width, keyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.Store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GqaLog.Error("验证码获取失败!", zap.Any("err", err))
		global.ErrorMessage("验证码获取失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(system.ResponseCaptcha{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功！", c)
	}
}
