package dao

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	media struct {
		db *sqlx.DB
	}
)

func NewMedia(db *sqlx.DB) repository.Media {
	return &media{db: db}
}

func (r *media) AddMedia(ctx context.Context, m *object.Media) (int64, error) {
	if m.MediaURL == "" {
		return 0, errors.New("none media_url")
	}

	query := "INSERT INTO attachment (media_url, create_at) VALUES (?, ?)"
	result, err := r.db.ExecContext(ctx, query, m.MediaURL, time.Now())
	if err != nil {
		return 0, fmt.Errorf("faild insert new media: %w", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("faild get last insert id: %w", err)
	}

	return id, nil
}
