package config

type Upload struct {
	AvatarSavePath string   `mapstructure:"avatar-save-path" json:"avatarSavePath" yaml:"avatar-save-path"`
	AvatarUrl      string   `mapstructure:"avatar-url" json:"avatarUrl" yaml:"avatar-url"`
	AvatarSize     int      `mapstructure:"avatar-size" json:"avatarSize" yaml:"avatar-size"`
	AvatarExt      []string `mapstructure:"avatar-ext" json:"avatarExt" yaml:"avatar-ext"`
	FileSavePath   string   `mapstructure:"file-save-path" json:"fileSavePath" yaml:"file-save-path"`
	FileUrl        string   `mapstructure:"file-url" json:"fileUrl" yaml:"file-url"`
	FileSize       int      `mapstructure:"file-size" json:"fileSize" yaml:"file-size"`
	FileExt        []string `mapstructure:"file-ext" json:"fileExt" yaml:"file-ext"`
}
