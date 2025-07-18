package internal

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator wraps the validator
type CustomValidator struct {
	validator *validator.Validate
}

// NewValidator creates a new custom validator instance
func NewValidator() *CustomValidator {
	return &CustomValidator{
		validator: validator.New(),
	}
}

// Validate validates the struct
func (cv *CustomValidator) Validate(i any) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
