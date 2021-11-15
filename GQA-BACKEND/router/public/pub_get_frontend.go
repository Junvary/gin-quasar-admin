package public

import "github.com/gin-gonic/gin"

type RouterGetFrontend struct{}

func (r *RouterGetFrontend) InitRouterGetFrontend(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取网站配置列表
		Router.GET("config-frontend-list", ApiGetFrontend.GetFrontend)
	}
	return Router
}
