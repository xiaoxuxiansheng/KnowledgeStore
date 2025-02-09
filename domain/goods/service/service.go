package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/south"
	"go.uber.org/dig"
)

type Service struct {
	repo south.Repository
}

func NewService(param ServiceParam) north.IService {
	return &Service{
		repo: param.Repo,
	}
}

type ServiceParam struct {
	dig.In
	Repo south.Repository
}

// 检索商品
func (s *Service) GetGoods(ctx context.Context, opts ...north.QueryOption) ([]*entity.GoodsEntity, int, error) {
	return nil, 0, nil
}

// 创建商品
func (s *Service) CreateGoods(ctx context.Context, req *entity.GoodsEntity) error {
	return nil
}

// 发布商品
func (s *Service) OnlineGoods(ctx context.Context, req *entity.GoodsEntity) error {
	return nil
}

// 下线商品
func (s *Service) OfflineGoods(ctx context.Context, req *entity.GoodsEntity) error {
	return nil
}

// 购买商品
func (s *Service) PurchaseGoods(ctx context.Context, req *entity.GoodsEntity) error {
	return nil
}
