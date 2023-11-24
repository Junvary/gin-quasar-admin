package private

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/gin-gonic/gin"
)

type ApiDept struct{}

func (a *ApiDept) GetDeptList(c *gin.Context) {
	var toGetDataList model.RequestGetDeptList
	if err := model.RequestShouldBindJSON(c, &toGetDataList); err != nil {
		return
	}
	if err, dataList, total := servicePrivate.ServiceDept.GetDeptList(toGetDataList); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "GetListFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessData(model.ResponsePage{
			Records:  dataList,
			Page:     toGetDataList.Page,
			PageSize: toGetDataList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiDept) EditDept(c *gin.Context) {
	var toEditData model.SysDept
	if err := model.RequestShouldBindJSON(c, &toEditData); err != nil {
		return
	}
	toEditData.UpdatedBy = utils.GetUsername(c)
	if err := servicePrivate.ServiceDept.EditDept(c, toEditData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "EditFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "EditSuccess"), c)
	}
}

func (a *ApiDept) AddDept(c *gin.Context) {
	var toAddData model.RequestAddDept
	if err := model.RequestShouldBindJSON(c, &toAddData); err != nil {
		return
	}
	var GqaModelWithCreatedByAndUpdatedBy = model.GqaModelWithCreatedByAndUpdatedBy{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddData.Status,
			Sort:      toAddData.Sort,
			Memo:      toAddData.Memo,
		},
	}
	addData := &model.SysDept{
		GqaModelWithCreatedByAndUpdatedBy: GqaModelWithCreatedByAndUpdatedBy,
		ParentCode:                        toAddData.ParentCode,
		DeptCode:                          toAddData.DeptCode,
		DeptName:                          toAddData.DeptName,
		Leader:                            toAddData.Leader,
	}
	if err := servicePrivate.ServiceDept.AddDept(c, *addData); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}

func (a *ApiDept) DeleteDeptById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.DeleteDeptById(toDeleteId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiDept) QueryDeptById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, data := servicePrivate.ServiceDept.QueryDeptById(toQueryId.Id); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": data}, utils.GqaI18n(c, "FindSuccess"), c)
	}
}

func (a *ApiDept) QueryUserByDept(c *gin.Context) {
	var deptCode model.RequestDeptCode
	if err := model.RequestShouldBindJSON(c, &deptCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceDept.QueryUserByDept(&deptCode); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "FindFailed")+err.Error(), c)
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
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "DeleteFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessageWithLog(utils.GetUsername(c)+utils.GqaI18n(c, "DeleteSuccess"), c)
	}
}

func (a *ApiDept) AddDeptUser(c *gin.Context) {
	var toAddDataUser model.RequestDeptUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddDataUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.AddDeptUser(&toAddDataUser); err != nil {
		model.ResponseErrorMessageWithLog(utils.GqaI18n(c, "AddFailed")+err.Error(), c)
	} else {
		model.ResponseSuccessMessage(utils.GqaI18n(c, "AddSuccess"), c)
	}
}
