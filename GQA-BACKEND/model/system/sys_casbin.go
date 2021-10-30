package system

type SysCasbin struct {
	Ptype    string `json:"Ptype" gorm:"column:ptype"`
	RoleCode string `json:"roleCode" gorm:"column:v0"`
	Path     string `json:"path" gorm:"column:v1"`
	Method   string `json:"method" gorm:"column:v2"`
}
