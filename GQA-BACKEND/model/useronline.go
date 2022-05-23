package model

type SysUserOnline struct {
	Username string  `json:"username" gorm:"comment:用户名;not null;uniqueIndex;"`
	Token    string  `json:"token" gorm:"comment:token;not null"`
	User     SysUser `json:"user" gorm:"foreignKey:Username;references:Username"`
}

type RequestGetUserOnlineList struct {
	RequestPageAndSort
	Username string `json:"username"`
}
