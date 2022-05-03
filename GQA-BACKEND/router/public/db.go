package public

import (
	"github.com/gin-gonic/gin"
)

type RouterDb struct{}

func (r *RouterDb) InitRouterDb(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		//检查数据库是否需要初始化
		publicGroup.POST("check-db", apiPublic.ApiDb.CheckDb)
		//初始化数据库
		publicGroup.POST("init-db", apiPublic.ApiDb.InitDb)
	}
	return publicGroup
}
