package public

import (
	"github.com/gin-gonic/gin"
)

type RouterCaptcha struct{}

func (r *RouterCaptcha) InitRouterCaptcha(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取验证码
		Router.GET("captcha", ApiCaptcha.Captcha)
	}
	return Router
}
