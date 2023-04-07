package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiConfigBackend struct{}

func (a *ApiConfigBackend) GetConfigBackendAll(c *gin.Context) {
	err, dataList := servicePublic.ServiceConfigBackend.GetConfigBackendAll()
	if err != nil {
		model.ResponseErrorMessage(utils.GqaI18n("GetBackendConfigFailed"), c)
	}
	model.ResponseSuccessData(gin.H{"records": dataList}, c)
}
