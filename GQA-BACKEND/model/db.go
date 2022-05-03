package model

type RequestDbInit struct {
	DbType     string `json:"db_type" binding:"required"`
	DbHost     string `json:"db_host" binding:"required"`
	DbPort     string `json:"db_port" binding:"required"`
	DbSchema   string `json:"db_schema" binding:"required"`
	DbUser     string `json:"db_user" binding:"required"`
	DbPassword string `json:"db_password" binding:"required"`
}
