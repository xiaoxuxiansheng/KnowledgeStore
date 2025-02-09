package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/south"
	"go.uber.org/dig"
)

type Service struct {
	repo south.Repository
}

func NewService(param ServiceParam) north.IService {
	return &Service{
		repo: param.Repo,
	}
}

type ServiceParam struct {
	dig.In
	Repo south.Repository
}

func (s *Service) CreateStore(ctx context.Context, req *entity.StoreEntity) error {
	return nil
}

func (s *Service) OnlineStore(ctx context.Context, req *entity.StoreEntity) error {
	return nil
}

func (s *Service) OfflineStore(ctx context.Context, req *entity.StoreEntity) error {
	return nil
}

func (s *Service) GetStores(ctx context.Context, opts ...north.QueryOption) ([]*entity.StoreEntity, int, error) {
	return nil, 0, nil
}
