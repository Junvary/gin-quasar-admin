package config

type Config struct {
	Zap    Zap    `yaml:"zap"`
	Mysql  Mysql  `yaml:"mysql"`
	System System `yaml:"system"`
}

type Zap struct {
	Prefix     string `yaml:"prefix"`
	Level      string `yaml:"level"`
	Path       string `yaml:"path"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxSize"`
	MaxBackups int    `yaml:"maxBackups"`
	MaxAge     int    `yaml:"maxAge"`
}

type Mysql struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	MaxIdle  int    `yaml:"maxIdle"`
	MaxOpen  int    `yaml:"maxOpen"`
}

type System struct {
	Port                  int    `yaml:"port"`
	SuccessCode           int    `yaml:"successCode"`
	SuccessMessage        string `yaml:"successMessage"`
	ErrorCode             int    `yaml:"errorCode"`
	ErrorMessage          string `yaml:"errorMessage"`
	BindError             string `yaml:"bindError"`
	GenPluginToPath       string `yaml:"genPluginToPath"`
	GenPluginTemplatePath string `yaml:"genPluginTemplatePath"`
	ImportPath            string `yaml:"importPath"`
	ExportPath            string `yaml:"exportPath"`
	TemplatePath          string `yaml:"templatePath"`
}
