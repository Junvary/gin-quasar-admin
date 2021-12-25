package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiTodoNote struct {}


func (a *ApiTodoNote) GetTodoNoteList(c *gin.Context) {
	var requestTodoNoteList system.RequestTodoNoteList
	if err := c.ShouldBindJSON(&requestTodoNoteList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	username := utils.GetUsername(c)
	if err, deptList, total := ServiceTodoNote.GetTodoNoteList(requestTodoNoteList, username); err != nil {
		global.GqaLog.Error("获取待办便签列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取待办便签列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  deptList,
			Page:     requestTodoNoteList.Page,
			PageSize: requestTodoNoteList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiTodoNote) EditTodoNote(c *gin.Context) {
	var toEditTodoNote system.SysTodoNote
	if err := c.ShouldBindJSON(&toEditTodoNote); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditTodoNote.UpdatedBy = utils.GetUsername(c)
	if err := ServiceTodoNote.EditTodoNote(toEditTodoNote); err != nil {
		global.GqaLog.Error("编辑待办便签失败！", zap.Any("err", err))
		global.ErrorMessage("编辑待办便签失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "编辑待办便签成功！")
		global.SuccessMessage("编辑待办便签成功！", c)
	}
}

func (a *ApiTodoNote) AddTodoNote(c *gin.Context) {
	var toAddTodoNote system.RequestAddTodoNote
	if err := c.ShouldBindJSON(&toAddTodoNote); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addTodoNote := &system.SysTodoNote{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
		},
		TodoDetail: toAddTodoNote.TodoDetail,
	}
	if err := ServiceTodoNote.AddTodoNote(*addTodoNote); err != nil {
		global.GqaLog.Error("添加待办便签失败！", zap.Any("err", err))
		global.ErrorMessage("添加待办便签失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加待办便签成功！", c)
	}
}

func (a *ApiTodoNote) DeleteTodoNote(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceTodoNote.DeleteTodoNote(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除待办便签失败！", zap.Any("err", err))
		global.ErrorMessage("删除待办便签失败，"+err.Error(), c)
	} else {
		global.GqaLog.Warn(utils.GetUsername(c) + "删除待办便签成功！")
		global.SuccessMessage("删除待办便签成功！", c)
	}
}

func (a *ApiTodoNote) QueryTodoNoteById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := ServiceTodoNote.QueryTodoNoteById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找待办便签失败！", zap.Any("err", err))
		global.ErrorMessage("查找待办便签失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找待办便签成功！", c)
	}
}