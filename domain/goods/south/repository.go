package south

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
)

/**
 * @brief: store 仓储服务
 */
type Repository interface {
	CreateGood(ctx context.Context, entity *entity.GoodsEntity) error

	GetGoods(ctx context.Context, opts ...north.QueryOption) ([]*entity.GoodsEntity, error)

	UpdateGoods(ctx context.Context, entity *entity.GoodsEntity) error
}
