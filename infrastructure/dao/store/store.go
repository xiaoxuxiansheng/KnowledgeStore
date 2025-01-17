package store

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/south"
)

type Dao struct{}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateStore(ctx context.Context, entity *entity.StoreEntity) error {
	return nil
}

func (d *Dao) GetStores(ctx context.Context, opts ...north.QueryOption) ([]*entity.StoreEntity, error) {
	return nil, nil
}
