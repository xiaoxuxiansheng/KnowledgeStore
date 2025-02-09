package dto

/**
 * @brief: 注册商户请求参数
 */
type RegisterBusinessReqDto struct {
	// 账号
	AccountId string
	// 密码
	Password string
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

/**
 * @brief: 创建店铺请求参数
 */
type CreateStoreReqDto struct {
	// 商户账号 id
	AccountId string
	// 店铺名称
	StoreName string
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

/**
 * @brief: 下线店铺请求参数
 */
type OfflineStoreReqDto struct {
	// 商户账号 id
	AccountId string `json:"account_id"`
	// 店铺名称
	StoreName string `json:"store_name"`
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
