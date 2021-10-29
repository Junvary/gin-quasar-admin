package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type ApiCaptcha struct {
}


func (a *ApiCaptcha) Captcha(c *gin.Context) {
	captchaCfg := global.GqaConfig.Captcha
	driver := base64Captcha.NewDriverDigit(captchaCfg.ImgHeight, captchaCfg.ImgWidth, captchaCfg.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, global.Store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.GqaLog.Error("验证码获取失败!", zap.Any("err", err))
		global.ErrorMessage("验证码获取失败，" + err.Error(), c)
	} else {
		global.SuccessMessageData(system.ResponseCaptcha{
			CaptchaId: id,
			PicPath:   b64s,
		}, "验证码获取成功！", c)
	}
}
