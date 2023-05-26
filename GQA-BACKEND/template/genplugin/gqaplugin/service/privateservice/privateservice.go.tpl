package privateservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
	gqaServicePrivate "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/service/private"
	gqaUtils "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/utils"
	"gorm.io/gorm"
)

{{ range .PluginModel }}
func Get{{.ModelName}}List(get{{.ModelName}}List model.RequestGet{{.ModelName}}List, username string) (err error, records []model.GqaPlugin{{$.PluginCode}}{{.ModelName}}, total int64) {
	pageSize := get{{.ModelName}}List.PageSize
	offset := get{{.ModelName}}List.PageSize * (get{{.ModelName}}List.Page - 1)
	var recordList []model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	var db *gorm.DB
	{{ if .WithDataPermission}}
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})); err != nil {
		return err, recordList, 0
	}
	{{ end }}
	//配置搜索

	err = db.Count(&total).Error
	if err != nil {
		return err, recordList, 0
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(get{{.ModelName}}List.SortBy, get{{.ModelName}}List.Desc)).Preload("CreatedByUser").Find(&recordList).Error
	return err, recordList, total
}

func Edit{{.ModelName}}(toEdit{{.ModelName}} model.GqaPlugin{{$.PluginCode}}{{.ModelName}}, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})); err != nil {
		return err
	}
	var record model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	if err = db.Where("id = ?", toEdit{{.ModelName}}.Id).First(&record).Error; err != nil {
		return err
	}

	err = db.Updates(gqaUtils.MergeMap(gqaUtils.GlobalModelToMap(&toEdit{{.ModelName}}.GqaModel),
		map[string]interface{}{
		    {{ $ModelName := .ModelName }}
		    {{ range .ColumnList }}
			"{{.ColumnName}}":      toEdit{{$ModelName}}.{{.ColumnName}},
			{{ end }}
		})).Error
	return err
}

func Add{{.ModelName}}(toAdd{{.ModelName}} model.GqaPlugin{{$.PluginCode}}{{.ModelName}}, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})); err != nil {
		return err
	}
	err = db.Create(&toAdd{{.ModelName}}).Error
	return err
}

func Delete{{.ModelName}}ById(id uint, username string) (err error) {
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})); err != nil {
		return err
	}
	var record model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	if err = db.Where("id = ?", id).First(&record).Error; err != nil {
		return err
	}
	err = db.Where("id = ?", id).Unscoped().Delete(&record).Error
	return err
}

func Query{{.ModelName}}ById(id uint, username string) (err error, recordInfo model.GqaPlugin{{$.PluginCode}}{{.ModelName}}) {
	var record model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	var db *gorm.DB
	if err, db = gqaServicePrivate.DeptDataPermission(username, gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})); err != nil {
		return err, record
	}
	err = db.Preload("CreatedByUser").Preload("UpdatedByUser").First(&record, "id = ?", id).Error
	return err, record
}
{{ end }}
