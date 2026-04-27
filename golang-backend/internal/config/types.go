package config

type DatabaseConfiguration struct {
	Type string `mapstructure:"type"`
	Path string `mapstructure:"path"`
}

type ServerConfiguration struct {
	Port int `mapstructure:"port"`
}

type AuthConfiguration struct {
	AccessTokenExpirationHours int    `mapstructure:"access_token_expiration_hours"`
	RefreshTokenExpirationDays int    `mapstructure:"refresh_token_expiration_days"`
	SecretKey                  string `mapstructure:"secret_key"`
}

type Configuration struct {
	Database DatabaseConfiguration `mapstructure:"database"`
	Server   ServerConfiguration   `mapstructure:"server"`
	Auth     AuthConfiguration     `mapstructure:"auth"`
}
