package repository

import (
	"context"

	"github.com/arvinpaundra/go-boilerplate/domain/auth/entity"
)

type UserWriter interface {
	Save(ctx context.Context, user entity.User) error
}
