package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiNoteTodo struct{}

func (a *ApiNoteTodo) GetNoteTodoList(c *gin.Context) {
	var requestNoteTodoList model.RequestGetNoteTodoList
	if err := model.RequestShouldBindJSON(c, &requestNoteTodoList); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err, deptList, total := servicePrivate.ServiceNoteTodo.GetNoteTodoList(requestNoteTodoList, username); err != nil {
		global.GqaLogger.Error("获取待办便签列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取待办便签列表失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  deptList,
			Page:     requestNoteTodoList.Page,
			PageSize: requestNoteTodoList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiNoteTodo) EditNoteTodo(c *gin.Context) {
	var toEditNoteTodo model.SysNoteTodo
	if err := model.RequestShouldBindJSON(c, &toEditNoteTodo); err != nil {
		return
	}
	toEditNoteTodo.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceNoteTodo.EditNoteTodo(toEditNoteTodo); err != nil {
		global.GqaLogger.Error("编辑待办便签失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑待办便签失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑待办便签成功！")
		model.ResponseSuccessMessage("编辑待办便签成功！", c)
	}
}

func (a *ApiNoteTodo) AddNoteTodo(c *gin.Context) {
	var toAddNoteTodo model.RequestAddNoteTodo
	if err := model.RequestShouldBindJSON(c, &toAddNoteTodo); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
		},
	}
	addNoteTodo := &model.SysNoteTodo{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		TodoDetail:                        toAddNoteTodo.TodoDetail,
	}
	if err := servicePrivate.ServiceNoteTodo.AddNoteTodo(*addNoteTodo); err != nil {
		global.GqaLogger.Error("添加待办便签失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加待办便签失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加待办便签成功！", c)
	}
}

func (a *ApiNoteTodo) DeleteNoteTodoById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceNoteTodo.DeleteNoteTodoById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除待办便签失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除待办便签失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除待办便签成功！")
		model.ResponseSuccessMessage("删除待办便签成功！", c)
	}
}

func (a *ApiNoteTodo) QueryNoteTodoById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := servicePrivate.ServiceNoteTodo.QueryNoteTodoById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找待办便签失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找待办便签失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dept}, "查找待办便签成功！", c)
	}
}
