package public_router

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/example/api/public_api"
	"github.com/gin-gonic/gin"
)

func InitPublicRouter(publicGroup *gin.RouterGroup) {
	{
		publicGroup.GET("news-list", public_api.GetNews)
	}
}
