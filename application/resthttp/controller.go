package resthttp

import (
	"github.com/arvinpaundra/go-boilerplate/core/validator"
	"gorm.io/gorm"
)

type Controller struct {
	db        *gorm.DB
	validator *validator.Validator
}

func NewController(db *gorm.DB, validator *validator.Validator) *Controller {
	return &Controller{
		db:        db,
		validator: validator,
	}
}
