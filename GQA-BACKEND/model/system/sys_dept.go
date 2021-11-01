package system

import "gin-quasar-admin/global"

type SysDept struct {
	global.GqaModel
	ParentId uint     `json:"parentId" gorm:"comment:父部门ID;"`
	DeptCode string  `json:"deptCode" gorm:"comment:部门编码;not null;uniqueIndex;"`
	DeptName string  `json:"deptName" gorm:"comment:部门名称;not null;unique;"`
	Phone    string  `json:"phone" gorm:"comment:联系电话;"`
	OwnerId  uint  `json:"ownerId" gorm:"comment:部门负责人ID;"`
	Owner    SysUser `json:"owner" gorm:"comment:部门负责人;foreignKey:OwnerId;references:Id"`
}

type RequestAddDept struct {
	Status   string `json:"status"`
	Sort     int    `json:"sort"`
	Remark     string `json:"remark"`
	ParentId uint     `json:"parentId"`
	DeptCode string  `json:"deptCode"`
	DeptName string  `json:"deptName"`
	Phone    string  `json:"phone"`
	OwnerId  uint  `json:"ownerId"`
}