package public_api

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/gqaplugin/xk/service/public_service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetWeaponLanguage(c *gin.Context)  {
	if err, language := public_service.GetWeaponLanguage(); err!=nil{
		global.GqaLog.Error("获取武器库语言列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取武器库语言列表失败！"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": language}, "获取武器库语言列表成功！", c)
	}
}

func GetWeaponNode(c *gin.Context)  {
	if err, node := public_service.GetWeaponNode(); err!=nil{
		global.GqaLog.Error("获取武器库节点列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取武器库节点列表失败！"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": node}, "获取武器库节点列表成功！", c)
	}
}
