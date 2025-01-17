package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/south"
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

func (s *Service) Register(ctx context.Context) (*entity.MerchantEntity, error) {
	return nil, nil
}

func (s *Service) GetMerchant(ctx context.Context, merchantId string) (*entity.MerchantEntity, error) {
	return nil, nil
}
