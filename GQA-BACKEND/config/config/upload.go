package config

type Upload struct {
	AvatarSavePath string   `mapstructure:"avatar-save-path" json:"avatarSavePath" yaml:"avatar-save-path"`
	AvatarUrl      string   `mapstructure:"avatar-url" json:"avatarUrl" yaml:"avatar-url"`
	FileSavePath   string   `mapstructure:"file-save-path" json:"fileSavePath" yaml:"file-save-path"`
	FileUrl        string   `mapstructure:"file-url" json:"fileUrl" yaml:"file-url"`
}
