package config

type Zap struct {
	Prefix     string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Level      string `mapstructure:"level" json:"level" yaml:"level"`
	Filepath   string `mapstructure:"filepath" json:"filepath" yaml:"filepath"`
	Filename   string `mapstructure:"filename" json:"filename" yaml:"filename"`
	MaxSize    int    `mapstructure:"max-size" json:"maxsize" yaml:"max-size"`
	MaxAge     int    `mapstructure:"max-age" json:"max_age" yaml:"max-age"`
	MaxBackups int    `mapstructure:"max-backups" json:"max_backups" yaml:"max-backups"`
}
