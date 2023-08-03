package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
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
	if a == nil || a.ID == 0 {
		return 0, errors.New("no account")
	}

	if s.Content == "" {
		return 0, errors.New("content is empty")
	}

	query := "INSERT INTO status (account_id, content, create_at) VALUES (?, ?, ?)"
	result, err := r.db.ExecContext(ctx, query, a.ID, s.Content, time.Now())
	if err != nil {
		return 0, fmt.Errorf("faild insert new status: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("faild get last insert id: %w", err)
	}
	return id, nil
}
