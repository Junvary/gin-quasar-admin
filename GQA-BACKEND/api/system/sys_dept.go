package system

import (
	"gin-quasar-admin/global"
	"gin-quasar-admin/model/system"
	"gin-quasar-admin/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ApiDept struct{}

func (a *ApiDept) GetDeptList(c *gin.Context) {
	var requestDeptList system.RequestDeptList
	if err := c.ShouldBindJSON(&requestDeptList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, deptList, total := ServiceDept.GetDeptList(requestDeptList); err != nil {
		global.GqaLog.Error("获取部门列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取部门列表失败，"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  deptList,
			Page:     requestDeptList.Page,
			PageSize: requestDeptList.PageSize,
			Total:    total,
		}, c)
	}
}

func (a *ApiDept) EditDept(c *gin.Context) {
	var toEditDept system.SysDept
	if err := c.ShouldBindJSON(&toEditDept); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditDept.UpdatedBy = utils.GetUsername(c)
	if err := ServiceDept.EditDept(toEditDept); err != nil {
		global.GqaLog.Error("编辑部门失败！", zap.Any("err", err))
		global.ErrorMessage("编辑部门失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑部门成功！", c)
	}
}

func (a *ApiDept) AddDept(c *gin.Context) {
	var toAddDept system.RequestAddDept
	if err := c.ShouldBindJSON(&toAddDept); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addDept := &system.SysDept{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddDept.Status,
			Sort:      toAddDept.Sort,
			Remark:    toAddDept.Remark,
		},
		ParentCode:    toAddDept.ParentCode,
		DeptCode:      toAddDept.DeptCode,
		DeptName:      toAddDept.DeptName,
		Phone:         toAddDept.Phone,
		OwnerUsername: toAddDept.OwnerUsername,
	}
	if err := ServiceDept.AddDept(*addDept); err != nil {
		global.GqaLog.Error("添加部门失败！", zap.Any("err", err))
		global.ErrorMessage("添加部门失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加部门成功！", c)
	}
}

func (a *ApiDept) DeleteDept(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceDept.DeleteDept(toDeleteId.Id); err != nil {
		global.GqaLog.Error("删除部门失败！", zap.Any("err", err))
		global.ErrorMessage("删除部门失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除部门成功！", c)
	}
}

func (a *ApiDept) QueryDeptById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := ServiceDept.QueryDeptById(toQueryId.Id); err != nil {
		global.GqaLog.Error("查找部门失败！", zap.Any("err", err))
		global.ErrorMessage("查找部门失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找部门成功！", c)
	}
}

func (a *ApiDept) QueryUserByDept(c *gin.Context) {
	var deptCode system.RequestDeptCode
	if err := c.ShouldBindJSON(&deptCode); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, userList := ServiceDept.QueryUserByDept(&deptCode); err != nil {
		global.GqaLog.Error("查找部门用户失败！", zap.Any("err", err))
		global.ErrorMessage("查找部门用户失败，"+err.Error(), c)
	} else {
		global.SuccessData(gin.H{"records": userList}, c)
	}
}

func (a *ApiDept) RemoveDeptUser(c *gin.Context) {
	var toDeleteDeptUser system.RequestDeptUser
	if err := c.ShouldBindJSON(&toDeleteDeptUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceDept.RemoveDeptUser(&toDeleteDeptUser); err != nil {
		global.GqaLog.Error("移除部门用户失败！", zap.Any("err", err))
		global.ErrorMessage("移除部门用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("移除部门用户成功！", c)
	}
}

func (a *ApiDept) AddDeptUser(c *gin.Context) {
	var toAddDeptUser system.RequestDeptUserAdd
	if err := c.ShouldBindJSON(&toAddDeptUser); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := ServiceDept.AddDeptUser(&toAddDeptUser); err != nil {
		global.GqaLog.Error("添加部门用户失败！", zap.Any("err", err))
		global.ErrorMessage("添加部门用户失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加部门用户成功！", c)
	}
}
