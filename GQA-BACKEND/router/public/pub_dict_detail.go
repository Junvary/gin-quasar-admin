package public

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterDictDetail struct{}

func (r *RouterDictDetail) InitRouterDictDetail(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiDictDetail := api.GroupApiApp.ApiPublic.ApiDictDetail
	{
		// 获取全部字典列表
		Router.POST("dict-detail-list", apiDictDetail.GetDictDetailList)
	}
	return Router
}
