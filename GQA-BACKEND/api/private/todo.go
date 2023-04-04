package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiTodo struct{}

func (a *ApiTodo) GetTodoList(c *gin.Context) {
	var requestTodoList model.RequestGetTodoList
	if err := model.RequestShouldBindJSON(c, &requestTodoList); err != nil {
		return
	}
	username := utils.GetUsername(c)
	if err, deptList, total := servicePrivate.ServiceTodo.GetTodoList(requestTodoList, username); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  deptList,
			Page:     requestTodoList.Page,
			PageSize: requestTodoList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiTodo) EditTodo(c *gin.Context) {
	var toEditTodo model.SysTodo
	if err := model.RequestShouldBindJSON(c, &toEditTodo); err != nil {
		return
	}
	toEditTodo.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceTodo.EditTodo(toEditTodo); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiTodo) AddTodo(c *gin.Context) {
	var toAddTodo model.RequestAddTodo
	if err := model.RequestShouldBindJSON(c, &toAddTodo); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
		},
	}
	addTodo := &model.SysTodo{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		TodoDetail:                        toAddTodo.TodoDetail,
	}
	if err := servicePrivate.ServiceTodo.AddTodo(*addTodo); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiTodo) DeleteTodoById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceTodo.DeleteTodoById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiTodo) QueryTodoById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := servicePrivate.ServiceTodo.QueryTodoById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dept}, utils.GqaI18n("FindSuccess"), c)
	}
}
