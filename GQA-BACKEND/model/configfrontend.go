package model

type SysConfigFrontend struct {
	GqaModelWithCreatedByAndUpdatedBy
	ConfigItem  string `json:"config_item" gorm:"comment:配置项;not null;index;"`
	ItemDefault string `json:"item_default" gorm:"comment:默认值;not null;"`
	ItemCustom  string `json:"item_custom" gorm:"comment:自定义值;"`
}

type RequestGetConfigFrontendList struct {
	RequestPageAndSort
	ConfigItem string `json:"config_item"`
	Memo       string ``
}

type RequestAddConfigFrontend struct {
	RequestAdd
	ConfigItem  string `json:"config_item"`
	ItemDefault string `json:"item_default"`
	ItemCustom  string `json:"item_custom"`
}
