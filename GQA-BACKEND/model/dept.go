package model

type SysDept struct {
	GqaModelWithCreatedByAndUpdatedBy
	//ParentCode <==> DeptCode
	ParentCode string    `json:"parent_code" gorm:"comment:父部门DeptCode;index;"`
	DeptCode   string    `json:"dept_code" gorm:"comment:部门编码;not null;uniqueIndex;"`
	DeptName   string    `json:"dept_name" gorm:"comment:部门名称;not null;"`
	Leader     string    `json:"leader" gorm:"comment:部门负责人username;"`
	LeaderUser SysUser   `json:"leader_user" gorm:"comment:部门负责人;foreignKey:Leader;references:Username"`
	Staff      []SysUser `json:"staff" gorm:"many2many:sys_dept_user;foreignKey:DeptCode;joinForeignKey:SysDeptDeptCode;references:Username;joinReferences:SysUserUsername;"`
	Children   []SysDept `json:"children" gorm:"foreignKey:ParentCode;references:DeptCode"`
}

type RequestGetDeptList struct {
	RequestPageAndSort
	DeptCode string `json:"dept_code"`
	DeptName string `json:"dept_name"`
}

type RequestAddDept struct {
	RequestAdd
	ParentCode string `json:"parent_code"`
	DeptCode   string `json:"dept_code"`
	DeptName   string `json:"dept_name"`
	Leader     string `json:"leader"`
}

type RequestDeptCode struct {
	DeptCode string `json:"dept_code"`
}

type RequestDeptUser struct {
	DeptCode string `json:"dept_code"`
	Username string `json:"username"`
}

type RequestDeptUserAdd struct {
	DeptCode string   `json:"dept_code"`
	Username []string `json:"username"`
}
