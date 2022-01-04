package system

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiNotice struct{}

func (a *ApiNotice) GetNoticeList(c *gin.Context) {
	var requestNoticeList system.RequestNoticeList
	if err := c.ShouldBindJSON(&requestNoticeList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, noticeList, total := ServiceNotice.GetNoticeList(requestNoticeList); err != nil {
		global.GqaLog.Error("获取消息列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取消息列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  noticeList,
			Page:     requestNoticeList.Page,
			PageSize: requestNoticeList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiNotice) EditNotice(c *gin.Context) {
	var toEditNotice system.SysNotice
	if err := c.ShouldBindJSON(&toEditNotice); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditNotice.UpdatedBy = utils.GetUsername(c)
	if err := ServiceNotice.EditNotice(toEditNotice); err != nil {
		global.GqaLog.Error("编辑消息失败！", zap.Any("err", err))
		global.ErrorMessage("编辑消息失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "编辑消息成功！")
		global.SuccessMessage("编辑消息成功！", c)
	}
}

func (a *ApiNotice) AddNotice(c *gin.Context) {
	var toAddNotice system.RequestAddNotice
	if err := c.ShouldBindJSON(&toAddNotice); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	username := utils.GetUsername(c)
	if err := ServiceNotice.AddNotice(toAddNotice, username); err != nil {
		global.GqaLog.Error("添加消息失败！", zap.Any("err", err))
		global.ErrorMessage("添加消息失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加消息成功！", c)
	}
}

func (a *ApiNotice) DeleteNotice(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceNotice.DeleteNotice(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除消息失败！", zap.Any("err", err))
		global.ErrorMessage("删除消息失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "删除消息成功！")
		global.SuccessMessage("删除消息成功！", c)
	}
}

func (a *ApiNotice) QueryNoticeById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, role := ServiceNotice.QueryNoticeById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找消息失败！", zap.Any("err", err))
		global.ErrorMessage("查找消息失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": role}, "查找消息成功！", c)
	}
}

func (a *ApiNotice) QueryNoticeByIdRead(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	username := utils.GetUsername(c)
	if err, role := ServiceNotice.QueryNoticeByIdRead(toQueryId.Id, username); err != nil {
		global.GqaLog.Error("查找消息失败！", zap.Any("err", err))
		global.ErrorMessage("查找消息失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": role}, "查找消息成功！", c)
	}
}

func (a *ApiNotice) SendNotice(c *gin.Context) {
	var toSendNotice system.SysNotice
	if err := c.ShouldBindJSON(&toSendNotice); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toSendNotice.UpdatedBy = utils.GetUsername(c)
	if err := ServiceNotice.SendNotice(toSendNotice); err != nil {
		global.GqaLog.Error("发送消息失败！", zap.Any("err", err))
		global.ErrorMessage("发送消息失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "发送消息成功！")
		global.SuccessMessage("发送消息成功！", c)
	}
}
