package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/gin-gonic/gin"
)

type ApiDict struct{}

func (a *ApiDict) GetDictAll(c *gin.Context) {
	err, dictList := servicePublic.ServiceDict.GetDictAll()
	if err != nil {
		model.ResponseErrorMessage("获取字典失败", c)
	}
	model.ResponseSuccessData(gin.H{"records": dictList}, c)
}
