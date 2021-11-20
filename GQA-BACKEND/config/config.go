package config

import "github.com/Junvary/gin-quasar-admin/GQA-BACKEND/config/config"

type Config struct {
	System  config.System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     config.Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   config.Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Upload  config.Upload  `mapstructure:"upload" json:"upload" yaml:"upload"`
}
