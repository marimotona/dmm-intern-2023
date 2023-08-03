package dao

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"yatter-backend-go/app/domain/object"
	"yatter-backend-go/app/domain/repository"

	"github.com/jmoiron/sqlx"
)

type (
	// Implementation for repository.Account
	account struct {
		db *sqlx.DB
	}

	// handler struct {
	// 	ar repository.Account
	// }
)

// Create accout repository
func NewAccount(db *sqlx.DB) repository.Account {
	return &account{db: db}
}

// Create HTTP request Handler
// func NewHandler(ar repository.Account) *handler {
// 	return &handler{ar: ar}
// }

// FindByUsername : ユーザ名からユーザを取得
func (r *account) FindByUsername(ctx context.Context, username string) (*object.Account, error) {
	entity := new(object.Account)
	err := r.db.QueryRowxContext(ctx, "select * from account where username = ?", username).StructScan(entity)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, fmt.Errorf("failed to find account from db: %w", err)
	}

	return entity, nil
}

func (r *account) CreateUser(ctx context.Context, a *object.Account) error {
	query := `
		INSERT INTO account (username, password_hash)
		VALUES (?, ?)
	`

	_, err := r.db.ExecContext(ctx, query, a.Username, a.PasswordHash)
	if err != nil {
		return fmt.Errorf("faild to insert account into db: %w", err)
	}
	return nil
}

func (r *account) GetUserAccounts(ctx context.Context, username string) ([]*object.Account, error) {
	query := `SELECT * FROM account WHERE username = ?`

	rows, err := r.db.QueryxContext(ctx, query, username)
	if err != nil {
		return nil, fmt.Errorf("faild to get user accounts: %w", err)
	}
	defer rows.Close()

	var accounts []*object.Account

	for rows.Next() {
		account := new(object.Account)
		err := rows.StructScan(account)

		if err != nil {
			return nil, fmt.Errorf("faild to scan: %w", err)
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}
