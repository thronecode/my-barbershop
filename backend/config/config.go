package config

import (
	"backend/sorry"

	"encoding/json"
	"os"
)

var config *Config

func LoadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return sorry.Err(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = &Config{}
	if err = decoder.Decode(config); err != nil {
		return sorry.Err(err)
	}

	return nil
}

func GetConfig() *Config {
	return config
}
