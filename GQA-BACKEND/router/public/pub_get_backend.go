package public

import "github.com/gin-gonic/gin"

type RouterGetBackend struct{}

func (r *RouterGetBackend) InitRouterGetBackend(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取网站配置列表
		Router.GET("config-backend-list", ApiGetBackend.GetBackend)
	}
	return Router
}
