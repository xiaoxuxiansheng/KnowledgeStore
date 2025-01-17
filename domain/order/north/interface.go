package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
)

type IService interface {
	CreateOrder(ctx context.Context, vo *entity.OrderVo) (*entity.OrderEntity, error)
	PayOrder(ctx context.Context, entity *entity.OrderEntity) error
	GetOrders(ctx context.Context, opts ...QueryOption) ([]*entity.OrderEntity, error)
}

type QueryOptions struct{}

type QueryOption func(*QueryOptions)
