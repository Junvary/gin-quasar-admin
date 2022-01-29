package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/vote/service/private_service"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func VoteHandle(c *gin.Context) {
	var toAddScore model.RequestAddScore
	if err := c.ShouldBindJSON(&toAddScore); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.VoteHandle(&toAddScore, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("投票失败！", zap.Any("err", err))
		global.ErrorMessage("投票失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("投票成功！", c)
	}
}

func CanVoteDy(c *gin.Context)  {
	if err, canVote := private_service.CanVoteDy(utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("检查投票结果失败！", zap.Any("err", err))
		global.ErrorMessage("检查投票结果失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": canVote}, "检查投票结果成功！", c)
	}
}

func CanVoteGl(c *gin.Context)  {
	if err, canVote := private_service.CanVoteGl(utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("检查投票结果失败！", zap.Any("err", err))
		global.ErrorMessage("检查投票结果失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": canVote}, "检查投票结果成功！", c)
	}
}