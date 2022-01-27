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

func VoteResultList(c *gin.Context)  {
	var getVoteScoreList model.RequestScoreResultList
	if err := c.ShouldBindJSON(&getVoteScoreList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, voter, total := private_service.VoteResultList(getVoteScoreList, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("获取投票结果列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取投票结果列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  voter,
			Page:     getVoteScoreList.Page,
			PageSize: getVoteScoreList.PageSize,
			Total:    total,
		}, c)
	}
}

func VoteResultChart(c *gin.Context)  {
	var getVoteScoreList model.RequestScoreResultList
	if err := c.ShouldBindJSON(&getVoteScoreList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, voter, total := private_service.VoteResultChart(getVoteScoreList, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("获取投票结果图形失败！", zap.Any("err", err))
		global.ErrorMessage("获取投票结果图形失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  voter,
			Page:     getVoteScoreList.Page,
			PageSize: getVoteScoreList.PageSize,
			Total:    total,
		}, c)
	}
}

