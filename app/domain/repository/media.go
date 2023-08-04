package repository

import (
	"context"
	"yatter-backend-go/app/domain/object"
)

type Media interface {
	AddMedia(ctx context.Context, m *object.Media) (int64, error)
}
