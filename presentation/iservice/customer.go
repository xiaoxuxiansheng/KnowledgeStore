package iservice

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/application/toc/dto"
)

type ToCService interface {

	/**
	 * @brief: 注册用户
	 * @param: dto.RegisterReqDto——注册请求参数
	 * @return: dto.RegisterRespDto——注册响应参数
	 * 1) 调用 user domain 完成用户注册
	 * 2）调用 account domain 创建用户账户
	 */
	RegisterCustomer(ctx context.Context, req *dto.RegisterCutomerReqDto) error

	/**
	 * @brief: 用户登录
	 * @param: dto.LoginReqDto——登录请求参数
	 * @return: dto.LoginRespDto——登录响应参数
	 * 1) 调用 account domain 校验用户是否合法
	 */
	LoginCustomer(ctx context.Context, req *dto.LoginCutomerReqDto) error

	/**
	 * @brief: 检索商铺
	 * @param: dto.GetStoresReqDto——检索店铺请求参数
	 * @return: dto.GetStoresRespDto——检索店铺响应参数
	 * 1) 调用 store domain 检索店铺
	 */
	ListStore(ctx context.Context, req *dto.ListStoresReqDto) (*dto.ListStoresRespDto, error)

	/**
	 * @brief: 检索商品
	 * @param: dto.GetGoodsReqDto——检索店铺请求参数
	 * @return: dto.GetGoodsRespDto——检索店铺响应参数
	 * 1) 调用 good domain 检索商品
	 */
	ListGoods(ctx context.Context, req *dto.ListGoodsReqDto) (*dto.ListGoodsRespDto, error)

	QueryGoods(ctx context.Context, req *dto.QueryGoodsReqDto) (*dto.QueryGoodsRespDto, error)
	/**
	 * @brief: 购买商品
	 * @param: dto.PurchaseGoodsReqDto——购买商品请求参数
	 * @return: dto.PurchaseGoodsRespDto——购买商品响应参数
	 * 1) 调用 good domain 锁定商品
	 * 2) 调用 order domain 创建订单
	 * 3) 调用 good domain 扣减商品
	 */
	BuyGoods(ctx context.Context, req *dto.BuyGoodsReqDto) error

	ListOrder(ctx context.Context, req *dto.ListOrderReqDto) (*dto.ListOrderRespDto, error)

	/**
	 * @brief: 支付订单
	 * @param: dto.PayOrderReqDto——支付订单请求参数
	 * @return: dto.PayOrderRespDto——支付订单响应参数
	 * 1）调用 order domain 锁定订单
	 * 2）调用 account domain 转移账户点数
	 * 3) 调用 order domain 更新订单成功
	 */
	PayOrder(ctx context.Context, req *dto.PayOrderReqDto) error
}
