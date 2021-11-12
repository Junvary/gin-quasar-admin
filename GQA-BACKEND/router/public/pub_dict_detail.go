package public

import (
	"github.com/gin-gonic/gin"
)

type RouterDictDetail struct{}

func (r *RouterDictDetail) InitRouterDictDetail(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//获取全部字典列表
		Router.POST("dict-detail-list", ApiDictDetail.GetDictDetailList)
	}
	return Router
}
