package config

type System struct {
	Port           int    `mapstructure:"port" json:"port" yaml:"port"`
	SuccessCode    int    `mapstructure:"success-code" json:"successCode" yaml:"success-code"`
	SuccessMessage string `mapstructure:"success-message" json:"successMessage" yaml:"success-message"`
	ErrorCode      int    `mapstructure:"error-code" json:"errorCode" yaml:"error-code"`
	ErrorMessage   string `mapstructure:"error-message" json:"errorMessage" yaml:"error-message"`
	GenPluginPath  string `mapstructure:"gen-plugin-path" json:"genPluginPath" yaml:"gen-plugin-path"`
}
