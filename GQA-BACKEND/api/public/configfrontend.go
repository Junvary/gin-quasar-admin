package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
)

type ApiConfigFrontend struct{}

func (a *ApiConfigFrontend) GetConfigFrontendAll(c *gin.Context) {
	err, configList := servicePublic.ServiceConfigFrontend.GetConfigFrontendAll()
	if err != nil {
		model.ResponseErrorMessage("获取前台配置失败", c)
	}
	model.ResponseSuccessData(gin.H{"records": configList}, c)
}
