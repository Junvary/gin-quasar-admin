package config

type JWT struct {
	Issuer    string `mapstructure:"issuer" json:"issuer" yaml:"issuer"`
	SecretKey string `mapstructure:"secret-key" json:"secretKey" yaml:"secret-key"`
	ExpiresAt int64  `mapstructure:"expires-at" json:"expiresAt" yaml:"expires-at"`
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`
}
