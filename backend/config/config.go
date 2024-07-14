package config

import (
	"os"
	"strconv"
)

var config *Config

// LoadConfig loads the config from the environment variables
func LoadConfig() error {
	config = &Config{
		Auth: AuthConfig{
			Secret: os.Getenv("AUTH_SECRET"),
		},
		Databases: []DatabaseConfig{
			{
				Driver:   os.Getenv("DB1_DRIVER"),
				Host:     os.Getenv("DB1_HOST"),
				Port:     os.Getenv("DB1_PORT"),
				User:     os.Getenv("DB1_USER"),
				Password: os.Getenv("DB1_PASSWORD"),
				DBName:   os.Getenv("DB1_DBNAME"),
				ReadOnly: getEnvAsBool("DB1_READONLY"),
				SSLMode:  os.Getenv("DB1_SSLMODE"),
			},
			{
				Driver:   os.Getenv("DB2_DRIVER"),
				Host:     os.Getenv("DB2_HOST"),
				Port:     os.Getenv("DB2_PORT"),
				User:     os.Getenv("DB2_USER"),
				Password: os.Getenv("DB2_PASSWORD"),
				DBName:   os.Getenv("DB2_DBNAME"),
				ReadOnly: getEnvAsBool("DB2_READONLY"),
				SSLMode:  os.Getenv("DB2_SSLMODE"),
			},
		},
	}

	return nil
}

// GetConfig returns the config
func GetConfig() *Config {
	return config
}

func getEnvAsBool(name string) bool {
	value := os.Getenv(name)
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}
