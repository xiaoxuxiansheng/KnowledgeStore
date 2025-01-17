package entity

type GoodsType int

type GoodsEntity struct {
	// 商品唯一 id
	GoodsId string
	// 商品从属的店铺 id
	StoreId string
	// 商品类型
	Type GoodsType
}

type GoodsVo struct{}
