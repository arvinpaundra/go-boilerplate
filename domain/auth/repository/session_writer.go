package repository

import (
	"context"

	"github.com/arvinpaundra/go-boilerplate/domain/auth/entity"
)

type SessionWriter interface {
	Save(ctx context.Context, refreshToken entity.Session) error
}
