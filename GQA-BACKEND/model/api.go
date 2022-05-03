package model

type SysApi struct {
	GqaModelWithCreatedByAndUpdatedBy
	ApiGroup  string `json:"api_group" gorm:"comment:Api分组;index;not null;"`
	ApiMethod string `json:"api_method" gorm:"comment:请求方法;default:POST;index;not null;"`
	ApiPath   string `json:"api_path" gorm:"comment:Api地址;index;not null;"`
}

type RequestGetApiList struct {
	RequestPageAndSort
	ApiGroup  string `json:"api_group"`
	ApiMethod string `json:"api_method"`
}

type RequestAddApi struct {
	RequestAdd
	ApiGroup  string `json:"api_group"`
	ApiMethod string `json:"api_method"`
	ApiPath   string `json:"api_path"`
}
