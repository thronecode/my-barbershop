package utils

import (
	"github.com/thronecode/my-barbershop/backend/internal/config"

	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// CheckPassword checks if a password matches a hashed password
func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// HashPassword hashes a password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// GenerateToken generates a JWT token
func GenerateToken(username string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"id":       id,
		"exp":      time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(config.GetConfig().Auth.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// IsValidToken validates a JWT token
func IsValidToken(tokenString string) bool {
	token, err := decodeToken(tokenString)
	if err != nil {
		return false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		exp := claims["exp"].(float64)
		return exp > float64(time.Now().Unix())
	}

	return false
}

// SessionData is the data stored in a JWT token
type SessionData struct {
	ID       int
	Username string
}

// GetSessionData retrieves the username from a JWT token
func GetSessionData(tokenString string) (SessionData, error) {
	data := SessionData{}

	token, err := decodeToken(tokenString)
	if err != nil {
		return data, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		data.ID = int(claims["id"].(float64))
		data.Username = claims["username"].(string)
	}

	return data, nil
}

func decodeToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("invalid signing method", jwt.ValidationErrorSignatureInvalid)
		}
		return []byte(config.GetConfig().Auth.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}
