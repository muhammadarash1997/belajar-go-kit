package services

import (
	"context"

	"github.com/go-kit/kit/log"
)

type Service interface {
	Add(ctx context.Context, numA, numB float32) (float32, error)
}

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) Service {
	return &service{
		logger: logger,
	}
}

func (s *service) Add(ctx context.Context, numA, numB float32) (float32, error) {
	return numA + numB, nil
}
