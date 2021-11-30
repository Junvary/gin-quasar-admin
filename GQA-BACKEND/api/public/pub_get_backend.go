package public

import (
	"gin-quasar-admin/global"
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
