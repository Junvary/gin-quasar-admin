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

func CandidateList(c *gin.Context)  {
	var getCandidateList model.RequestCandidateList
	if err := c.ShouldBindJSON(&getCandidateList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, candidate, total := private_service.CandidateList(getCandidateList, utils.GetUsername(c)); err!=nil{
		global.GqaLog.Error("获取候选人列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取候选人列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  candidate,
			Page:     getCandidateList.Page,
			PageSize: getCandidateList.PageSize,
			Total:    total,
		}, c)
	}
}

func CandidateAdd(c *gin.Context)  {
	var toAddCandidate model.RequestAddCandidate
	if err := c.ShouldBindJSON(&toAddCandidate); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.CandidateAdd(&toAddCandidate, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加候选人失败！", zap.Any("err", err))
		global.ErrorMessage("添加候选人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加候选人成功！", c)
	}
}

func CandidateDelete(c *gin.Context)  {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.CandidateDelete(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除候选人失败！", zap.Any("err", err))
		global.ErrorMessage("删除候选人失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除候选人成功！", c)
	}
}
