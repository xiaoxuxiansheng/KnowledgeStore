package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/entity"
)

type IService interface {
	CreateOrder(ctx context.Context, vo *entity.OrderVo) (*entity.OrderEntity, error)
	LockOrder(ctx context.Context, req *entity.OrderEntity) (*entity.OrderEntity, error)
	UnlockOrder(ctx context.Context, req *entity.OrderEntity) error
	FinishOrder(ctx context.Context, entity *entity.OrderEntity) error
	GetOrders(ctx context.Context, opts ...QueryOption) ([]*entity.OrderEntity, int, error)
}

type QueryOptions struct {
	UserId string
	Status string
	Page   int
	Limit  int
}

type QueryOption func(*QueryOptions)

func WithUserId(userId string) QueryOption {
	return func(qo *QueryOptions) {
		qo.UserId = userId
	}
}

func WithStatus(status string) QueryOption {
	return func(qo *QueryOptions) {
		qo.Status = status
	}
}

func WithPageLimit(page, limit int) QueryOption {
	return func(qo *QueryOptions) {
		qo.Page = page
		qo.Limit = limit
	}
}
