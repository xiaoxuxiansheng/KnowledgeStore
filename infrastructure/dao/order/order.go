package order

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/south"
)

type Dao struct{}

func NewDao() south.Repository {
	return &Dao{}
}

func (d *Dao) CreateOrder(ctx context.Context, vo *entity.OrderVo) error {
	return nil
}

func (d *Dao) UpdateOrder(ctx context.Context, vo *entity.OrderVo) error {
	return nil
}

func (d *Dao) GetOrders(ctx context.Context, opts ...north.QueryOption) ([]*entity.OrderEntity, error) {
	return nil, nil
}
