package repository

import (
	"context"

	"github.com/arvinpaundra/go-boilerplate/domain/auth/entity"
)

type UserReader interface {
	FindByEmail(ctx context.Context, email string) (entity.User, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}
