package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/service/private_service"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoterList(c *gin.Context) {
	var getVoterList model.RequestVoterList
	if err := c.ShouldBindJSON(&getVoterList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, voter, total := private_service.VoterList(getVoterList, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("获取投票人列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取投票人列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  voter,
			Page:     getVoterList.Page,
			PageSize: getVoterList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoterAdd(c *gin.Context) {
	var toAddVoter model.RequestAddVoter
	if err := c.ShouldBindJSON(&toAddVoter); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.VoterAdd(&toAddVoter, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加投票人失败！", zap.Any("err", err))
		global.ErrorMessage("添加投票人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加投票人成功！", c)
	}
}

func VoterDelete(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.VoterDelete(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除投票人失败！", zap.Any("err", err))
		global.ErrorMessage("删除投票人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除投票人成功！", c)
	}
}
