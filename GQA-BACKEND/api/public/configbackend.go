package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
)

type ApiConfigBackend struct{}

func (a *ApiConfigBackend) GetConfigBackendAll(c *gin.Context) {
	err, configList := servicePublic.ServiceConfigBackend.GetConfigBackendAll()
	if err != nil {
		model.ResponseErrorMessage("获取后台配置失败", c)
	}
	model.ResponseSuccessData(gin.H{"records": configList}, c)
}
