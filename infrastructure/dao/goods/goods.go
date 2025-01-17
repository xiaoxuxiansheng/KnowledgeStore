package goods

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/south"
)

type Dao struct{}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateGood(ctx context.Context, entity *entity.GoodsEntity) error {
	return nil
}

func (d *Dao) GetGoods(ctx context.Context, opts ...north.QueryOption) ([]*entity.GoodsEntity, error) {
	return nil, nil
}

func (d *Dao) UpdateGoods(ctx context.Context, entity *entity.GoodsEntity) error {
	return nil
}
