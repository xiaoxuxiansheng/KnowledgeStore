package entity

type GoodsEntity struct {
	// 商品名称
	GoodsName string
	// 商品从属的店铺 id
	MerchantId string
	// 商品类型
	Category string
	// 商品类型
	GoodsContent string
	// 店铺名称
	StoreName string
	Points    int
	// 是否已拥有过该商品
	Owned bool
}
