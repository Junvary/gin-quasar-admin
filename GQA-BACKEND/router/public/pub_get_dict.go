package public

import (
	"github.com/gin-gonic/gin"
)

type RouterGetDict struct{}

func (r *RouterGetDict) InitRouterGetDict(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取全部字典列表
		Router.GET("dict-detail-list", ApiGetDict.GetDict)
	}
	return Router
}
