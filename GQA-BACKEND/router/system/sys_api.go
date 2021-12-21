package system

import (
	"github.com/gin-gonic/gin"
)

type RouterApi struct{}

func (r *RouterApi) InitRouterApi(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiGroup := Router.Group("api")/*.Use(middleware.LogOperationHandler())*/
	{
		//获取api列表
		apiGroup.POST("api-list", ApiApi.GetApiList)
		//编辑api信息
		apiGroup.PUT("api-edit", ApiApi.EditApi)
		//新增api
		apiGroup.POST("api-add", ApiApi.AddApi)
		//删除api
		apiGroup.DELETE("api-delete", ApiApi.DeleteApi)
	}
	return Router
}

