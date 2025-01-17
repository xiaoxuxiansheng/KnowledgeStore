package entity

// 店铺分类
type StoreCategory int

const (
	StoreCategoryTech StoreCategory = 1
)

type StoreEntity struct {
	// 店铺唯一 id
	StoreId string
	// 店铺从属的商家 id
	MerchantId string
	// 店铺分类
	Category StoreCategory
}

type StoreVo struct{}
