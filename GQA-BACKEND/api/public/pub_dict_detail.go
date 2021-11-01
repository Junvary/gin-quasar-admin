package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDictDetail struct {
}

func (a *ApiDictDetail) GetDictDetailList(c *gin.Context) {
	var pageInfo global.RequestPage
	_ = c.ShouldBindJSON(&pageInfo)
	if err, dictList, total := service.GroupServiceApp.ServicePublic.GetDictDetailList(pageInfo); err != nil {
		global.GqaLog.Error("获取字典列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取字典列表失败，" + err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:     dictList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}
