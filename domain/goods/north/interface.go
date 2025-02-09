package north

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
)

type IService interface {
	// 检索商品
	GetGoods(ctx context.Context, opts ...QueryOption) ([]*entity.GoodsEntity, int, error)
	// 创建商品
	CreateGoods(ctx context.Context, req *entity.GoodsEntity) error
	// 发布商品
	OnlineGoods(ctx context.Context, req *entity.GoodsEntity) error
	// 下线商品
	OfflineGoods(ctx context.Context, req *entity.GoodsEntity) error
	// 购买商品
	PurchaseGoods(ctx context.Context, req *entity.GoodsEntity) error
}

type QueryOptions struct {
	Category  string
	UserId    string
	StoreName string
	GoodsName string
	Page      int
	Limit     int
}

type QueryOption func(*QueryOptions)

func WithUserId(userId string) QueryOption {
	return func(qo *QueryOptions) {
		qo.UserId = userId
	}
}

func WithStoreName(storeName string) QueryOption {
	return func(qo *QueryOptions) {
		qo.StoreName = storeName
	}
}

func WithGoodsName(goodsName string) QueryOption {
	return func(qo *QueryOptions) {
		qo.GoodsName = goodsName
	}
}

func WithPageLimit(page, limit int) QueryOption {
	return func(qo *QueryOptions) {
		qo.Page = page
		qo.Limit = limit
	}
}

func WithCategory(category string) QueryOption {
	return func(qo *QueryOptions) {
		qo.Category = category
	}
}
