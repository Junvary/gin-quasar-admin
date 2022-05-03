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
		global.GqaLogger.Error("获取部门列表失败！", zap.Any("err", err))
		model.ResponseErrorMessage("获取部门列表失败，"+err.Error(), c)
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
		global.GqaLogger.Error("编辑部门失败！", zap.Any("err", err))
		model.ResponseErrorMessage("编辑部门失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "编辑部门成功！")
		model.ResponseSuccessMessage("编辑部门成功！", c)
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
		global.GqaLogger.Error("添加部门失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加部门失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加部门成功！", c)
	}
}

func (a *ApiDept) DeleteDeptById(c *gin.Context) {
	var toDeleteId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toDeleteId); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.DeleteDeptById(toDeleteId.Id); err != nil {
		global.GqaLogger.Error("删除部门失败！", zap.Any("err", err))
		model.ResponseErrorMessage("删除部门失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "删除部门成功！")
		model.ResponseSuccessMessage("删除部门成功！", c)
	}
}

func (a *ApiDept) QueryDeptById(c *gin.Context) {
	var toQueryId model.RequestQueryById
	if err := model.RequestShouldBindJSON(c, &toQueryId); err != nil {
		return
	}
	if err, dept := servicePrivate.ServiceDept.QueryDeptById(toQueryId.Id); err != nil {
		global.GqaLogger.Error("查找部门失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找部门失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessageData(gin.H{"records": dept}, "查找部门成功！", c)
	}
}

func (a *ApiDept) QueryUserByDept(c *gin.Context) {
	var deptCode model.RequestDeptCode
	if err := model.RequestShouldBindJSON(c, &deptCode); err != nil {
		return
	}
	if err, userList := servicePrivate.ServiceDept.QueryUserByDept(&deptCode); err != nil {
		global.GqaLogger.Error("查找部门用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("查找部门用户失败，"+err.Error(), c)
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
		global.GqaLogger.Error("移除部门用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("移除部门用户失败，"+err.Error(), c)
	} else {
		global.GqaLogger.Warn(utils.GetUsername(c) + "移除部门用户成功！")
		model.ResponseSuccessMessage("移除部门用户成功！", c)
	}
}

func (a *ApiDept) AddDeptUser(c *gin.Context) {
	var toAddDeptUser model.RequestDeptUserAdd
	if err := model.RequestShouldBindJSON(c, &toAddDeptUser); err != nil {
		return
	}
	if err := servicePrivate.ServiceDept.AddDeptUser(&toAddDeptUser); err != nil {
		global.GqaLogger.Error("添加部门用户失败！", zap.Any("err", err))
		model.ResponseErrorMessage("添加部门用户失败，"+err.Error(), c)
	} else {
		model.ResponseSuccessMessage("添加部门用户成功！", c)
	}
}
