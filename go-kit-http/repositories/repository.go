package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/muhammadarash1997/go-kit-http/entities"
)

var RepositoryError = errors.New("Unable to handle Repo Request")

type Repository interface {
	CreateUser(ctx context.Context, user entities.User) error
	GetUser(ctx context.Context, id string) (string, error)
}

type repository struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepository(db *sql.DB, logger log.Logger) Repository {
	return &repository{db, logger}
}

func (r *repository) CreateUser(ctx context.Context, user entities.User) error {
	sql := `
		INSERT INTO users (id, email, password)
		VALUES ($1, $2, $3)
	`

	if user.Email != "" || user.Password != "" {
		return RepositoryError
	}

	_, err := r.db.ExecContext(ctx, sql, user.ID, user.Email, user.Password)
	if err != nil {
		return RepositoryError
	}

	return nil
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	var email string
	err := r.db.QueryRow("SELECT email FROM users WHERE id = $1", id).Scan(&email)
	if err != nil {
		return "", RepositoryError
	}

	return email, nil
}
