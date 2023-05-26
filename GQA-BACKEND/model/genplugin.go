package model

type SysGenPluginList struct {
	GqaModelWithCreatedByAndUpdatedBy
	PluginSort uint   `json:"plugin_sort" gorm:"comment:插件排序"`
	PluginCode string `json:"plugin_code" gorm:"comment:插件编码;index"`
	PluginName string `json:"plugin_name" gorm:"comment:插件名称;"`
	PluginFile string `json:"plugin_file"  gorm:"comment:插件位置"`
}

type RequestGetGenPluginList struct {
	RequestPageAndSort
	PluginSort uint   `json:"plugin_sort"`
	PluginCode string `json:"plugin_code"`
	PluginName string `json:"plugin_name"`
}

type SysGenPlugin struct {
	PluginSort  uint          `json:"plugin_sort"`
	PluginCode  string        `json:"plugin_code"`
	PluginName  string        `json:"plugin_name"`
	PluginModel []PluginModel `json:"plugin_model"`
}

type PluginModel struct {
	ModelName          string       `json:"model_name"`
	WithGqaColumn      bool         `json:"with_gqa_column"`
	WithPublicList     bool         `json:"with_public_list"`
	WithDataPermission bool         `json:"with_data_permission"`
	WithLogOperation   bool         `json:"with_log_operation"`
	ColumnList         []ColumnList `json:"column_list"`
}

type ColumnList struct {
	ColumnName    string `json:"column_name"`
	ColumnType    string `json:"column_type"`
	ColumnComment string `json:"column_comment"`
	ColumnDefault string `json:"column_default"`
}
