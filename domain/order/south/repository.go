package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/north"
)

type Repository interface {
	CreateOrder(ctx context.Context, vo *entity.OrderVo) error
	UpdateOrder(ctx context.Context, vo *entity.OrderVo) error
	GetOrders(ctx context.Context, opts ...north.QueryOption) ([]*entity.OrderEntity, error)
}
