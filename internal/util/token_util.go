package util

import (
	"time"
	"toko-api-fiber/internal/exception"
	"toko-api-fiber/internal/model"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenUtil interface {
	CreateToken(auth *model.Auth) (string, error)
	ParseToken(jwtToken string) (*model.Auth, error)
}

type TokenUtilImpl struct {
	SecretKey string
}

func NewTokenUtil(config *viper.Viper) TokenUtil {
	secretKey := config.GetString("JWT_SECRET_KEY")
	return &TokenUtilImpl{
		SecretKey: secretKey,
	}
}

func (t *TokenUtilImpl) CreateToken(auth *model.Auth) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       auth.ID,
		"username": auth.Username,
		"email":    auth.Email,
		"exp":      time.Now().Add(24 * time.Hour * 30).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(t.SecretKey))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}

func (t *TokenUtilImpl) ParseToken(jwtToken string) (*model.Auth, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (any, error) {
		return []byte(t.SecretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claim := token.Claims.(jwt.MapClaims)

	expFloat := claim["exp"].(float64)
	exp := int64(expFloat)

	if exp < time.Now().Unix() {
		return nil, exception.ErrUnauthorized
	}

	idFloat := claim["id"].(float64)
	id := int64(idFloat)

	username := claim["username"].(string)
	email := claim["email"].(string)

	auth := &model.Auth{
		ID:       id,
		Username: username,
		Email:    email,
	}

	return auth, nil
}
