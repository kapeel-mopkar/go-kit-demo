package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
	"github.com/kapeel-mopkar/go-kit-demo/account/db/model"
)

var RepoErr = errors.New("Unable to handle repository errors")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) model.Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreateUser(ctx context.Context, user model.User) error {
	query := `INSERT into users (id, email, password) 
	VALUES (?, ?, ?)`

	if user.Email == "" || user.Password == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, query, user.ID, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (repo *repo) GetUser(ctc context.Context, id string) (string, error) {
	var email string
	err := repo.db.QueryRow("SELECT email from users WHERE id=?", id).Scan(&email)
	if err != nil {
		return "", RepoErr
	}
	return email, nil
}
