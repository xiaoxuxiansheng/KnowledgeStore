package tob

import (
	"github.com/xiaoxuxiansheng/KnowledgeStore/application/tob/dto"
)

const (
	ErrnoUnknown = -1
	ErrnoSuccess = 0
)

/**
 * @brief: 注册商家账号请求参数
 */
type RegisterBusinessReqVo struct {
	// 商家账号
	AccountId string `json:"account_id" binding:"required"`
	// 商家密码
	Password string `json:"password" binding:"required"`
}

func (r *RegisterBusinessReqVo) ToDto() *dto.RegisterBusinessReqDto {
	return &dto.RegisterBusinessReqDto{
		AccountId: r.AccountId,
		Password:  r.Password,
	}
}

/**
 * @brief: 登录商家账号请求参数
 */
type LoginBusinessReqVo struct {
	// 账号
	AccountId string `json:"account_id" binding:"required"`
	// 密码
	Password string `json:"password" binding:"required"`
}

func (l *LoginBusinessReqVo) ToDto() *dto.LoginBusinessReqDto {
	return &dto.LoginBusinessReqDto{
		AccountId: l.AccountId,
		Password:  l.Password,
	}
}

/**
 * @brief: 创建店铺请求参数
 */
type CreateStoreReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
}

func (c *CreateStoreReqVo) ToDto() *dto.CreateStoreReqDto {
	return &dto.CreateStoreReqDto{
		AccountId: c.AccountId,
		StoreName: c.StoreName,
	}
}

/**
 * @brief: 发布店铺请求参数
 */
type OnlineStoreReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
}

func (o *OnlineStoreReqVo) ToDto() *dto.OnlineStoreReqDto {
	return &dto.OnlineStoreReqDto{
		AccountId: o.AccountId,
		StoreName: o.StoreName,
	}
}

/**
 * @brief: 下线店铺请求参数
 */
type OfflineStoreReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
}

func (o *OfflineStoreReqVo) ToDto() *dto.OfflineStoreReqDto {
	return &dto.OfflineStoreReqDto{
		AccountId: o.AccountId,
		StoreName: o.StoreName,
	}
}

/**
 * @brief: 创建商品请求参数
 */
type CreateGoodsReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
	// 商品分类
	Category string `json:"category"`
	// 商品名称
	GoodsName string `json:"goods_name"`
	// 商品内容
	GoodsContent string `json:"goods_content"`
}

func (c *CreateGoodsReqVo) ToDto() *dto.CreateGoodsReqDto {
	return &dto.CreateGoodsReqDto{
		AccountId:    c.AccountId,
		StoreName:    c.StoreName,
		Category:     c.Category,
		GoodsName:    c.GoodsName,
		GoodsContent: c.GoodsContent,
	}
}

/**
 * @brief: 发布商品请求参数
 */
type OnlineGoodsReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
	// 商品名称
	GoodsName string `json:"goods_name"`
}

func (o *OnlineGoodsReqVo) ToDto() *dto.OnlineGoodsReqDto {
	return &dto.OnlineGoodsReqDto{
		AccountId: o.AccountId,
		StoreName: o.StoreName,
		GoodsName: o.GoodsName,
	}
}

/**
 * @brief: 下线商品请求参数
 */
type OfflineGoodsReqVo struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
	// 商品名称
	GoodsName string `json:"goods_name"`
}

func (o *OfflineGoodsReqVo) ToDto() *dto.OfflineGoodsReqDto {
	return &dto.OfflineGoodsReqDto{
		AccountId: o.AccountId,
		StoreName: o.StoreName,
		GoodsName: o.GoodsName,
	}
}
