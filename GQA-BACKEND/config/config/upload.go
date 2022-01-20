package config

type Upload struct {
	AvatarSavePath      string `mapstructure:"avatar-save-path" json:"avatarSavePath" yaml:"avatar-save-path"`
	AvatarUrl           string `mapstructure:"avatar-url" json:"avatarUrl" yaml:"avatar-url"`
	FileSavePath        string `mapstructure:"file-save-path" json:"fileSavePath" yaml:"file-save-path"`
	FileUrl             string `mapstructure:"file-url" json:"fileUrl" yaml:"file-url"`
	WebLogoSavePath     string `mapstructure:"web-logo-save-path" json:"webLogoSavePath" yaml:"web-logo-save-path"`
	WebLogoUrl          string `mapstructure:"web-logo-url" json:"webLogoUrl" yaml:"web-logo-url"`
	BannerImageSavePath string `mapstructure:"banner-image-save-path" json:"bannerImageSavePath" yaml:"banner-image-save-path"`
	BannerImageUrl      string `mapstructure:"banner-image-url" json:"bannerImageUrl" yaml:"banner-image-url"`
}
