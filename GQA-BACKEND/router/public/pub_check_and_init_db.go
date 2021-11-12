package public

import (
	"github.com/gin-gonic/gin"
)

type RouterCheckDb struct {}

func (r *RouterCheckDb)InitRouterCheckAndInitDb(Router *gin.RouterGroup) (R gin.IRoutes) {
	{
		//检查数据库是否需要初始化
		Router.GET("check-db", ApiCheckAndInitDb.CheckDb)
		//初始化数据库
		Router.POST("init-db", ApiCheckAndInitDb.InitDb)
	}
	return Router
}