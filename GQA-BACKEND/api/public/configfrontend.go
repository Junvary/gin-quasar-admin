package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiConfigFrontend struct{}

func (a *ApiConfigFrontend) GetConfigFrontendAll(c *gin.Context) {
	err, dataList := servicePublic.ServiceConfigFrontend.GetConfigFrontendAll()
	if err != nil {
		model.ResponseErrorMessage(utils.GqaI18n("GetFrontendConfigFailed"), c)
	}
	model.ResponseSuccessData(gin.H{"records": dataList}, c)
}
