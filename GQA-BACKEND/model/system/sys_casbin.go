package system

type SysCasbin struct {
	Ptype string `json:"Ptype" gorm:"column:ptype"`
	V0    string `json:"V0" gorm:"column:v0"`
	V1    string `json:"V1" gorm:"column:v1"`
	V2    string `json:"V2" gorm:"column:v2"`
	V3    string `json:"V3" gorm:"column:v3"`
	V4    string `json:"V4" gorm:"column:v4"`
}
