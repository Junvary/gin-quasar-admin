package public

import (
	"gin-quasar-admin/api"
	"github.com/gin-gonic/gin"
)

type RouterCheckDb struct {
}

func (r *RouterCheckDb)InitRouterCheckAndInitDb(Router *gin.RouterGroup) (R gin.IRoutes) {
	apiCheckAndInitDb := api.GroupApiApp.ApiPublic.ApiCheckAndInitDb
	{
		Router.GET("check-db", apiCheckAndInitDb.CheckDb)
		Router.POST("init-db", apiCheckAndInitDb.InitDb)
	}
	return Router
}