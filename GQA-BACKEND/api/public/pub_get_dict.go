package public

import (
	"gin-quasar-admin/global"
	"github.com/gin-gonic/gin"
)

type ApiGetDict struct {}

func (a *ApiGetDict) GetDict(c *gin.Context) {
	err, dict := ServiceGetDict.GetDict()
	if err != nil {
		global.ErrorMessage("获取网站字典配置失败，"+err.Error(), c)
	}
	global.SuccessData(gin.H{"records": dict}, c)
}
