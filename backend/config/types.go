package config

type Config struct {
	Auth      AuthConfig       `json:"auth" binding:"required"`
	Databases []DatabaseConfig `json:"databases"`
}

type AuthConfig struct {
	Secret string `json:"secret" binding:"required"`
}

type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	ReadOnly bool   `json:"readonly"`
}
