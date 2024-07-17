package utils

import (
	"github.com/thronecode/my-barbershop/backend/internal/config"

	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCheckPassword(t *testing.T) {
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)

	assert.True(t, CheckPassword(hashedPassword, password))

	assert.False(t, CheckPassword(hashedPassword, "wrongpassword"))
}

func TestHashPassword(t *testing.T) {
	password := "testpassword"
	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	assert.NotEqual(t, password, hashedPassword)
}

func TestGenerateToken(t *testing.T) {
	config.SetConfig(config.Config{
		Auth: config.AuthConfig{
			Secret: "testsecret",
		},
	})

	token, err := GenerateToken("testuser", 1)
	require.NoError(t, err)
	assert.NotEmpty(t, token)

	claims := jwt.MapClaims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.GetConfig().Auth.Secret), nil
	})
	require.NoError(t, err)
	assert.Equal(t, "testuser", claims["username"])
	assert.Equal(t, float64(1), claims["id"])
}

func TestIsValidToken(t *testing.T) {
	config.SetConfig(config.Config{
		Auth: config.AuthConfig{
			Secret: "testsecret",
		},
	})

	token, err := GenerateToken("testuser", 1)
	require.NoError(t, err)

	assert.True(t, IsValidToken(token))

	assert.False(t, IsValidToken("invalidtoken"))
}

func TestGetSessionData(t *testing.T) {
	config.SetConfig(config.Config{
		Auth: config.AuthConfig{
			Secret: "testsecret",
		},
	})

	token, err := GenerateToken("testuser", 1)
	require.NoError(t, err)

	data, err := GetSessionData(token)
	require.NoError(t, err)
	assert.Equal(t, 1, data.ID)
	assert.Equal(t, "testuser", data.Username)

	_, err = GetSessionData("invalidtoken")
	assert.Error(t, err)
}
