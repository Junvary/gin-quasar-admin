package public

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiCheckAndInitDb struct {

}

func (a *ApiCheckAndInitDb) CheckDb(c *gin.Context)  {
	if global.GqaDb != nil {
		global.GqaLog.Info("数据库无需初始化")
		global.SuccessMessageData(gin.H{"needInit": false}, "数据库无需初始化", c)
		return
	} else {
		global.GqaLog.Info("数据库需要初始化")
		global.SuccessMessageData(gin.H{"needInit": true}, "数据库需要初始化", c)
		return
	}
}

func (a *ApiCheckAndInitDb) InitDb(c *gin.Context)  {
	if global.GqaDb != nil {
		global.GqaLog.Error("已存在数据库配置，无须再次初始化！")
		global.ErrorMessage("已存在数据库配置，无须再次初始化！", c)
		return
	}

	var initDbInfo system.RequestInitDb
	if err := c.ShouldBindJSON(&initDbInfo); err != nil {
		global.GqaLog.Error("参数校验不通过！", zap.Any("err", err))
		global.ErrorMessage("参数校验不通过！", c)
		return
	}
	if err := service.GroupServiceApp.ServiceCheckAndInitDb.CheckAndInitDb(initDbInfo); err != nil {
		global.GqaLog.Error("自动创建数据库失败！", zap.Any("err", err))
		global.ErrorMessage("自动创建数据库失败，请查看后台日志，检查后在进行初始化！", c)
		return
	}
	global.SuccessMessage("系统初始化成功！", c)
}
