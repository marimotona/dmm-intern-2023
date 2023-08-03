package dao

import (
	"context"
	"database/sql"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"
)

type (
	status struct {
		db *sql.DB
	}
)

func NewStatus(db *sql.DB) repository.Status {
	return &status{db: db}
}

func (r *status) AddStatus(ctx context.Context, a *object.Account, s *object.Status) (int64, error) {
	return a.ID, nil
}
