package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDept struct {
}

func (a *ApiDept) GetDeptList(c *gin.Context) {
	var pageInfo system.RequestPage
	_ = c.ShouldBindJSON(&pageInfo)
	if err, deptList, total := service.GroupServiceApp.ServiceSystem.GetDeptList(pageInfo); err != nil {
		global.GqaLog.Error("获取部门列表失败：", zap.Any("err", err))
		global.ErrorMessage("获取部门列表失败，" + err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:     deptList,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiDept) EditDept(c *gin.Context) {
	var toEditDept system.SysDept
	_ = c.ShouldBindJSON(&toEditDept)
	if err := service.GroupServiceApp.ServiceSystem.EditDept(toEditDept); err != nil {
		global.GqaLog.Error("编辑部门失败!", zap.Any("err", err))
		global.ErrorMessage("编辑部门失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("编辑部门成功！", c)
	}
}

func (a *ApiDept) AddDept(c *gin.Context) {
	var toAddDept system.RequestAddDept
	_ = c.ShouldBindJSON(&toAddDept)
	addDept := &system.SysDept{
		ParentId: toAddDept.ParentId,
		DeptCode: toAddDept.DeptCode,
		DeptName: toAddDept.DeptName,
		Phone: toAddDept.Phone,
		OwnerId: toAddDept.OwnerId,
	}
	if err := service.GroupServiceApp.ServiceSystem.AddDept(*addDept); err != nil {
		global.GqaLog.Error("添加部门失败！", zap.Any("err", err))
		global.ErrorMessage("添加部门失败！"+err.Error(), c)
	} else {
		global.SuccessMessage("添加部门成功！", c)
	}
}

func (a *ApiDept) DeleteDept(c *gin.Context) {
	var toDeleteId system.RequestDelete
	_ = c.ShouldBindJSON(&toDeleteId)
	if err := service.GroupServiceApp.ServiceSystem.DeleteDept(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除部门失败！", zap.Any("err", err))
		global.ErrorMessage("删除部门失败，" + err.Error(), c)
	} else {
		global.SuccessMessage("删除部门成功！", c)
	}
}

func (a *ApiDept) QueryDeptById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	_ = c.ShouldBindJSON(&toQueryId)
	if err, dept := service.GroupServiceApp.ServiceSystem.QueryDeptById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找部门失败！", zap.Any("err", err))
		global.ErrorMessage("查找部门失败，" + err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找部门成功！", c)
	}
}
