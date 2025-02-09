package service

import (
	"context"
	"fmt"

	"github.com/xiaoxuxiansheng/KnowledgeStore/application/tob/dto"
	accountnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/north"
	goodsnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
	merchantnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/north"
	storenorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"

	"go.uber.org/dig"
)

type Service struct {
	merchantService merchantnorth.IService
	accountService  accountnorth.IService
	storeService    storenorth.IService
	goodsService    goodsnorth.IService
}

func NewService(param ServiceParam) iservice.ToBService {
	return &Service{
		merchantService: param.MerchantService,
		accountService:  param.AccountService,
		storeService:    param.StoreService,
		goodsService:    param.GoodsService,
	}
}

type ServiceParam struct {
	dig.In
	MerchantService merchantnorth.IService
	AccountService  accountnorth.IService
	StoreService    storenorth.IService
	GoodsService    goodsnorth.IService
}

/**
 * @brief: 注册商户
 * @param: dto.RegisterReqDto——注册请求参数
 * 1) 调用 merchant domain 完成商户实例创建
 * 2）调用 account domain 创建商户账户实例
 */
func (s *Service) RegisterBusiness(ctx context.Context, req *dto.RegisterBusinessReqDto) error {
	// 1) 注册 business 实例
	if err := s.merchantService.RegisterMerchant(ctx, req.ToMerchantEntity()); err != nil {
		return err
	}

	// 2) 创建 business 账户
	return s.accountService.CreateAccount(ctx, req.ToAccountEntity())
}

/**
 * @brief: 登录商户
 * @param: dto.LoginReqDto——登录请求参数
 * 1）调用 account domain 验证商户账号和密码是否合法
 */
func (s *Service) LoginBusiness(ctx context.Context, req *dto.LoginBusinessReqDto) error {
	// 1) 校验 business 账号是否合法
	_, err := s.accountService.GetAccount(ctx, req.ToAccountVo())
	return err
}

/**
 * @brief: 创建店铺
 * @param: dto.CreateStoreReqDto——创建商铺请求参数
 * 1) 调用 merchant domain 校验商户账号是否合法
 * 2) 调用 store domain 创建店铺
 */
func (s *Service) CreateStore(ctx context.Context, req *dto.CreateStoreReqDto) error {
	// 1) 查询商家信息
	if _, err := s.merchantService.QueryMerchant(ctx, req.AccountId); err != nil {
		return err
	}
	// 2) 创建商铺信息
	return s.storeService.CreateStore(ctx, req.ToStoreEntity())
}

/**
 * @brief: 开放店铺
 * @param: dto.OnlineStoreReqDto——开放商铺请求参数
 * 1) 调用 store domain 获取并锁定店铺
 * 2) 调用 store domain 更新店铺状态为上线
 */
func (s *Service) OnlineStore(ctx context.Context, req *dto.OnlineStoreReqDto) error {
	return s.storeService.OnlineStore(ctx, req.ToStoreEntity())
}

/**
 * @brief: 下线店铺
 * @param: dto.OfflineStoreReqDto——下线商铺请求参数
 * 1) 调用 store domain 获取并锁定店铺
 * 2) 调用 store domain 更新店铺状态为下线
 */
func (s *Service) OfflineStore(ctx context.Context, req *dto.OfflineStoreReqDto) error {
	return s.storeService.OfflineStore(ctx, req.ToStoreEntity())
}

/**
 * @brief: 创建商品
 * @param: dto.CreateGoodsReqDto——创建商品请求参数
 * 1) 调用 store domain 查询店铺校验合法性
 * 2) 调用 goods domain 创建商品
 */
func (s *Service) CreateGoods(ctx context.Context, req *dto.CreateGoodsReqDto) error {
	stores, _, err := s.storeService.GetStores(ctx, storenorth.WithStoreName(req.StoreName), storenorth.WithPageLimit(1, 1))
	if err != nil || len(stores) != 1 {
		return fmt.Errorf("invalid store name: %s, store cnt: %d, err: %v", req.StoreName, len(stores), err)
	}
	return s.goodsService.CreateGoods(ctx, req.ToGoodsEntity())
}

/**
 * @brief: 发布商品
 * @param: dto.OnlineGoodsReq——发布商品请求参数
 * 1) 调用 goods domain 获取并锁定商品
 * 2) 调用 goods domain 更新商品为发布状态
 */
func (s *Service) OnlineGoods(ctx context.Context, req *dto.OnlineGoodsReqDto) error {
	return s.goodsService.OnlineGoods(ctx, req.ToGoodsEntity())
}

/**
 * @brief: 下线商品
 * @param: dto.OfflineGoodsReq——下线商品请求参数
 * 1) 调用 goods domain 获取并锁定商品
 * 2) 调用 goods domain 更新商品为下线状态
 */
func (s *Service) OfflineGoods(ctx context.Context, req *dto.OfflineGoodsReqDto) error {
	return s.goodsService.OfflineGoods(ctx, req.ToGoodsEntity())
}
