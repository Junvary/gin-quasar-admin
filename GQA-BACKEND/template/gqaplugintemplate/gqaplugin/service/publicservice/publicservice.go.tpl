package publicservice

import (
	gqaGlobal "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"
	"github.com/Junvary/gin-quasar-admin/GQA-BACKEND/gqaplugin/{{.PluginCode}}/model"
	gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"
)

{{ range .PluginModel }}
{{ if .WithPublicList }}
func Get{{.ModelName}}List(get{{.ModelName}}List model.RequestGet{{.ModelName}}List) (err error, records []model.GqaPlugin{{$.PluginCode}}{{.ModelName}}, total int64) {
	pageSize := get{{.ModelName}}List.PageSize
	offset := get{{.ModelName}}List.PageSize * (get{{.ModelName}}List.Page - 1)
	db := gqaGlobal.GqaDb.Model(&model.GqaPlugin{{$.PluginCode}}{{.ModelName}}{})
	var recordList []model.GqaPlugin{{$.PluginCode}}{{.ModelName}}
	//配置搜索
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(pageSize).Offset(offset).Order(gqaModel.OrderByColumn(get{{.ModelName}}List.SortBy, get{{.ModelName}}List.Desc)).Preload("CreatedByUser").Find(&get{{.ModelName}}List).Error
	return err, recordList, total
}
{{ end }}
{{ end }}
