package model

import gqaModel "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/model"

{{ range .PluginModel }}
type GqaPlugin{{$.PluginCode}}{{.ModelName}} struct {

	gqaModel.GqaModelWithCreatedByAndUpdatedBy
	{{ range .ColumnList }}
	{{.ColumnName}} {{.ColumnType }}  `json:"{{.ColumnName}}" gorm:"comment:{{.ColumnComment}};default:{{.ColumnDefault}}"`
	{{ end }}
}

type RequestAdd{{.ModelName}} struct {
	gqaModel.RequestAdd
	{{ range .ColumnList }}
    {{.ColumnName}} {{.ColumnType }}  `json:"{{.ColumnName}}"`
    {{ end }}
}

type RequestGet{{.ModelName}}List struct {
	gqaModel.RequestPageAndSort
}
{{ end }}
