package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/south"
	"go.uber.org/dig"
)

type Service struct {
	repo south.Repository
}

type ServiceParam struct {
	dig.In
	Repo south.Repository
}

func NewService(param ServiceParam) north.IService {
	return &Service{
		repo: param.Repo,
	}
}

func (s *Service) RegisterUser(ctx context.Context, req *entity.UserEntity) error {
	return nil
}

func (s *Service) GetUser(ctx context.Context, userId string) (*entity.UserEntity, error) {
	return nil, nil
}
