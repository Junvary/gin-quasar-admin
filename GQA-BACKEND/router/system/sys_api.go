package system

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterApi struct{}

func (r *RouterApi) InitRouterApi(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiGroup := Router.Group("api")
	apiApi := api.GroupApiApp.ApiSystem.ApiApi
	{
		// 获取api列表
		apiGroup.POST("api-list", apiApi.GetApiList)
		// 编辑api信息
		apiGroup.PUT("api-edit", apiApi.EditApi)
		// 新增api
		apiGroup.POST("api-add", apiApi.AddApi)
		// 删除api
		apiGroup.DELETE("api-delete", apiApi.DeleteApi)
	}
	return Router
}

