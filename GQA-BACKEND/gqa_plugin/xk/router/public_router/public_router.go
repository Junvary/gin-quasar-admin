package public_router

import (
	"gin-quasar-admin/gqa_plugin/xk/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	{
		publicGroup.GET("news-list", public_api.GetNews)
	}
}
