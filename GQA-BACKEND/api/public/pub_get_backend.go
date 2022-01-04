package public

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/gin-gonic/gin"
)

type ApiGetBackend struct {}

func (a *ApiGetBackend) GetBackend(c *gin.Context) {
	err, backend := ServiceGetBackend.GetBackend()
	if err != nil {
		global.ErrorMessage("获取后台前台配置失败，"+err.Error(), c)
	}
	global.SuccessData(gin.H{"records": backend}, c)
}
