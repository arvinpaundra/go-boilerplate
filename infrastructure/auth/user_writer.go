package auth

import (
	"context"

	"github.com/arvinpaundra/go-boilerplate/domain/auth/entity"
	"github.com/arvinpaundra/go-boilerplate/domain/auth/repository"
	"github.com/arvinpaundra/go-boilerplate/model"
	"gorm.io/gorm"
)

var _ repository.UserWriter = (*UserWriterRepository)(nil)

type UserWriterRepository struct {
	db *gorm.DB
}

func NewUserWriterRepository(db *gorm.DB) UserWriterRepository {
	return UserWriterRepository{db: db}
}

func (r UserWriterRepository) Save(ctx context.Context, user entity.User) error {
	userModel := user.ToModel()

	err := r.db.WithContext(ctx).
		Model(&model.User{}).
		Create(&userModel).
		Error

	if err != nil {
		return err
	}

	return nil
}
