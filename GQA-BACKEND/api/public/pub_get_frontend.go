package public

import (
	"gin-quasar-admin/global"
	"github.com/gin-gonic/gin"
)

type ApiGetFrontend struct {}

func (a *ApiGetFrontend) GetFrontend(c *gin.Context) {
	err, frontend := ServiceGetFrontend.GetFrontend()
	if err != nil {
		global.ErrorMessage("获取网站前台配置失败，"+err.Error(), c)
	}
	global.SuccessData(gin.H{"records": frontend}, c)
}