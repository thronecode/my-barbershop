package config

import (
	"encoding/json"
	"os"
)

var config *Config

func LoadConfig(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config = &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *Config {
	return config
}
