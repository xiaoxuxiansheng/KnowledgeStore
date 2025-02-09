package dto

import (
	accountentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/account/entity"
	goodsentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/goods/entity"
	merchantentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/merchant/entity"
	storeentity "github.com/xiaoxuxiansheng/KnowledgeStore/domain/store/entity"
)

/**
 * @brief: 注册商户请求参数
 */
type RegisterBusinessReqDto struct {
	// 账号
	AccountId string
	// 密码
	Password string
}

func (r *RegisterBusinessReqDto) ToMerchantEntity() *merchantentity.MerchantEntity {
	return &merchantentity.MerchantEntity{
		AccountId: r.AccountId,
	}
}

func (r *RegisterBusinessReqDto) ToAccountEntity() *accountentity.AccountEntity {
	return &accountentity.AccountEntity{
		Points: 0,
		AccountVo: accountentity.AccountVo{
			Type:      accountentity.AccountTypeMerchant,
			AccountId: r.AccountId,
			Passwd:    r.Password,
		},
	}
}

/**
 * @brief: 登录商户请求参数
 */
type LoginBusinessReqDto struct {
	// 账号
	AccountId string
	// 密码
	Password string
}

func (l *LoginBusinessReqDto) ToAccountVo() *accountentity.AccountVo {
	return &accountentity.AccountVo{
		Type:      accountentity.AccountTypeMerchant,
		AccountId: l.AccountId,
		Passwd:    l.Password,
	}
}

/**
 * @brief: 创建店铺请求参数
 */
type CreateStoreReqDto struct {
	// 商户账号 id
	AccountId string
	// 店铺名称
	StoreName string
}

func (c *CreateStoreReqDto) ToStoreEntity() *storeentity.StoreEntity {
	return &storeentity.StoreEntity{
		MerchantId: c.AccountId,
		StoreName:  c.StoreName,
	}
}

/**
 * @brief: 开放店铺请求参数
 */
type OnlineStoreReqDto struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
}

func (o *OnlineStoreReqDto) ToStoreEntity() *storeentity.StoreEntity {
	return &storeentity.StoreEntity{
		MerchantId: o.AccountId,
		StoreName:  o.StoreName,
	}
}

/**
 * @brief: 下线店铺请求参数
 */
type OfflineStoreReqDto struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
}

func (o *OfflineStoreReqDto) ToStoreEntity() *storeentity.StoreEntity {
	return &storeentity.StoreEntity{
		MerchantId: o.AccountId,
		StoreName:  o.StoreName,
	}
}

/**
 * @brief: 创建商品请求参数
 */
type CreateGoodsReqDto struct {
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

func (c *CreateGoodsReqDto) ToGoodsEntity() *goodsentity.GoodsEntity {
	return &goodsentity.GoodsEntity{
		GoodsName:    c.GoodsName,
		MerchantId:   c.AccountId,
		Category:     c.Category,
		GoodsContent: c.GoodsContent,
		StoreName:    c.StoreName,
	}
}

/**
 * @brief: 上线商品请求参数
 */
type OnlineGoodsReqDto struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
	// 商品名称
	GoodsName string `json:"goods_name"`
}

func (o *OnlineGoodsReqDto) ToGoodsEntity() *goodsentity.GoodsEntity {
	return &goodsentity.GoodsEntity{
		GoodsName:  o.GoodsName,
		MerchantId: o.AccountId,
		StoreName:  o.StoreName,
	}
}

/**
 * @brief: 下线商品请求参数
 */
type OfflineGoodsReqDto struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
	// 商品名称
	GoodsName string `json:"goods_name"`
}

func (o *OfflineGoodsReqDto) ToGoodsEntity() *goodsentity.GoodsEntity {
	return &goodsentity.GoodsEntity{
		GoodsName:  o.GoodsName,
		MerchantId: o.AccountId,
		StoreName:  o.StoreName,
	}
}
