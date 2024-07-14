package config

// Config is the struct that holds the configuration
type Config struct {
	Auth      AuthConfig       `json:"auth" binding:"required"`
	Databases []DatabaseConfig `json:"databases"`
}

// AuthConfig is the struct that holds the auth configuration
type AuthConfig struct {
	Secret string `json:"secret" binding:"required"`
}

// DatabaseConfig is the struct that holds the database configuration
type DatabaseConfig struct {
	Driver   string `json:"driver"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	ReadOnly bool   `json:"readonly"`
	SSLMode  string `json:"sslmode"`
}
