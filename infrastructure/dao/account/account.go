package account

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/south"
)

type Dao struct{}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateAccount(ctx context.Context, userEntity *entity.AccountEntity) error {
	return nil
}

func (d *Dao) GetAccount(ctx context.Context, vo *entity.AccountVo) (*entity.AccountEntity, error) {
	return nil, nil
}

func (d *Dao)Transaction(ctx context.Context, payer *entity.AccountEntity, payee *entity.AccountEntity, tx *entity.TransactionEntity) error{
	return nil 
}
