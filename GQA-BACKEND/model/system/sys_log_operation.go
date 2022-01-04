package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysLogOperation struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	OperationUsername string `json:"operationUsername" gorm:"comment:请求用户名;index"`
	OperationIp       string `json:"operationIp" gorm:"comment:请求IP;"`
	OperationMethod   string `json:"operationMethod" gorm:"comment:请求方法"`
	OperationApi      string `json:"operationApi" gorm:"comment:请求Api;"`
	OperationStatus   int    `json:"operationStatus" gorm:"comment:请求状态;"`
}

type RequestLogOperationList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 SysLogOperation 中的字段
	OperationUsername string `json:"operationUsername"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysLogOperation
}
