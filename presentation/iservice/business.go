package iservice

import (
	"context"

	"github.com/xiaoxuxiansheng/KnowledgeStore/application/business/dto"
)

type BusinessService interface {
	/**
	 * @brief: 注册商户
	 * @param: dto.RegisterReqDto——注册请求参数
	 * @return: dto.RegisterRespDto——注册响应参数
	 * 1) 调用 merchant domain 完成用户注册
	 * 2）调用 account domain 创建商户账户
	 */
	Register(ctx context.Context, req *dto.RegisterReqDto) (*dto.RegisterRespDto, error)

	/**
	 * @brief: 登录商户
	 * @param: dto.LoginReqDto——登录请求参数
	 * @return: dto.LoginRespDto——登录响应参数
	 * 1) 调用 account domain 校验商户账号是否合法
	 */
	Login(ctx context.Context, req *dto.LoginReqDto) (*dto.LoginRespDto, error)

	/**
	 * @brief: 创建店铺
	 * @param: dto.CreateStoreReqDto——创建商铺请求参数
	 * @return: dto.CreateStoreRespDto——创建商铺响应参数
	 * 1) 调用 merchant domain 校验商户账号是否合法
	 * 2) 调用 store domain 创建店铺
	 */
	CreateStore(ctx context.Context, req *dto.CreateStoreReqDto) (*dto.CreateStoreRespDto, error)

	/**
	 * @brief: 开放店铺
	 * @param: dto.OpenStoreReqDto——开放商铺请求参数
	 * @return: dto.OpenStoreRespDto——开放商铺响应参数
	 * 1) 调用 store domain 获取并锁定店铺
	 * 2) 调用 store domain 更新店铺状态为上线
	 */
	OnlineStore(ctx context.Context, req *dto.OnlineStoreReqDto) (*dto.OnlineStoreRespDto, error)

	/**
	 * @brief: 下线店铺
	 * @param: dto.OpenStoreReqDto——下线商铺请求参数
	 * @return: dto.OpenStoreRespDto——下线商铺响应参数
	 * 1) 调用 store domain 获取并锁定店铺
	 * 2) 调用 store domain 更新店铺状态为下线
	 */
	OfflineStore(ctx context.Context, req *dto.OfflineStoreReqDto) (*dto.OfflineStoreRespDto, error)

	/**
	 * @brief: 创建商品
	 * @param: dto.CreateGoodsReqDto——创建商品请求参数
	 * @return: dto.CreateGoodsRespDto——创建商品响应参数
	 * 1) 调用 store domain 查询店铺校验合法性
	 * 2) 调用 goods domain 创建商品
	 */
	CreateGoods(ctx context.Context, req *dto.CreateGoodsReqDto) (*dto.CreateGoodsRespDto, error)

	/**
	 * @brief: 发布商品
	 * @param: dto.OnlineGoodsReq——发布商品请求参数
	 * @return: dto.OnlineGoodsRespDto——发布商品响应参数
	 * 1) 调用 goods domain 获取并锁定商品
	 * 2) 调用 goods domain 更新商品为发布状态
	 */
	OnlineGoods(ctx context.Context, req *dto.OnlineGoodsReqDto) (*dto.OnlineGoodsRespDto, error)

	/**
	 * @brief: 下线商品
	 * @param: dto.OfflineGoodsReq——下线商品请求参数
	 * @return: dto.OfflineGoodsRespDto——下线商品响应参数
	 * 1) 调用 goods domain 获取并锁定商品
	 * 2) 调用 goods domain 更新商品为下线状态
	 */
	OfflineGoods(ctx context.Context, req *dto.OfflineGoodsReqDto) (*dto.OfflineGoodsRespDto, error)
}
