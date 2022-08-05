package services

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/google/uuid"
	"github.com/muhammadarash1997/go-kit-http/entities"
	"github.com/muhammadarash1997/go-kit-http/repositories"
)

type Service interface {
	CreateUser(ctx context.Context, email string, password string) (string, error)
	GetUser(ctx context.Context, id string) (string, error)
}

type service struct {
	repository repositories.Repository
	logger     log.Logger
}

func NewService(repository repositories.Repository, logger log.Logger) Service {
	return &service{repository, logger}
}

func (s *service) CreateUser(ctx context.Context, email string, password string) (string, error) {
	logger := log.With(s.logger, "method", "CreateUser")

	id := uuid.New().String()

	user := entities.User{
		ID:       id,
		Email:    email,
		Password: password,
	}

	err := s.repository.CreateUser(ctx, user)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Create User", id)

	return "Success", nil
}

func (s *service) GetUser(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetUser")

	email, err := s.repository.GetUser(ctx, id)
	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get User", id)

	return email, nil
}
