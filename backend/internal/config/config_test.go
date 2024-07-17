package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setEnvVars() {
	os.Setenv("MYBARBERSHOP_AUTH_SECRET", "test_secret")
	os.Setenv("MYBARBERSHOP_DB1_DRIVER", "postgres")
	os.Setenv("MYBARBERSHOP_DB1_HOST", "localhost")
	os.Setenv("MYBARBERSHOP_DB1_PORT", "5432")
	os.Setenv("MYBARBERSHOP_DB1_USER", "user1")
	os.Setenv("MYBARBERSHOP_DB1_PASSWORD", "password1")
	os.Setenv("MYBARBERSHOP_DB1_DBNAME", "db1")
	os.Setenv("MYBARBERSHOP_DB1_READONLY", "false")
	os.Setenv("MYBARBERSHOP_DB1_SSLMODE", "disable")
	os.Setenv("MYBARBERSHOP_DB2_DRIVER", "postgres")
	os.Setenv("MYBARBERSHOP_DB2_HOST", "localhost")
	os.Setenv("MYBARBERSHOP_DB2_PORT", "5432")
	os.Setenv("MYBARBERSHOP_DB2_USER", "user2")
	os.Setenv("MYBARBERSHOP_DB2_PASSWORD", "password2")
	os.Setenv("MYBARBERSHOP_DB2_DBNAME", "db2")
	os.Setenv("MYBARBERSHOP_DB2_READONLY", "true")
	os.Setenv("MYBARBERSHOP_DB2_SSLMODE", "require")
}

func TestGetEnvAsBool(t *testing.T) {
	os.Setenv("TEST_BOOL", "true")
	assert.True(t, getEnvAsBool("TEST_BOOL"), "Expected true when env var is true")

	os.Setenv("TEST_BOOL", "false")
	assert.False(t, getEnvAsBool("TEST_BOOL"), "Expected false when env var is false")

	os.Setenv("TEST_BOOL", "invalid")
	assert.False(t, getEnvAsBool("TEST_BOOL"), "Expected false when env var is invalid")
}

func TestLoadConfig(t *testing.T) {
	setEnvVars()

	err := LoadConfig()
	assert.NoError(t, err, "Expected no error when loading config")

	config := GetConfig()

	assert.NotNil(t, config, "Config should not be nil")
	assert.Equal(t, "test_secret", config.Auth.Secret)
	assert.Equal(t, "postgres", config.Databases[0].Driver)
	assert.Equal(t, "localhost", config.Databases[0].Host)
	assert.Equal(t, "5432", config.Databases[0].Port)
	assert.Equal(t, "user1", config.Databases[0].User)
	assert.Equal(t, "password1", config.Databases[0].Password)
	assert.Equal(t, "db1", config.Databases[0].DBName)
	assert.False(t, config.Databases[0].ReadOnly)
	assert.Equal(t, "disable", config.Databases[0].SSLMode)

	assert.Equal(t, "postgres", config.Databases[1].Driver)
	assert.Equal(t, "localhost", config.Databases[1].Host)
	assert.Equal(t, "5432", config.Databases[1].Port)
	assert.Equal(t, "user2", config.Databases[1].User)
	assert.Equal(t, "password2", config.Databases[1].Password)
	assert.Equal(t, "db2", config.Databases[1].DBName)
	assert.True(t, config.Databases[1].ReadOnly)
	assert.Equal(t, "require", config.Databases[1].SSLMode)
}

func TestSetConfig(t *testing.T) {
	newConfig := Config{
		Auth: AuthConfig{
			Secret: "new_secret",
		},
		Databases: []DatabaseConfig{
			{
				Driver:   "mysql",
				Host:     "127.0.0.1",
				Port:     "3306",
				User:     "user_new",
				Password: "password_new",
				DBName:   "db_new",
				ReadOnly: true,
				SSLMode:  "verify-full",
			},
		},
	}

	SetConfig(newConfig)

	config := GetConfig()

	assert.NotNil(t, config, "Config should not be nil")
	assert.Equal(t, "new_secret", config.Auth.Secret)
	assert.Equal(t, "mysql", config.Databases[0].Driver)
	assert.Equal(t, "127.0.0.1", config.Databases[0].Host)
	assert.Equal(t, "3306", config.Databases[0].Port)
	assert.Equal(t, "user_new", config.Databases[0].User)
	assert.Equal(t, "password_new", config.Databases[0].Password)
	assert.Equal(t, "db_new", config.Databases[0].DBName)
	assert.True(t, config.Databases[0].ReadOnly)
	assert.Equal(t, "verify-full", config.Databases[0].SSLMode)
}
