package config

import "gin-quasar-admin/config/config"

type Config struct {
	JWT     config.JWT     `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	System  config.System  `mapstructure:"system" json:"system" yaml:"system"`
	Zap     config.Zap     `mapstructure:"zap" json:"zap" yaml:"zap"`
	Mysql   config.Mysql   `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Casbin  config.Casbin  `mapstructure:"casbin" json:"casbin" yaml:"mysql"`
	Upload  config.Upload  `mapstructure:"upload" json:"upload" yaml:"upload"`
}
