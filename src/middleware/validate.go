package middleware

import (
	validator "gopkg.in/go-playground/validator.v9"
)

// DefaultValidator 默认验证器
type DefaultValidator struct {
	Validator *validator.Validate
}

// Validate 参数验证
func (cv *DefaultValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}
