package public

import "github.com/gin-gonic/gin"

type RouterCaptcha struct{}

func (r *RouterCaptcha) InitRouterCaptcha(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("/get-captcha", apiPublic.ApiCaptcha.GetCaptcha)
	}
	return publicGroup
}
