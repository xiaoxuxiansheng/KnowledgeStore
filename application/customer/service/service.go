package service

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/application/customer/dto"
	"github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"
	"go.uber.org/dig"
)

type Service struct {
	userService    north.IService
	accountService north.IService
	storeService   north.IService
	goodsService   north.IService
}

func NewService(param ServiceParam) iservice.CustomerService {
	return &Service{
		userService:    param.UserService,
		accountService: param.AccountService,
		storeService:   param.StoreService,
		goodsService:   param.GoodsService,
	}
}

type ServiceParam struct {
	dig.In
	UserService    north.IService
	AccountService north.IService
	StoreService   north.IService
	GoodsService   north.IService
}

/**
 * @brief: 注册用户
 * @param: dto.RegisterReqDto——注册请求参数
 * @return: dto.RegisterRespDto——注册响应参数
 * 1) 调用 user domain 完成用户注册
 * 2）调用 account domain 创建用户账户
 */
func (s *Service) Register(ctx context.Context, req *dto.RegisterReqDto) (*dto.RegisterRespDto, error) {
	return nil, nil
}

/**
 * @brief: 用户登录
 * @param: dto.LoginReqDto——登录请求参数
 * @return: dto.LoginRespDto——登录响应参数
 * 1) 调用 account domain 校验用户是否合法
 */
func (s *Service) Login(ctx context.Context, req *dto.LoginReqDto) (*dto.LoginRespDto, error) {
	return nil, nil
}

/**
 * @brief: 检索商铺
 * @param: dto.GetStoresReqDto——检索店铺请求参数
 * @return: dto.GetStoresRespDto——检索店铺响应参数
 * 1) 调用 store domain 检索店铺
 */
func (s *Service) GetStores(ctx context.Context, req *dto.GetStoresReqDto) (*dto.GetStoresRespDto, error) {
	return nil, nil
}

/**
 * @brief: 检索商品
 * @param: dto.GetGoodsReqDto——检索店铺请求参数
 * @return: dto.GetGoodsRespDto——检索店铺响应参数
 * 1) 调用 good domain 检索商品
 */
func (s *Service) GetGoods(ctx context.Context, req *dto.GetGoodsReqDto) (*dto.GetGoodsRespDto, error) {
	return nil, nil
}

/**
 * @brief: 购买商品
 * @param: dto.PurchaseGoodsReqDto——购买商品请求参数
 * @return: dto.PurchaseGoodsRespDto——购买商品响应参数
 * 1) 调用 good domain 锁定商品
 * 2) 调用 order domain 创建订单
 * 3) 调用 good domain 扣减商品
 */
func (s *Service) PurchaseGoods(ctx context.Context, req *dto.PurchaseGoodsReqDto) (*dto.PurchaseGoodsRespDto, error) {
	return nil, nil
}

/**
 * @brief: 支付订单
 * @param: dto.PayOrderReqDto——支付订单请求参数
 * @return: dto.PayOrderRespDto——支付订单响应参数
 * 1）调用 order domain 锁定订单
 * 2）调用 account domain 转移账户点数
 * 3) 调用 order domain 更新订单成功
 */
func (s *Service) PayOrder(ctx context.Context, req *dto.PayOrderReqDto) (*dto.PayOrderRespDto, error) {
	return nil, nil
}
