package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
)

type IService interface {
	// 检索商品
	GetGoods(ctx context.Context, opts ...QueryOption) ([]*entity.GoodsEntity, error)
	// 创建商品
	CreateGoods(ctx context.Context, vo *entity.GoodsVo) (*entity.GoodsEntity, error)
	// 发布商品
	OnlineGoods(ctx context.Context, entity *entity.GoodsEntity) error
	// 下线商品
	OfflineGoods(ctx context.Context, entity *entity.GoodsEntity) error
	// 购买商品
	PurchaseGoods(ctx context.Context, entity *entity.GoodsEntity) error
}

type QueryOptions struct{}

type QueryOption func(*QueryOptions)
