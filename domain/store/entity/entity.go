package entity

// 店铺分类
type StoreCategory int

type StoreEntity struct {
	// 店铺名称
	StoreName string
	// 店铺从属的商家 id
	MerchantId string
}

type StoreVo struct{}
