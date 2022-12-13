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
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
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
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
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
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiNoteTodo) DeleteNoteTodoById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceNoteTodo.DeleteNoteTodoById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiNoteTodo) QueryNoteTodoById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := servicePrivate.ServiceNoteTodo.QueryNoteTodoById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dept}, utils.GqaI18n("FindSuccess"), c)
	}
}
