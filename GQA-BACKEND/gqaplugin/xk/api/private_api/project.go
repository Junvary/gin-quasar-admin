package private_api

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"github.com/Junvary/gqa-plugin-xk/service/private_service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func GetProjectList(c *gin.Context) {
	var getProjectList model.RequestProjectList
	if err := c.ShouldBindJSON(&getProjectList); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, project, total := private_service.GetProjectList(getProjectList, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("获取项目列表失败！", zap.Any("err", err))
		global.ErrorMessage("获取项目列表失败！"+err.Error(), c)
	} else {
		global.SuccessData(system.ResponsePage{
			Records:  project,
			Page:     getProjectList.Page,
			PageSize: getProjectList.PageSize,
			Total:    total,
		}, c)
	}
}

func EditProject(c *gin.Context) {
	var toEditProject model.GqaPluginXkProject
	if err := c.ShouldBindJSON(&toEditProject); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	toEditProject.UpdatedBy = utils.GetUsername(c)
	if err := private_service.EditProject(toEditProject, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("编辑项目失败！", zap.Any("err", err))
		global.ErrorMessage("编辑项目失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑项目成功！", c)
	}
}

func EditProjectDetail(c *gin.Context) {
	var toEditProjectDetail model.RequestEditProjectDetail
	if err := c.ShouldBindJSON(&toEditProjectDetail); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.EditProjectDetail(toEditProjectDetail); err != nil {
		global.GqaLog.Error("编辑项目失败！", zap.Any("err", err))
		global.ErrorMessage("编辑项目失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("编辑项目成功！", c)
	}
}

func AddProject(c *gin.Context) {
	var toAddProject model.RequestAddProject
	if err := c.ShouldBindJSON(&toAddProject); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	addProject := &model.GqaPluginXkProject{
		GqaModel: global.GqaModel{
			CreatedBy: utils.GetUsername(c),
			Status:    toAddProject.Status,
			Sort:      toAddProject.Sort,
			Remark:    toAddProject.Remark,
		},
		ProjectId:      uuid.New(),
		ProjectName:    toAddProject.ProjectName,
		Demand:         toAddProject.Demand,
		LeaderUsername: toAddProject.LeaderUsername,
		Player:         toAddProject.Player,
		Language:       toAddProject.Language,
		Node:           toAddProject.Node,
	}
	if err := private_service.AddProject(*addProject, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("添加项目失败！", zap.Any("err", err))
		global.ErrorMessage("添加项目失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("添加项目成功！", c)
	}
}

func DeleteProject(c *gin.Context) {
	var toDeleteId system.RequestQueryById
	if err := c.ShouldBindJSON(&toDeleteId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err := private_service.DeleteProject(toDeleteId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("删除项目失败！", zap.Any("err", err))
		global.ErrorMessage("删除项目失败，"+err.Error(), c)
	} else {
		global.SuccessMessage("删除项目成功！", c)
	}
}

func QueryProjectById(c *gin.Context) {
	var toQueryId system.RequestQueryById
	if err := c.ShouldBindJSON(&toQueryId); err != nil {
		global.GqaLog.Error("模型绑定失败！", zap.Any("err", err))
		global.ErrorMessage("模型绑定失败，"+err.Error(), c)
		return
	}
	if err, dept := private_service.QueryProjectById(toQueryId.Id, utils.GetUsername(c)); err != nil {
		global.GqaLog.Error("查找项目失败！", zap.Any("err", err))
		global.ErrorMessage("查找项目失败，"+err.Error(), c)
	} else {
		global.SuccessMessageData(gin.H{"records": dept}, "查找项目成功！", c)
	}
}
