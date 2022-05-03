package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiNotice struct{}

func (a *ApiNotice) GetNoticeList(c *gin.Context) {
	var requestNoticeList model.RequestGetNoticeList
	if err := model.RequestShouldBindJSON(c, &requestNoticeList); err != nil {
		return
	}
	if err, noticeList, total := servicePrivate.ServiceNotice.GetNoticeList(requestNoticeList); err != nil {
		global.GqaLogger.Error("获取消息列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取消息列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  noticeList,
			Page:     requestNoticeList.Page,
			PageSize: requestNoticeList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiNotice) EditNotice(c *gin.Context) {
	var toEditNotice model.SysNotice
	if err := model.RequestShouldBindJSON(c, &toEditNotice); err != nil {
		return
	}
	toEditNotice.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.EditNotice(toEditNotice); err != nil {
		global.GqaLogger.Error("编辑消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑消息失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑消息成功！")
		model.ResponseSuccessMessage("编辑消息成功！", c)
	}
}

func (a *ApiNotice) AddNotice(c *gin.Context) {
	var toAddNotice model.RequestAddNotice
	if err := model.RequestShouldBindJSON(c, &toAddNotice); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.AddNotice(toAddNotice, username); err != nil {
		global.GqaLogger.Error("添加消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加消息失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加消息成功！", c)
	}
}

func (a *ApiNotice) DeleteNoticeById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceNotice.DeleteNoticeById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除消息失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除消息成功！")
		model.ResponseSuccessMessage("删除消息成功！", c)
	}
}

func (a *ApiNotice) QueryNoticeById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, role := servicePrivate.ServiceNotice.QueryNoticeById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找消息失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": role}, "查找消息成功！", c)
	}
}

func (a *ApiNotice) QueryNoticeReadById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err, role := servicePrivate.ServiceNotice.QueryNoticeReadById(toQueryId.Id, username); err != nil {
		global.GqaLogger.Error("查找消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找消息失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": role}, "查找消息成功！", c)
	}
}

func (a *ApiNotice) SendNotice(c *gin.Context) {
	var toSendNotice model.SysNotice
	if err := model.RequestShouldBindJSON(c, &toSendNotice); err != nil {
		return
	}
	toSendNotice.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceNotice.SendNotice(toSendNotice); err != nil {
		global.GqaLogger.Error("发送消息失败！", zap.Any("err", err))
		model.ResponseErrorMessage("发送消息失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "发送消息成功！")
		model.ResponseSuccessMessage("发送消息成功！", c)
	}
}
