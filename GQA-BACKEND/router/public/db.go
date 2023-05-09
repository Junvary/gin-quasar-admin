package public

import (
	"github.com/gin-gonic/gin"
)

type RouterDb struct{}

func (r *RouterDb) InitRouterDb(publicGroup *gin.RouterGroup) (R gin.IRoutes) {
	{
		publicGroup.POST("check-db", apiPublic.ApiDb.CheckDb)
		publicGroup.POST("init-db", apiPublic.ApiDb.InitDb)
	}
	return publicGroup
}
