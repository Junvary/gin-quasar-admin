package private_service

import (
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/system"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"github.com/Junvary/gqa-plugin-xk/model"
	"gorm.io/gorm"
)

func GetProjectList(getProjectList model.RequestProjectList, username string) (err error, project []model.GqaPluginXkProject, total int64) {
	pageSize := getProjectList.PageSize
	offset := getProjectList.PageSize * (getProjectList.Page - 1)
	var projectList []model.GqaPluginXkProject
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkProject{})); err != nil {
		return err, projectList, 0
	}
	//配置搜索
	if getProjectList.ProjectName != "" {
		db = db.Where("project_name like ?", "%"+getProjectList.ProjectName+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(global.OrderByColumn(getProjectList.SortBy, getProjectList.Desc)).
		Preload("Leader").Find(&projectList).Error
	return err, projectList, total
}

func EditProject(toEditProject model.GqaPluginXkProject, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkProject{})); err != nil {
		return err
	}
	var project model.GqaPluginXkProject
	if err = db.Where("id = ?", toEditProject.Id).First(&project).Error; err != nil {
		return err
	}
	//err = db.Updates(&toEditProject).Error
	err = db.Updates(utils.MergeMap(utils.GlobalModelToMap(&toEditProject.GqaModel), map[string]interface{}{
		"project_name":    toEditProject.ProjectName,
		"demand":          toEditProject.Demand,
		"leader_username": toEditProject.LeaderUsername,
		"player":          toEditProject.Player,
		"language":        toEditProject.Language,
		"node":            toEditProject.Node,
	})).Error
	return err
}

func EditProjectDetail(toEditProjectDetail model.RequestEditProjectDetail) (err error) {
	var projectDetail model.GqaPluginXkProjectDetail
	if err = global.GqaDb.Where("project_id = ?", toEditProjectDetail.ProjectId).Unscoped().Delete(&projectDetail).Error; err != nil {
		return err
	}
	err = global.GqaDb.Create(&toEditProjectDetail.ProjectDetail).Error
	return err
}

func AddProject(toAddProject model.GqaPluginXkProject, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkProject{})); err != nil {
		return err
	}
	err = db.Create(&toAddProject).Error
	return err
}

func DeleteProject(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkProject{})); err != nil {
		return err
	}
	var project model.GqaPluginXkProject
	if err = db.Where("id = ?", id).First(&project).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&project).Error
	return err
}

func QueryProjectById(id uint, username string) (err error, projectInfo model.GqaPluginXkProject) {
	var project model.GqaPluginXkProject
	var db *gorm.DB
	if err, db = system.DeptDataPermission(username, global.GqaDb.Model(&model.GqaPluginXkProject{})); err != nil {
		return err, project
	}
	err = db.Preload("Leader").Preload("CreatedByUser").Preload("UpdatedByUser").Preload("ProjectDetail").
		First(&project, "id = ?", id).Error
	return err, project
}
