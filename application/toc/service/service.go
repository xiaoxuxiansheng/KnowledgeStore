package service

import (
	"context"
	"fmt"

	"github.com/xiaoxuxiansheng/KnowledgeStore/application/toc/dto"
	accountnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/north"
	goodsnorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/north"
	ordernorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/order/north"
	storenorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/north"
	usernorth "github.com/xiaoxuxiansheng/KnowledgeStore/domain/user/north"
	"github.com/xiaoxuxiansheng/KnowledgeStore/presentation/iservice"

	"go.uber.org/dig"
)

type Service struct {
	userService    usernorth.IService
	accountService accountnorth.IService
	storeService   storenorth.IService
	goodsService   goodsnorth.IService
	orderService   ordernorth.IService
}

func NewService(param ServiceParam) iservice.ToCService {
	return &Service{
		userService:    param.UserService,
		accountService: param.AccountService,
		storeService:   param.StoreService,
		goodsService:   param.GoodsService,
		orderService:   param.OrderService,
	}
}

type ServiceParam struct {
	dig.In
	UserService    usernorth.IService
	AccountService accountnorth.IService
	StoreService   storenorth.IService
	GoodsService   goodsnorth.IService
	OrderService   ordernorth.IService
}

/**
 * @brief: 注册用户
 * @param: dto.RegisterReqDto——注册请求参数
 * @return: dto.RegisterRespDto——注册响应参数
 * 1) 调用 user domain 创建用户
 * 2）调用 account domain 创建用户账户
 */
func (s *Service) RegisterCustomer(ctx context.Context, req *dto.RegisterCutomerReqDto) error {
	if err := s.userService.RegisterUser(ctx, req.ToUserEntity()); err != nil {
		return err
	}
	return s.accountService.CreateAccount(ctx, req.ToAccountEntity())
}

/**
 * @brief: 用户登录
 * @param: dto.LoginReqDto——登录请求参数
 * @return: dto.LoginRespDto——登录响应参数
 * 1) 调用 account domain 校验用户是否合法
 */
func (s *Service) LoginCustomer(ctx context.Context, req *dto.LoginCutomerReqDto) error {
	_, err := s.accountService.GetAccount(ctx, req.ToAccountVo())
	return err
}

/**
 * @brief: 检索商铺
 * @param: dto.GetStoresReqDto——检索店铺请求参数
 * @return: dto.GetStoresRespDto——检索店铺响应参数
 * 1) 调用 store domain 检索在线的店铺
 */
func (s *Service) ListStore(ctx context.Context, req *dto.ListStoresReqDto) (*dto.ListStoresRespDto, error) {
	entities, total, err := s.storeService.GetStores(ctx, req.ToQueryOptions()...)
	if err != nil {
		return nil, err
	}
	return dto.NewListStoresRespDto(total, entities), nil
}

/**
 * @brief: 检索店铺内的在线商品
 * @param: dto.ListGoodsReqDto——检索店铺请求参数
 * @return: dto.ListGoodsRespDt——检索店铺响应参数
 * 1) 调用 goods domain 检索商品
 */
func (s *Service) ListGoods(ctx context.Context, req *dto.ListGoodsReqDto) (*dto.ListGoodsRespDto, error) {
	entities, total, err := s.goodsService.GetGoods(ctx, req.ToQueryOptions()...)
	if err != nil {
		return nil, err
	}
	return dto.NewListGoodsRespDto(total, entities), nil
}

/**
 * @brief: 查询商品详情
 * @param: dto.QueryGoodsReqDto——查询商品详情请求参数
 * @return: dto.QueryGoodsRespDto——查询商品详情响应参数
 * 1) 调用 good domain 检索商品
 */
func (s *Service) QueryGoods(ctx context.Context, req *dto.QueryGoodsReqDto) (*dto.QueryGoodsRespDto, error) {
	entities, _, err := s.goodsService.GetGoods(ctx, req.ToQueryOptions()...)
	if err != nil || len(entities) != 1 {
		return nil, fmt.Errorf("invalid good name: %s, store cnt: %d, err: %v", req.GoodsName, len(entities), err)
	}
	return dto.NewQueryGoodsRespDto(entities[0]), nil
}

/**
 * @brief: 购买商品
 * @param: dto.BuyGoodsReqDto——购买商品请求参数
 * 1) 调用 goods domain 校验商品合法性
 * 2) 调用 order domain 创建订单
 */
func (s *Service) BuyGoods(ctx context.Context, req *dto.BuyGoodsReqDto) error {
	goods, _, err := s.goodsService.GetGoods(ctx, req.ToQueryGoodsOption()...)
	if err != nil || len(goods) != 1 {
		return fmt.Errorf("invalid good name: %s, store cnt: %d, err: %v", req.GoodsName, len(goods), err)
	}

	_, err = s.orderService.CreateOrder(ctx, req.ToOrderVo(goods[0].Points))
	return err
}

/**
 * @brief: 查看订单列表
 * @param: dto.ListOrderReqDto——查看订单列表请求参数
 * @return: dto.ListOrderRespDto——查看订单列表响应参数
 * 1）调用 order domain 检索订单
 */
func (s *Service) ListOrder(ctx context.Context, req *dto.ListOrderReqDto) (*dto.ListOrderRespDto, error) {
	entities, total, err := s.orderService.GetOrders(ctx, req.ToQueryOptions()...)
	if err != nil {
		return nil, err
	}
	return dto.NewListOrderRespDto(total, entities), nil
}

/**
 * @brief: 支付订单
 * @param: dto.PayOrderReqDto——支付订单请求参数
 * @return: dto.PayOrderRespDto——支付订单响应参数
 * 1）调用 order domain 锁定订单
 * 2）调用 account domain 转移账户点数
 * 3) 调用 order domain 更新订单状态
 */
func (s *Service) PayOrder(ctx context.Context, req *dto.PayOrderReqDto) error {
	order, err := s.orderService.LockOrder(ctx, req.ToOrderEntity())
	if err != nil {
		return err
	}
	payer, payee := req.ToAccountEntities(order)
	if _, err = s.accountService.Transaction(ctx, payer, payee); err != nil {
		return s.orderService.FinishOrder(ctx, order)
	}
	return s.orderService.UnlockOrder(ctx, order)
}
