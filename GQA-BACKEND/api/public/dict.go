package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiDict struct{}

func (a *ApiDict) GetDictAll(c *gin.Context) {
	err, dataList := servicePublic.ServiceDict.GetDictAll()
	if err != nil {
		model.ResponseErrorMessage(utils.GqaI18n("GetDictFailed"), c)
	}
	model.ResponseSuccessData(gin.H{"records": dataList}, c)
}
