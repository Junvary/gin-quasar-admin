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

func VoterRandomList(c *gin.Context) {
	var getVoterRandomList model.RequestVoterRandomList
	if err := c.ShouldBindJSON(&getVoterRandomList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, voter, total := private_service.VoterRandomList(getVoterRandomList, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("获取随机投票人列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取随机投票人列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  voter,
			Page:     getVoterRandomList.Page,
			PageSize: getVoterRandomList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoterRandomGet(c *gin.Context) {
	var getVoterRandomNumber model.RequestGetVoterRandom
	if err := c.ShouldBindJSON(&getVoterRandomNumber); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, voter, total := private_service.VoterRandomGet(getVoterRandomNumber, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("随机选择投票人失败！", zap.Any("err", err))
		global.ErrorMessage("随机选择投票人失败！"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": voter, "randomNumber": total}, c)
	}
}

func VoterRandomAdd(c *gin.Context) {
	var toAddVoterRandom model.RequestAddVoterRandom
	if err := c.ShouldBindJSON(&toAddVoterRandom); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.VoterRandomAdd(&toAddVoterRandom, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加投票人失败！", zap.Any("err", err))
		global.ErrorMessage("添加投票人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加投票人成功！", c)
	}
}

func VoterRandomDelete(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.VoterRandomDelete(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除投票人失败！", zap.Any("err", err))
		global.ErrorMessage("删除投票人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除投票人成功！", c)
	}
}
