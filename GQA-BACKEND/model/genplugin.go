package model

type SysGenPlugin struct {
	PluginCode  string        `json:"plugin_code"`
	PluginName  string        `json:"plugin_name"`
	WithModel   bool          `json:"with_model"`
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
