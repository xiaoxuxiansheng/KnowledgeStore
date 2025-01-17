package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/south"
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

func (s *Service) CreateOrder(ctx context.Context, vo *entity.OrderVo) (*entity.OrderEntity, error) {
	return nil, nil
}

func (s *Service) PayOrder(ctx context.Context, entity *entity.OrderEntity) error {
	return nil
}

func (s *Service) GetOrders(ctx context.Context, opts ...north.QueryOption) ([]*entity.OrderEntity, error) {
	return nil, nil
}
