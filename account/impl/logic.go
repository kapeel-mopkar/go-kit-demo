package impl

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
	"github.com/kapeel-mopkar/go-kit-demo/account"
	"github.com/kapeel-mopkar/go-kit-demo/account/db/model"
)

type service struct {
	repository model.Repository
	logger     log.Logger
}

func NewService(repo model.Repository, log log.Logger) account.Service {
	return &service{
		repository: repo,
		logger:     log,
	}
}

func (s service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	uuid, _ := uuid.NewV4()
	id := uuid.String()

	user := model.User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	if err := s.repository.CreateUser(ctx, user); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Created User with id - ", id)

	return "Success", nil

}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")
	email, err := s.repository.GetUser(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Retrieved User with id - ", id)

	return email, nil

}
