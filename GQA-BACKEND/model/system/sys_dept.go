package system

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/global"

type SysDept struct {
	UpdatedByUser *SysUser `json:"updatedByUser" gorm:"foreignKey:UpdatedBy;references:Username"`
	CreatedByUser *SysUser `json:"createdByUser" gorm:"foreignKey:CreatedBy;references:Username"`
	global.GqaModel
	//ParentCode 对应 DeptCode
	ParentCode    string    `json:"parentCode" gorm:"comment:父部门DeptCode;index;"`
	DeptCode      string    `json:"deptCode" gorm:"comment:部门编码;not null;uniqueIndex;"`
	DeptName      string    `json:"deptName" gorm:"comment:部门名称;not null;"`
	Phone         string    `json:"phone" gorm:"comment:联系电话;"`
	OwnerUsername string    `json:"ownerUsername" gorm:"comment:部门负责人username;"`
	Owner         SysUser   `json:"owner" gorm:"comment:部门负责人;foreignKey:OwnerUsername;references:Username"`
	User          []SysUser `json:"user" gorm:"many2many:sys_dept_user;foreignKey:DeptCode;jointForeignKey:SysDeptDeptCode;references:Username;joinReferences:SysUserUsername;"`
}

type RequestAddDept struct {
	Status        string `json:"status"`
	Sort          uint   `json:"sort"`
	Remark        string `json:"remark"`
	ParentCode    string `json:"parentCode"`
	DeptCode      string `json:"deptCode"`
	DeptName      string `json:"deptName"`
	Phone         string `json:"phone"`
	OwnerUsername string `json:"ownerUsername"`
}

type RequestDeptList struct {
	global.RequestPageAndSort
	//可扩充的模糊搜索项，参考上面 RequestAddDept 中的字段
	DeptCode string `json:"deptCode"`
	DeptName string `json:"deptName"`
	//全部可搜索，直接放开模型，并从service里面配置搜索逻辑
	//SysDept
}

type RequestDeptCode struct {
	DeptCode string `json:"deptCode"`
}

type RequestDeptUser struct {
	DeptCode string `json:"deptCode"`
	Username string `json:"username"`
}

type RequestDeptUserAdd struct {
	DeptCode string   `json:"deptCode"`
	Username []string `json:"username"`
}
