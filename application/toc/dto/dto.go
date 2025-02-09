package dto

/**
 * @brief: 注册账号请求参数
 */
type RegisterCutomerReqDto struct {
	// 用户账号
	AccountId string
	// 商家密码
	Password string
}

/**
 * @brief: 登录请求参数
 */
type LoginCutomerReqDto struct {
	// 用户账号
	AccountId string
	// 商家密码
	Password string
}

/**
 * @brief: 检索店铺请求参数
 */
type ListStoresReqDto struct {
	StoreName string
	Page      int
	Limit     int
}

/**
 * @brief: 检索店铺响应参数
 */
type ListStoresRespDto struct {
	Total int
	List  []*StoreDataDto
}

type StoreDataDto struct {
	StoreName string
}

/**
 * @brief: 检索商品请求参数
 */
type ListGoodsReqDto struct {
	// 商店名称
	StoreName string
	// 商品分类
	Category string
	// 商品名称
	GoodsName string
	// 页码
	Page int
	// 大小
	Limit int
}

/**
 * @brief: 检索商品响应参数
 */
type ListGoodsRespDto struct {
	Total int
	List  []*GoodsDataDto
}

type GoodsDataDto struct {
	GoodsName string
	Category  string
}

/**
 * @brief: 查询商品详情请求参数
 */
type QueryGoodsReqDto struct {
	AccountId string
	StoreName string
	GoodsName string
}

/**
 * @brief: 查询商品详情响应参数
 */
type QueryGoodsRespDto struct {
	Owned        bool
	GoodsContent string
	Category     string
	GoodsName    string
}

/**
 * @brief: 购买商品请求参数
 */
type BuyGoodsReqDto struct {
	AccountId string
	StoreName string
	GoodsName string
}

type ListOrderReqDto struct {
	AccountId string
	Status    string
}

type ListOrderRespDto struct {
	Total  int
	Status string
	List   []*OrderDataDto
}

type OrderDataDto struct {
	OrderId   string
	Amount    int
	StoreName string
	GoodsName string
}

/**
 * @brief: 支付订单请求参数
 */
type PayOrderReqDto struct {
	AccountId string
	OrderId   string
}
