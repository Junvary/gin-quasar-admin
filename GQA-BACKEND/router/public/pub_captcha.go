package public

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterCaptcha struct{}

func (r *RouterCaptcha) InitRouterCaptcha(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiCaptcha := api.GroupApiApp.ApiPublic.ApiCaptcha
	{
		Router.GET("captcha", apiCaptcha.Captcha)
	}
	return Router
}
