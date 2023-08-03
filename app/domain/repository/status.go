package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Status interface {
	AddStatus(ctx context.Context, a *object.Account, s *object.Status) (int64, error)
	FindByID(ctx context.Context, id int64) (*object.Status, error)
}
