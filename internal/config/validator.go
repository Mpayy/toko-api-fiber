package config

import (
	"toko-api-fiber/internal/exception"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func NewValidator(config *viper.Viper) *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("valid_price", exception.ValidPrice)
	return validate
}
