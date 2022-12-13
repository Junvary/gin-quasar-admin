package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDept struct{}

func (a *ApiDept) GetDeptList(c *gin.Context) {
	var requestDeptList model.RequestGetDeptList
	if err := model.RequestShouldBindJSON(c, &requestDeptList); err != nil {
		return
	}
	if err, deptList, total := servicePrivate.ServiceDept.GetDeptList(requestDeptList); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("GetListFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  deptList,
			Page:     requestDeptList.Page,
			PageSize: requestDeptList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiDept) EditDept(c *gin.Context) {
	var toEditDept model.SysDept
	if err := model.RequestShouldBindJSON(c, &toEditDept); err != nil {
		return
	}
	toEditDept.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceDept.EditDept(toEditDept); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("EditFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("EditFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("EditSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("EditSuccess"), c)
	}
}

func (a *ApiDept) AddDept(c *gin.Context) {
	var toAddDept model.RequestAddDept
	if err := model.RequestShouldBindJSON(c, &toAddDept); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDept.Status,
			Sort:      toAddDept.Sort,
			Memo:      toAddDept.Memo,
		},
	}
	addDept := &model.SysDept{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddDept.ParentCode,
		DeptCode:                          toAddDept.DeptCode,
		DeptName:                          toAddDept.DeptName,
		Leader:                            toAddDept.Leader,
	}
	if err := servicePrivate.ServiceDept.AddDept(*addDept); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}

func (a *ApiDept) DeleteDeptById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.DeleteDeptById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiDept) QueryDeptById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := servicePrivate.ServiceDept.QueryDeptById(toQueryId.Id); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dept}, utils.GqaI18n("FindSuccess"), c)
	}
}

func (a *ApiDept) QueryUserByDept(c *gin.Context) {
	var deptCode model.RequestDeptCode
	if err := model.RequestShouldBindJSON(c, &deptCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceDept.QueryUserByDept(&deptCode); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("FindFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiDept) RemoveDeptUser(c *gin.Context) {
	var toDeleteDeptUser model.RequestDeptUser
	if err := model.RequestShouldBindJSON(c, &toDeleteDeptUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.RemoveDeptUser(&toDeleteDeptUser); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("DeleteFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("DeleteFailed")+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + utils.GqaI18n("DeleteSuccess"))
		model.ResponseSuccessMessage(utils.GqaI18n("DeleteSuccess"), c)
	}
}

func (a *ApiDept) AddDeptUser(c *gin.Context) {
	var toAddDeptUser model.RequestDeptUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddDeptUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.AddDeptUser(&toAddDeptUser); err != nil {
		global.GqaLogger.Error(utils.GqaI18n("AddFailed"), zap.Any("err", err))
		model.ResponseErrorMessage(utils.GqaI18n("AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n("AddSuccess"), c)
	}
}
