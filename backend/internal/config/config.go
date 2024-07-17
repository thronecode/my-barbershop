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
			Secret: os.Getenv("MYBARBERSHOP_AUTH_SECRET"),
		},
		Databases: []DatabaseConfig{
			{
				Driver:   os.Getenv("MYBARBERSHOP_DB1_DRIVER"),
				Host:     os.Getenv("MYBARBERSHOP_DB1_HOST"),
				Port:     os.Getenv("MYBARBERSHOP_DB1_PORT"),
				User:     os.Getenv("MYBARBERSHOP_DB1_USER"),
				Password: os.Getenv("MYBARBERSHOP_DB1_PASSWORD"),
				DBName:   os.Getenv("MYBARBERSHOP_DB1_DBNAME"),
				ReadOnly: getEnvAsBool("MYBARBERSHOP_DB1_READONLY"),
				SSLMode:  os.Getenv("MYBARBERSHOP_DB1_SSLMODE"),
			},
			{
				Driver:   os.Getenv("MYBARBERSHOP_DB2_DRIVER"),
				Host:     os.Getenv("MYBARBERSHOP_DB2_HOST"),
				Port:     os.Getenv("MYBARBERSHOP_DB2_PORT"),
				User:     os.Getenv("MYBARBERSHOP_DB2_USER"),
				Password: os.Getenv("MYBARBERSHOP_DB2_PASSWORD"),
				DBName:   os.Getenv("MYBARBERSHOP_DB2_DBNAME"),
				ReadOnly: getEnvAsBool("MYBARBERSHOP_DB2_READONLY"),
				SSLMode:  os.Getenv("MYBARBERSHOP_DB2_SSLMODE"),
			},
		},
	}

	return nil
}

// GetConfig returns the config
func GetConfig() *Config {
	return config
}

// SetConfig sets the config
func SetConfig(newConfig Config) {
	config = &newConfig
}

func getEnvAsBool(name string) bool {
	value := os.Getenv(name)
	boolValue, err := strconv.ParseBool(value)
	if err != nil {
		return false
	}
	return boolValue
}
