package repository

import (
	"context"

	"github.com/arvinpaundra/go-boilerplate/domain/auth/entity"
)

type SessionReader interface {
	FindByRefreshToken(ctx context.Context, userId int64, refreshToken string) (entity.Session, error)
}
