package system

type RequestInitDb struct {
	DbType   string `json:"dbType" binding:"required"`
	Host     string `json:"host" binding:"required"`
	Port     string `json:"port" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	DbName   string `json:"dbName" binding:"required"`
}
