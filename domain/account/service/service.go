package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/south"
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

func (s *Service) CreateAccount(ctx context.Context, account *entity.AccountEntity) error {
	return nil
}

func (s *Service) GetAccount(ctx context.Context, vo *entity.AccountVo) (*entity.AccountEntity, error) {
	return nil, nil
}

func (s *Service) Transaction(ctx context.Context, payer *entity.AccountEntity, payee *entity.AccountEntity, tx *entity.TransactionVo) (*entity.TransactionEntity, error) {
	return nil, nil
}
